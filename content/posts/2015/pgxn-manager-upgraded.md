---
title: PGXN Manager Upgraded
slug: pgxn-manager-upgraded
date: 2015-08-31T21:48:54Z
aliases: [/post/128057409023/pgxn-manager-upgraded, /post/128057409023]
tags: [Release, Upgrade, Versions]
---

I took a little time this summer to finally address some nagging issues in
[PGXN Manager], the site to which extensions are uploaded and added to the
[master repository]. Most of the changes were internal, improving the
interface through which I sometimes have to re-index a release. There are also
a couple of minor changes to the sample `Makefile` in the [How To]. But the
most important change for new releases going forward is that version ordering
is now enforced. That means two things:

- Distribution versions **must** be greater than the version of the previous
  release. You can no longer release `v1.2.0` today and `v1.1.0` tomorrow.
- Extension versions **must** be greater than or equal to versions in previous
  releases. It's pretty common to have a new release version but have embedded
  extensions be the same version as before. But they can't be any less than
  before.

The changes have been applied to the reindexing code, as well, which also
prevents versions from being greater than in subsequent releases. That won't
come up very often---most of the time, I end up reindexing something that has
just been released. But in the future I expect to add an interface for release
managers to [reindex their own distributions], so it may be that someone
reindexes a release that's a few versions old. Or not. But just in case, it'll
be handled.

So what are you going to release on PGXN today? [Get to it!][How To].



  [PGXN Manager]: https://manager.pgxn.org/
  [master repository]: https://master.pgxn.org/
  [How To]: https://manager.pgxn.org/howto
  [reindex their own distributions]: https://github.com/pgxn/pgxn-manager/issues/49
