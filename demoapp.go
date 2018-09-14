package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func ok(c echo.Context) error {
	return c.String(http.StatusOK, "ok\n")
}

func fail(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.String(http.StatusInternalServerError, "fail!\n")
	}
	codeNum, err := strconv.Atoi(code)
	if err != nil {
		return err
	}
	text := c.Param("text")
	if text == "" {
		text = "fail!"
	}
	return c.String(codeNum, text+"\n")
}

func randFail(c echo.Context) error {
	percent, err := strconv.Atoi(c.Param("percent"))
	if err != nil {
		return err
	}
	if rand.Intn(100) <= percent {
		return fail(c)
	}
	return ok(c)
}

func env(c echo.Context) error {
	w := c.Response()
	for _, env := range os.Environ() {
		w.Write([]byte(env + "\n"))
	}
	return nil
}

func req(c echo.Context) error {
	r := c.Request()
	w := c.Response()
	fmt.Fprintf(w, "URL: %s\n", r.RequestURI)
	fmt.Fprintf(w, "method: %s\n", r.Method)
	fmt.Fprintf(w, "protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "remote address: %s\n", r.RemoteAddr)
	for name := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", name, r.Header.Get(name))
	}
	return nil
}

func info(c echo.Context) error {
	w := c.Response()
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "hostname: %s\n", hostname)
	fmt.Fprintf(w, "ipaddress: %s\n", getIpList())
	fmt.Fprintf(w, "uid: %d\n", os.Getuid())
	fmt.Fprintf(w, "gid: %d\n", os.Getgid())
	return nil
}

func getIpList() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	ipList := []string{}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ipList = append(ipList, ip.String())
		}
	}
	return strings.Join(ipList, ", ")
}

func kill(c echo.Context) error {
	os.Exit(0)
	return nil
}

func slow(c echo.Context) error {
	seconds := c.Param("seconds")
	if seconds == "" {
		return c.String(http.StatusInternalServerError, "fail!\n")
	}
	secondsNum, err := strconv.ParseFloat(seconds, 64)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(int64(secondsNum * float64(time.Second))))
	return c.String(http.StatusOK, "ok\n")
}

func main() {
	e := echo.New()

	// Just OK
	e.Any("/", ok)
	e.Any("/ok", ok)

	// Getting info
	e.Any("/info", info)
	e.Any("/env", env)
	e.Any("/req", req)

	// Failing
	e.Any("/fail", fail)
	e.Any("/fail/:code", fail)
	e.Any("/fail/:code/:text", fail)
	e.Any("/rand/fail/:percent", randFail)
	e.Any("/kill", kill)
	e.Any("/slow/:seconds", slow)

	port := os.Getenv("DEMO_PORT")
	if port == "" {
		port = "8080"
	}

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":" + port))
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
