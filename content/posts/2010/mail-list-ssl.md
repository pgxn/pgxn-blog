---
title: Mail List, SSL
slug: mail-list-ssl
date: 2010-10-09T06:01:08Z
aliases: [/post/1274234177/mail-list-ssl, /post/1274234177]
tags: [Mail List, SSL, Basic Auth]
---

I'm *this* close to having PGXN Manager ready for a limited beta. I've got it
running on [Kineticode]'s server, and have been tweaking things here and
there, fixing some bugs and filling in a few missing bits. In the next couple
of days I'll get some more kinks worked out and then start inviting folks in.
If you're interested -- especially if you have an extension you'd like to
release, leave a comment here to let me know.

Better yet, sign up for our new [mail list]. I've set this up so that PGXN
users can have a place to meet and discuss things (naturally). I expect
discussion will be about how to create proper distribution archives (hint:
it's all about the [`META.json`] file); issues with PGXN Manager or the
network itself, and directions for ongoing development.

Oh, one thing I wanted to run by you here: I currently have two sections in
PGXN Manager: public and users only. Right now the only difference is that all
the users-only stuff is under the `/auth/` URI, which is itself limited by a
basic auth challenge. I find this setup a bit hinky, though, because the link
to `/about`, for example, appears in the nav menu for both sections, and if
you're logged in and click it, it will look like you're logged out.

So I was thinking perhaps that I'd change it so that the difference is that
user-only is on port 443 (SSL) and public on port 80. That way I could have
links to `about` on both and one wouldn't appear to be logged out by clicking
it from the SSL site. Because, you know, they're effectively two different
sites.

Thoughts? Is using the SSL divide perhaps the most natural way to separate
user-only access from public access?

More soon.

  [Kineticode]: https://kineticode.com/
  [mail list]: https://groups.google.com/group/pgxn-users
  [`META.json`]: https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec
