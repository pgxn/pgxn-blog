---
title: How To, Example Makefile Updated
slug: howto-updated
date: 2013-06-21T00:28:13Z
aliases: [/post/53477365497/howto-updated, /post/53477365497]
tags: []
---

I deployed a new version of [Manager] last week, and it included a major
update to the [How To]. A lot of the changes are narrative: I tried to keep
things shorter and more to-the-point. But from a technical point of view,
perhaps the most important improvements are in the `Makefile` example. Thanks
to ongoing work by [Cédric Villemain], it now tries harder to stay out of your
way, while integrating better with PostgreSQL's own `make` targets.

The main changes:

- The setting of `$EXTENSION` and `$EXTVERSION` is now done by reading the
  `META.JSON` file.
- Thanks to the `?=` operator, `$PG_CONFIG` can now be passed as a param as
  well as an environment variable. That is, you can do
  `PG_CONFIG=/my/pg_config make` or `make PG_CONFIG=/my/pg_config`.
- New targets are defined only after PGXS is loaded, so that targets can be
  overridden if necessary.
- A new `dist` target creates a PGXN-ready zip file, assuming that your code
  is maintained in Git.

Have a look at the [resulting diff for semver] to get an idea how you might
want to update your own `Makefile`s.

  [Manager]: https://manager.pgxn.org/
  [How To]: https://manager.pgxn.org/howto
  [Cédric Villemain]: https://www.linkedin.com/in/cedricvillemain
  [resulting diff for semver]: https://github.com/theory/pg-semver/compare/v0.3.0%E2%80%A6v0.4.0#diff-3
