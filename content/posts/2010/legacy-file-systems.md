---
title: On Legacy File Systems
slug: legacy-file-systems
date: 2010-09-08T23:41:00Z
aliases: [/post/1088721089/legacy-file-systems, /post/1088721089]
tags: [URI Templates, File Systems, Hashing, Names, Mirrors]
---

When we last looked at the [organization of the mirror], I was pretty happy
with the design except for one thing: The use of the letter hashing variables
in the URI templates. It's just ugly and, damnit, is it really necessary
anymore? I mean, sure, I had to [work around this issue in Bricolage] back in
2005, but on modern file systems like zfs and ext3, does it really matter
anymore?

I was chatting about this with [Schwern] just now, and he thought it just
didn't matter anymore. So I asked my fellow [PGX] associate [Jeff Frost] about
this, and he said, "Used to be anytime you went above 1000 entries in one
directory, things would start to slow down. Certainly reiserfs, xfs, jfs, and
zfs don't suffer that issue." But what about ext3?

I happen to have a box with ext3, so I tested it. Here's what I found. To stat
one file among 20,000, `time` says:

    real    0m0.005s
    user    0m0.000s
    sys     0m0.000s

Not bad. And on file among 200?

    real    0m0.009s
    user    0m0.000s
    sys     0m0.010s

Well, you can't get much closer than that. What about subdirectories? To stat
a file inside one of 20,000 subdirectories, `time` tells me:

    real    0m0.015s
    user    0m0.010s
    sys     0m0.000s

And a file inside one of 200 subdirectories:

    real    0m0.005s
    user    0m0.000s
    sys     0m0.000s

Well, I can live with that. I suppose there are some file systems out there
that still have this problem, but you know what? I'm not going to worry about
them. It's thinking too far in advance anyway (if PGXN has this kind of
scaling problem we'll be lucky!), and the farther out we get, the less of a
problem it is.

So you know what? I'm not going to use the hashing of extension, distribution,
and owner names. Let the file systems worry about that performance, not me.



  [organization of the mirror]: https://blog.pgxn.org/post/988613682/restful-directory
  [work around this issue in Bricolage]: https://github.com/bricoleurs/bricolage/commit/37c3ac85006503bde4240f642a315ff4c3fb425b
  [Schwern]: https://schwern.dreamhosters.com/
  [PGX]: https://pgexperts.com/
  [Jeff Frost]: https://www.facebook.com/frosty996
