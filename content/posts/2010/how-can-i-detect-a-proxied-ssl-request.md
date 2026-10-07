---
title: How Can I Detect a Proxied SSL Request?
slug: how-can-i-detect-a-proxied-ssl-request
date: 2010-10-11T15:00:29Z
aliases: [/post/1291485816/how-can-i-detect-a-proxied-ssl-request, /post/1291485816]
tags: [mod_proxy, mod_ssl, SSL, Proxy, URI, Apache, Plack, Request]
---

I've been thinking about how best to create two sections of the PGXN Manager
site, one that requires authentication and is on SSL and one that's public. So
far, I've just had all the authenticated stuff go to /auth (because I'm using
basic auth), but what I want to do now is require authentication if the
connection is via SSL. However, using a reverse proxy with mod_proxy to map to
the PGXN Manger Plack app running on an internal port, I've found no
environment variables passed through that would allow me, in code, to
determine whether a request is via a proxied SSL or non-SSL connection. Most
irritating.

So tell me what you think about the alternative plan I've come up with.

I'll have two plack apps, one mapped to /auth and one mapped to /no-auth. The
former will require authentication and the latter will not. They'll have
separate dispatch tables, of course. Then I'll have the SSL site proxied to
/auth and the non-SSL site proxied to /no-auth. Makes sense, right?

The only hangup I can see (though maybe you can see others?) is that my
current method of generating URIs knows nothing about proxies. So If I link to
/auth/account, when requests come through the proxy, it should actually create
a link to /account. Does it make sense to use relative links instead of
absolute links for all links to avoid this issue? I think it might be kind of
annoying, because not all the code is aware of the current URI, though there
are ways to deal with that.

Thoughts?

I guess I could stick with absolute URLs, and then have the [`uri_for`] use
`$req−>uri−>path` for a proxied request and `$req−>path_info` for a
non-proxied request.

Sure would be nice if there was some way to tell from the environment that a
request was forwarded from an SSL connection, though. Alas, there are only
three extra environment variable set by the proxy server:

- HTTP_X_FORWARDED_FOR
- HTTP_X_FORWARDED_HOST
- HTTP_X_FORWARDED_SERVER

No HTTP_X_FORWARDED_PORT or HTTP_X_FORWARDED_SSL or SSL_ENABLED or anything
like that. Maybe I'm missing something in my reading of the [mod_proxy
documentation]?

  [`uri_for`]: https://github.com/theory/pgxn-manager/blob/master/lib/PGXN/Manager/Request.pm#L14
  [mod_proxy documentation]: https://httpd.apache.org/docs/2.2/mod/mod_proxy.html
