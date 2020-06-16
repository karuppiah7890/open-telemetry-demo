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

Now, I want to see how to incorporate open telemetry in this!

Checking this

https://github.com/open-telemetry/opentelemetry-go

And there's an example here

https://github.com/open-telemetry/opentelemetry-go/blob/master/example/README.md

So I cloned the source code in my local

```bash
$ git clone git@github.com:open-telemetry/opentelemetry-go.git
```

Ran the server

```bash
$ cd opentelemetry-go/example/http
$ go run server/server.go
```

And ran the client too

```bash
$ cd opentelemetry-go/example/http
$ go run client/client.go
```

I could see the following output in the server

```bash
{
	"SpanContext": {
		"TraceID": "91285635bd7088d30a2aabfd2a48fa60",
		"SpanID": "e706a775c8091a82",
		"TraceFlags": 1
	},
	"ParentSpanID": "29d24de648b1d82d",
	"SpanKind": 1,
	"Name": "hello",
	"StartTime": "2020-06-12T08:11:50.127696+05:30",
	"EndTime": "2020-06-12T08:11:50.127727324+05:30",
	"Attributes": [
		{
			"Key": "http.method",
			"Value": {
				"Type": "STRING",
				"Value": "GET"
			}
		},
		{
			"Key": "http.target",
			"Value": {
				"Type": "STRING",
				"Value": "/hello"
			}
		},
		{
			"Key": "http.scheme",
			"Value": {
				"Type": "STRING",
				"Value": "http"
			}
		},
		{
			"Key": "http.host",
			"Value": {
				"Type": "STRING",
				"Value": "localhost:7777"
			}
		},
		{
			"Key": "http.user_agent",
			"Value": {
				"Type": "STRING",
				"Value": "Go-http-client/1.1"
			}
		},
		{
			"Key": "http.flavor",
			"Value": {
				"Type": "STRING",
				"Value": "1.1"
			}
		},
		{
			"Key": "net.transport",
			"Value": {
				"Type": "STRING",
				"Value": "IP.TCP"
			}
		},
		{
			"Key": "net.peer.name",
			"Value": {
				"Type": "STRING",
				"Value": "[::1]"
			}
		},
		{
			"Key": "net.peer.port",
			"Value": {
				"Type": "INT64",
				"Value": 49417
			}
		},
		{
			"Key": "net.host.name",
			"Value": {
				"Type": "STRING",
				"Value": "localhost"
			}
		},
		{
			"Key": "net.host.port",
			"Value": {
				"Type": "INT64",
				"Value": 7777
			}
		}
	],
	"MessageEvents": [
		{
			"Name": "handling this...",
			"Attributes": null,
			"Time": "2020-06-12T08:11:50.127709+05:30"
		}
	],
	"Links": null,
	"StatusCode": 0,
	"StatusMessage": "",
	"HasRemoteParent": true,
	"DroppedAttributeCount": 0,
	"DroppedMessageEventCount": 0,
	"DroppedLinkCount": 0,
	"ChildSpanCount": 0,
	"Resource": null,
	"InstrumentationLibrary": {
		"Name": "example/server",
		"Version": ""
	}
}
```

I could see the following output in the client side

