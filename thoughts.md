# Thoughts

These are the thoughts and my experiences while building
this demo.

I started off with a simple golang web server. I didn't
want to import the big web frameworks / libraries like
gin, fasthttp etc.

Just plain http.

Used this for reference

https://gowebexamples.com/http-server/

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

The next step, is to use fibonacci service itself to get the answer to bigger
sequence number. Actually there are multiple ways to get the number given the
sequence number n, through mathematical ways. But I want to use the brute force
method and do it in a simple manner.

That is, when a number greater than 1 comes as input to the service, it uses a
HTTP client to call itself to get the value of the number and it keeps going
on in a loop till the 0the or 1st element is given as answer and then the
responses are sent back slowly. This is similar to recusive functions.

I'm going to create a client first for this! :)

The resource that I'm going to use for this is

http://networkbit.ch/golang-http-client/

Okay, so I created a client for the service and wired it up and everything but
messed up with something and got a weird error. Finally added some extra strings
to the errors to understand where the error was occuring. Found out it's at
parsing. Of course. I messed up with request path 🙈 Fixing it now! It's mostly
hardcoded, to run in the local at a fixed port now. It want it to just work.
Meh :P

For a second I thought it doesn't work when I gave the sequence number `2` and
I got the result as `1`. The code was correct, I even fixed the copy paste
error of putting `n-1` twice, instead of `n-1` and `n-2`. Anyways, the result
is actually correct. I tried it with `3`, `4` too. Gotta try bigger numbers now!

```bash
$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 2 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 14:12:57 GMT
Content-Length: 21

{"FibonacciNumber":1}

$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 3 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 14:11:09 GMT
Content-Length: 21

{"FibonacciNumber":2}

$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 4 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 14:11:13 GMT
Content-Length: 21

{"FibonacciNumber":3}
```

For bigger numbers, I tried `10` and `20`. `20` took time!

```bash
$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 10 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 14:13:11 GMT
Content-Length: 22

{"FibonacciNumber":55}

$ curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 20 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 14:13:22 GMT
Content-Length: 24

{"FibonacciNumber":6765}

$ time curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 20 }'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 11 Jun 2020 14:14:09 GMT
Content-Length: 24

{"FibonacciNumber":6765}
real    0m5.816s
user    0m0.006s
sys     0m0.008s

$ time curl -i -X POST -H "Content-Type: application/json" localhost:8080/fibonacci -d '{ "sequenceNumber": 30 }'
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 11 Jun 2020 14:14:32 GMT
Content-Length: 177

error while sending fibonacci request for sequence number 29: Post "http://localhost:8080/fibonacci": context deadline exceeded (Client.Timeout exceeded while awaiting headers)

real    0m10.013s
user    0m0.006s
sys     0m0.009s

```

If you notice, for `30`, it just timed out. This is because I had set a time
out for all client requests - `10 seconds`. Looks like the first request that
got sent for `29` sequence number took a lot of time, more than 10 seconds and
hence everything just failed.
