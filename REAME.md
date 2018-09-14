demoapp
=======

Toy microservice to play, learn or do demos!


API
---

* Just OK:
  * `/`, `/ok`: Just answer ok.

* Getting info:
  * `/info` | General info: hostname, ip, uid & gid.
  * `/env` | Server environment variables list.
  * `/req` | Request headers, method, path & remote address.
  
* Failing:
  * `/fail`: Generates a 500 Internal Server Error.
  * `/fail/{STATUS-CODE}`, `/fail/{STATUS-CODE}/{STATUS_TEXT}`: Answer a custom http status code.

* Crazy things:
  * '/rand/fail/{PERCENT}': Generates an error a random percent of requests.
  * `/kill`: Just stop the microservice.
  * `/slow/{SECONDS}`: A request that takes SECONDS to finish.


Examples
--------




Use it
------