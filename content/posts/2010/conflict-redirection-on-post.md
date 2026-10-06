---
title: Conflict and Redirection on POST
slug: conflict-redirection-on-post
date: 2010-09-29T21:10:00Z
aliases: [/post/1212136083/conflict-redirection-on-post, /post/1212136083]
tags: [Http, Post, Redirect, RDG, Post/Redirect/Get, Conflict, HTTP Status Codes]
---

Had an interesting discussion on [#plack]. The upload form, which takes a POST
request for an upload, sends a redirect on a successful form submission. This
is known as the [Post/Redirect/Get (RDG)] pattern, though I didn't know that
before [DuckDuckGoing it] it today. But I realized that the code was *not*
redirecting on a failed form submission, but reloading the form. I was
thinking that one should always redirect on POST, and so was looking into a
Rails-like `flash` pattern to cache an error message and the form contents on
the redirect. But in this discussion, I learned a few things:

1.  One should use [HTTP 1.1 status code 303] rather than [status code 302]
    for these sorts of redirects. Apparently, this status code, called "See
    Other," is specifically designed for use redirecting from a POST. Thanks
    to [counfound] for pointing that out.

2.  The whole point of RDG is to prevent a double form submission. But the
    truth is, in the case of an error, you *want* a second form submission.
    For errors -- such as a username conflict or a malformed distribution
    archive -- the POST failed, so you show the form again along with a
    message about the problem and how to fix it.

What's nice about this is that it's already the way I was doing it! In fact,
for such errors, I'm returning [HTTP status code 409], which indicates that
the request failed due to a conflict, and the user should fix it and resubmit.
So no need to redirect on failure. Yay!

Note that for XMLHttpRequests, I'm not redirecting at all, but simply
returning the proper status code (success or conflict, generally) and a
fragment to be used to show an error message (or "success") (in HTML, JSON, or
plain text, depending on the requestor's preferred type).

I think this works pretty well, and I'm pleased to be making good use of HTTP.
Does it make sense to you?



  [#plack]: irc://irc.perl.org/#plack
  [Post/Redirect/Get (RDG)]: https://en.wikipedia.org/wiki/Post/Redirect/Get
  [DuckDuckGoing it]: https://duckduckgo.com/?q=redirect+from+post
  [HTTP 1.1 status code 303]: https://tools.ietf.org/html/rfc2616#section-10.3.4
  [status code 302]: https://tools.ietf.org/html/rfc2616#section-10.3.3
  [counfound]: https://blog.weftsoar.net/
  [HTTP status code 409]: https://tools.ietf.org/html/rfc2616#section-10.4.10