```bash
Sending request...
{
	"SpanContext": {
		"TraceID": "91285635bd7088d30a2aabfd2a48fa60",
		"SpanID": "a92a391549494e17",
		"TraceFlags": 1
	},
	"ParentSpanID": "fcd38a851a862f69",
	"SpanKind": 3,
	"Name": "http.dns",
	"StartTime": "2020-06-12T08:11:50.122266+05:30",
	"EndTime": "2020-06-12T08:11:50.125720277+05:30",
	"Attributes": [
		{
			"Key": "http.host",
			"Value": {
				"Type": "STRING",
				"Value": "localhost"
			}
		}
	],
	"MessageEvents": null,
	"Links": null,
	"StatusCode": 0,
	"StatusMessage": "",
	"HasRemoteParent": false,
	"DroppedAttributeCount": 0,
	"DroppedMessageEventCount": 0,
	"DroppedLinkCount": 0,
	"ChildSpanCount": 0,
	"Resource": null,
	"InstrumentationLibrary": {
		"Name": "go.opentelemetry.io/otel/instrumentation/httptrace",
		"Version": ""
	}
}
{
	"SpanContext": {
		"TraceID": "91285635bd7088d30a2aabfd2a48fa60",
		"SpanID": "e5756a90f465aac7",
		"TraceFlags": 1
	},
	"ParentSpanID": "fcd38a851a862f69",
	"SpanKind": 3,
	"Name": "http.connect",
	"StartTime": "2020-06-12T08:11:50.126454+05:30",
	"EndTime": "2020-06-12T08:11:50.126896362+05:30",
	"Attributes": [
		{
			"Key": "http.remote",
			"Value": {
				"Type": "STRING",
				"Value": "[::1]:7777"
			}
		}
	],
	"MessageEvents": null,
	"Links": null,
	"StatusCode": 0,
	"StatusMessage": "",
	"HasRemoteParent": false,
	"DroppedAttributeCount": 0,
	"DroppedMessageEventCount": 0,
	"DroppedLinkCount": 0,
	"ChildSpanCount": 0,
	"Resource": null,
	"InstrumentationLibrary": {
		"Name": "go.opentelemetry.io/otel/instrumentation/httptrace",
		"Version": ""
	}
}
{
	"SpanContext": {
		"TraceID": "91285635bd7088d30a2aabfd2a48fa60",
		"SpanID": "fcd38a851a862f69",
		"TraceFlags": 1
	},
	"ParentSpanID": "29d24de648b1d82d",
	"SpanKind": 3,
	"Name": "http.getconn",
	"StartTime": "2020-06-12T08:11:50.122146+05:30",
	"EndTime": "2020-06-12T08:11:50.127134248+05:30",
	"Attributes": [
		{
			"Key": "http.host",
			"Value": {
				"Type": "STRING",
				"Value": "localhost:7777"
			}
		},
		{
			"Key": "http.remote",
			"Value": {
				"Type": "STRING",
				"Value": "[::1]:7777"
			}
		},
		{
			"Key": "http.local",
			"Value": {
				"Type": "STRING",
				"Value": "[::1]:49417"
			}
		}
	],
	"MessageEvents": null,
	"Links": null,
	"StatusCode": 0,
	"StatusMessage": "",
	"HasRemoteParent": false,
	"DroppedAttributeCount": 0,
	"DroppedMessageEventCount": 0,
	"DroppedLinkCount": 0,
	"ChildSpanCount": 2,
	"Resource": null,
	"InstrumentationLibrary": {
		"Name": "go.opentelemetry.io/otel/instrumentation/httptrace",
		"Version": ""
	}
}
{
	"SpanContext": {
		"TraceID": "91285635bd7088d30a2aabfd2a48fa60",
		"SpanID": "8365d7df6d9f2c4b",
		"TraceFlags": 1
	},
	"ParentSpanID": "29d24de648b1d82d",
	"SpanKind": 3,
	"Name": "http.send",
	"StartTime": "2020-06-12T08:11:50.127347+05:30",
	"EndTime": "2020-06-12T08:11:50.127352261+05:30",
	"Attributes": null,
	"MessageEvents": null,
	"Links": null,
	"StatusCode": 0,
	"StatusMessage": "",
	"HasRemoteParent": false,
	"DroppedAttributeCount": 0,
	"DroppedMessageEventCount": 0,
	"DroppedLinkCount": 0,
	"ChildSpanCount": 0,
	"Resource": null,
	"InstrumentationLibrary": {
		"Name": "go.opentelemetry.io/otel/instrumentation/httptrace",
		"Version": ""
	}
}
{
	"SpanContext": {
		"TraceID": "91285635bd7088d30a2aabfd2a48fa60",
		"SpanID": "635882cda0da2883",
		"TraceFlags": 1
	},
	"ParentSpanID": "29d24de648b1d82d",
	"SpanKind": 3,
	"Name": "http.receive",
	"StartTime": "2020-06-12T08:11:50.12852+05:30",
	"EndTime": "2020-06-12T08:11:50.128639383+05:30",
	"Attributes": null,
	"MessageEvents": null,
	"Links": null,
	"StatusCode": 0,
	"StatusMessage": "",
	"HasRemoteParent": false,
	"DroppedAttributeCount": 0,
	"DroppedMessageEventCount": 0,
	"DroppedLinkCount": 0,
	"ChildSpanCount": 0,
	"Resource": null,
	"InstrumentationLibrary": {
		"Name": "go.opentelemetry.io/otel/instrumentation/httptrace",
		"Version": ""
	}
}
{
	"SpanContext": {
		"TraceID": "91285635bd7088d30a2aabfd2a48fa60",
		"SpanID": "29d24de648b1d82d",
		"TraceFlags": 1
	},
	"ParentSpanID": "0000000000000000",
	"SpanKind": 1,
	"Name": "say hello",
	"StartTime": "2020-06-12T08:11:50.122029+05:30",
	"EndTime": "2020-06-12T08:11:50.128762183+05:30",
	"Attributes": null,
	"MessageEvents": null,
	"Links": null,
	"StatusCode": 0,
	"StatusMessage": "",
	"HasRemoteParent": false,
	"DroppedAttributeCount": 0,
	"DroppedMessageEventCount": 0,
	"DroppedLinkCount": 0,
	"ChildSpanCount": 4,
	"Resource": null,
	"InstrumentationLibrary": {
		"Name": "example/client",
		"Version": ""
	}
}
Response Received: Hello, world!



Waiting for few seconds to export spans ...

Inspect traces on stdout
```

