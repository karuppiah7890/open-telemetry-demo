# Thoughts

These are the thoughts and my experiences while building
this demo.

I started off with a simple golang web server. I didn't
want to import the big web frameworks / libraries like
gin, fasthttp etc.

Just plain http.

Added a ping endpoint. It will respond to all kind of request methods
actually 😅 and can also take request bodies but of course it's not going to
read them or anything. Standard plain text response always. that's it.

I'm not really adding any tests here. I want to keep it simple. And no tests
really. I'm okay with it. I have noticed myself overdoing testing, so I'll have
to learn better and then come back to do better testing if I really want to do
it 😅

Nex is the Fibonacci endpoint. It will take only `POST` requests and will read
from the request body and respond to it. O.o I'm afraid I might need some tests.
Anyways. Meh :P Let's see. I'll try to keep it too simple. I just want to see
a working demo. Not gonna go crazy about robustness in this case :)

So I built the fibonacci service to reply for the numbers 0 and 1 as they are
easy! So, this is how it currently looks like!

```bash
$ curl -i localhost:8080/ping
HTTP/1.1 200 OK
Date: Thu, 11 Jun 2020 13:36:38 GMT
Content-Length: 4
Content-Type: text/plain; charset=utf-8

pong

$ curl -i localhost:8080/
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 11 Jun 2020 13:36:51 GMT
Content-Length: 19

404 page not found

$ curl -i localhost:8080/blah
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 11 Jun 2020 13:36:54 GMT
Content-Length: 19

404 page not found

$ curl -i localhost:8080/fibonacci
HTTP/1.1 400 Bad Request
Date: Thu, 11 Jun 2020 13:37:14 GMT
Content-Length: 0

$ curl -i -X PUT localhost:8080/fibonacci
HTTP/1.1 400 Bad Request
Date: Thu, 11 Jun 2020 13:37:19 GMT
Content-Length: 0

$ curl -i -X POST localhost:8080/fibonacci
HTTP/1.1 400 Bad Request
Date: Thu, 11 Jun 2020 13:37:23 GMT
Content-Length: 0

$ curl -i -X POST -H "Content-Type: text/plain" localhost:8080/fibonacci
HTTP/1.1 400 Bad Request
Date: Thu, 11 Jun 2020 13:37:34 GMT
Content-Length: 0

$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 11 Jun 2020 13:37:40 GMT
Content-Length: 29

unexpected end of JSON input

$ curl -i -X POST -H "content-type: application/json" localhost:8080/fibonacci
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 11 Jun 2020 13:38:07 GMT
Content-Length: 29

unexpected end of JSON input

$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 0 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 13:38:36 GMT
Content-Length: 21

{"FibonacciNumber":0}

$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 1 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 13:38:45 GMT
Content-Length: 21

{"FibonacciNumber":1}

$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 2 }'
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 11 Jun 2020 13:38:55 GMT
Content-Length: 25

don't know the answer :p
```

And for another too, same

```
$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 10 }'
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 11 Jun 2020 13:39:04 GMT
Content-Length: 25

don't know the answer :p
```

So, for `0` and `1` it works now. Gotta make it work for others!


