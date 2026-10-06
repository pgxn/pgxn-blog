---
title: PGXN Release Badges
slug: badges
date: 2015-04-11T04:34:08Z
aliases: [/post/116087351668/badges, /post/116087351668]
tags: [Badges, Gemfury, Travis CI, Coveralls, GitHub]
---

Looks like it's been close to two years since my last post on the PGXN blog.
Apologies for that. I've thought for a while maybe I should organize an
"extension of the week" series or something. Would there be interest in such a
thing?

Meanwhile, I'm finally getting back to posting to report on a fun thing you
can now do with your PGXN distributions. Thanks to the [Version Badge] service
from the nice folks at [Gemfury], you can badge your distributions! Badges
look like this:

[![PGXN version]][1]

You've no doubt seem similar badges for Ruby, Perl, and Python modules. Now
the fun comes to PGXN. Want in? Assuming you have a distribution named
`pgfoo`, just put code like this into the README file:

    [![PGXN version](https://badge.fury.io/pg/pgfoo.svg)](https://badge.fury.io/pg/pgfoo)

This is [Markdown] format; use the syntax appropriate to your preferred README
format to get the badge to show up on GitHub and PGXN.

That's it! The badge will show the current releases version on PGXN, and the
button will link through to PGXN.

Use [Travis CI]? You can badge your build status, too, as I've done for
[pgTAP], like this:

[![Build Status]][2]

    [![Build Status](https://travis-ci.org/theory/pgtap.png)](https://travis-ci.org/theory/pgtap)

[Coveralls] provides patches, too. I've used them for [Sqitch], though I've
not yet taken the time figure out how to do coverage testing with PostgreSQL
extensions. If you have, you can badge your current coverage like so:

[![Coverage Status]][3]

    [![Coverage Status](https://coveralls.io/repos/theory/sqitch/badge.svg)](https://coveralls.io/r/theory/sqitch)

So get badging, and show off your PGXN distributions GitHub and elsewhere!



  [Version Badge]: https://badge.fury.io
  [Gemfury]: https://gemfury.com/
  [PGXN version]: https://badge.fury.io/pg/semver.svg
  [1]: https://badge.fury.io/pg/semver
  [Markdown]: https://daringfireball.net/projects/markdown/
  [Travis CI]: https://travis-ci.org
  [pgTAP]: https://pgxn.org/dist/pgtap
  [Build Status]: https://travis-ci.org/theory/pgtap.png
  [2]: https://travis-ci.org/theory/pgtap
  [Coveralls]: https://coveralls.io
  [Sqitch]: https://github.com/theory/sqitch
  [Coverage Status]: https://coveralls.io/repos/theory/sqitch/badge.svg
  [3]: https://coveralls.io/r/theory/sqitch