So, it said the traces will be on the standard output, the console, in the order
they were exported.

Now I need to check the code and then understand what the above output means!

I read the code. I can understand to some extent what they are doing. As in,
the functions they call. But I don't get the meaning of the functions or why
they are called. I can only understand the basics of the code.

Digging in, I can see that the open telemetry specification can help me in this
regard. So, I'm reading that now!

https://github.com/open-telemetry/opentelemetry-specification

Starting with the overview here

https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/overview.md

---

Okay, so, looking at the overview, I understood some basics. What's a trace,
what's a span, what's span context, trace ID, span ID, parent ID etc.

And there were more stuff, which I didn't completely understand, but that's
fine. I'll get to it later I guess, when the time comes.

I tried read the spans again this time. It made more sense. Also, when I looked
at the names of the spans at the `Name` field, I could understand better!

So, the client has the following names in it's spans

```
http.dns
http.connect
http.getconn
http.send
http.receive
say hello
```

And the server has this one name in it's one span

```
hello
```

Now, it kind of makes sense about what that many spans were all about.
And then there's this attributes field, which also gives some extra information
about each span. Like, for dns it has details about the dns name that was
resolved.

For `hello` on the server side, it had attributes like what's the request method,
what's the transport protocol (TCP) and then version of http, http scheme - http,
and then URL path `/hello`, user agent, server host and port, client host and
port.

Similarly for other spans too, there were appropriate attributes! :) 

Now, I gotta see how to do this in my fibonacci program and also export it to
some place to visualize this cool thing - traces and spans.

I read the client and server code. I noticed what they do.

In the client code, they first initialize a tracer, that is - create an exporter,
in this case it was a stdout exporter and then create a trace provider using the
exporter and set it as the global trace provider. And something about sampling,
where sampling is always done. And then they create a correlation context which
I can't find in the span output. Then the client creates a span, with a name
and then some stuff is injected into the request. I think this stuff includes
span context, attributes and correlation related entries.

After that the client sends the request!

In the server code, it extracts some stuff from the client's request. This is
the same stuff client injected into the request - span context, correlation
related entries and attributes. It then creates a req with the correlation
details (not sure what this part is). And then a span is started / created!
It's given a name and attributes and the remote span context from the client
so that the span that is created gets a parent span ID based on the client's
span details. There's also an event that's added to the span and then the
response is sent and the span is ended.

Now, I need to learn how to use some visualization like Jaegar, Zipkin. Actually,
check if I'm even right about these tools. There's also some tool called
open telemetry collector. Gotta see what that is.

Okay, now I'm checking this out -

https://opentelemetry.io/docs/workshop/resources/
https://docs.google.com/presentation/d/1nVhLIyqn_SiDo78jFHxnMdxYlnT0b7tYOHz3Pu4gzVQ/edit#slide=id.g80f778c091_0_156

This workshop material is really sleek and good! I think I could use this if I'm
going to give a talk on it ;) :D

---

Okay, so I got stuck looking for how to export to open collector. I did look
for different example code in different places. Today I spent a lot of time
trying to make the example in this repo work

https://github.com/open-telemetry/opentelemetry-collector/tree/master/examples

Finally I realized there's an example for collector export in this
repo itself


