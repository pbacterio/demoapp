FROM golang:alpine as builder
RUN apk --no-cache add bash upx;
RUN wget -q -O - https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ADD . /go/src/github.com/pbacterio/demoapp/
WORKDIR /go/src/github.com/pbacterio/demoapp
RUN ./build.sh

FROM scratch
COPY --from=builder /go/src/github.com/pbacterio/demoapp/build/linux/demoapp /
CMD ["/demoapp"]
