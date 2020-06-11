# Thoughts

These are the thoughts and my experiences while building
this demo.

I started off with a simple golang web server. I didn't
want to import the big web frameworks / libraries like
gin, fasthttp etc.

Just plain http.

Added a ping endpoint. It will respond to all kind of request methods
actually 😅 and can also take request bodies but of course it's not going to
read them or anything. standard plain text response always. that's it.

I'm not really adding any tests here. I want to keep it simple. And no tests
really. I'm okay with it. I have noticed myself overdoing testing, so I'll have
to learn better and then come back to do better testing if I really want to do
it 😅
