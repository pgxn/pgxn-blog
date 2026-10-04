---
title: Inter-documentation and image links now work on PGXN
slug: relative-links
date:  2026-09-08T22:12:29Z
lastMod: 2026-09-08T22:12:29Z
description: |
  From the Department of It’s About Time: Inter-documentation and image links
  now work on PGXN.
tags: [Images, Links, Site]
aliases: [/post/827227334118604800/from-the-department-of-its-about-time]
---

Way back in 2015, I opened a [PGXN API issue] to allow links between documents
rendered by PGXN to work. In 2024 I followed up with [another issue] to enable
relative links to images to work. I mean, everyone wants this, right? It's how
your favorite source code sit works.

Now so does PGXN. As of today, links between documents and to images work.
Check it out in the newly-released [chdb extension docs], which both link to
the `chdb` and `chdb_hook` docs and to some nice benchmark graphs.

Or see the [Apache Age docs], which includes not only documentation images but
also decorative SVGs for the headers.

I'm gradually reindexing all of the extensions so that any other documentation
with such links work work; it should be done by tomorrow. Then, at along last,
your extensions can look their very best.

  [PGXN API issue]: https://github.com/pgxn/pgxn-api/issues/38
    "Teach Parser to resolve doc links to the .html files they generate"
  [another issue]: https://github.com/pgxn/pgxn-api/issues/40
    "Teach Parser to resolve relative Image Links"
  [chdb extension docs]: https://pgxn.org/dist/chdb/0.1.1/ "chdb 0.1.1 on PGXN"
  [Apache Age docs]: https://pgxn.org/dist/apacheage/1.4.0/ "ApacheAge 1.4.0 on PGXN"
