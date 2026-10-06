---
title: New HOWTO
slug: new-howto
date: 2011-05-13T20:34:46Z
aliases: [/post/5458118596/new-howtwo, /post/5458118596]
tags: [HOWTO, Build, Makefile, README, Changes, Documentation, META.json, Control File]
---

I updated the [howto] yesterday. This document explains how to create a PGXN
distribution. If you're interested in releasing PostgreSQL extensions on
[PGXN], this document is worth a read.

In essence, it's really simple: Just create a [`META.json`] and upload. But to
get the full benefit, there are quite a few other recommendations. Already
familiar with it? Here's the checklist:

- Create a [`META.json`]
- Create a [control file]
- Create a [`Makefile`]
- Implement the code in the `sql` and `src` directories
- Write tests in the `test` directory
- Write a [Text::Markup]-recognizable `README`
- Write [Text::Markup]-recognizable documentation in the `doc` directory
- Consider including other files: `Changes`, `LICENSE`, `INSTALL`, `COPYING`,
  `AUTHORS`
- [Release it][howto]!

Be sure to read the [howto] for details. Got feedback or suggestions? Leave a
comment!

Oh, and check out [pgxn-utils] and simplify your extension-development life.



  [howto]: https://manager.pgxn.org/
  [PGXN]: https://pgxn.org/
  [`META.json`]: https://pgxn.org/spec/
  [control file]: https://www.postgresql.org/docs/9.1/static/extend-extensions.html
  [`Makefile`]: https://www.postgresql.org/docs/current/static/xfunc-c.html#XFUNC-C-PGXS
  [Text::Markup]: https://search.cpan.org/perldoc?Text::Markup
  [pgxn-utils]: https://github.com/guedes/pgxn-utils/
