---
title: PGXN API Docs Published
slug: api-docs-published
date: 2011-04-27T00:27:55Z
aliases: [/post/4970682941/api-docs-published, /post/4970682941]
tags: [API, Documentation, Docs, API Docs, Zip, Client, Daniele Varrazzo]
---

The [PGXN API Documentation] is up! I've just finished writing the docs for
the lightweight REST API provided by all PGXN mirrors. It also documents how
the same APIs differ when provided by the [API server]. Their are four more
documents to write still, APIs provided by the API server but not the mirrors
(including full-text search); I should be able to get those done tomorrow.

So if you're interested in use the API for various things, please [have a
look][PGXN API Documentation]! I'd appreciate any feedback or corrections.

Speaking of users of the API, [Daniele Varrazzo] has started writing a [PGXN
client app] in Python. This is so awesome! It makes me happy that people are
able to just get going on this. And Daniele did it before I'd written the
docs, just from reading the PGXN source code. Nice!

Anyway, getting these docs written was the last barrier I had to releasing
various pieces of PGXN on CPAN and generally being able to get on with my
life. I wanted to write these docs first so that I would have a little more
freedom to change things if something struct me as especially stupid.
Fortunately, as I've written the [main site] as a thin client for the API,
most of the lameness has already been excised. I think the only change I'd
like to make is to change the `{char}` URI template variable for the
[`userlist` API] to `{letter}`, because only lowercased ASCII letters a-z are
allowed. It's more accurate.

Another thing I think I'll change is the file name suffix used for the
distribution download files. Currently it's `.pgz`, but after [Daniele
complained] about it, I [asked around] and the concensus seems to be to change
it to `.zip`. I'll likely do that tomorrow, too. Fortunately it's pretty easy:
Just edit the configuration file and rename the files on the master mirror. At
least it *should* be that easy!

There was one other thing in the APIs the felt a bit silly, but I don't
remember what it was, so screw it.

Anyway, feedback on the API docs would be greatly appreciated. And -- *get
hacking!*.

  [PGXN API Documentation]: https://github.com/pgxn/pgxn-api/wiki
  [API server]: https://api.pgxn.org/
  [Daniele Varrazzo]: https://profiles.google.com/daniele.varrazzo/about
  [PGXN client app]: https://github.com/dvarrazzo/pgxn-client/
  [main site]: https://pgxn.org/
  [`userlist` API]: https://github.com/pgxn/pgxn-api/wiki/userlist-api
  [Daniele complained]: https://blog.pgxn.org/post/4854707157/pgxn-manager
  [asked around]: https://groups.google.com/group/pgxn-users/browse_thread/thread/1209538be03be8a4
