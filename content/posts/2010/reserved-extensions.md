---
title: Reserved Extensions
slug: reserved-extensions
date: 2010-10-21T20:43:00Z
aliases: [/post/1368261274/reserved-extensions, /post/1368261274]
tags: [Reserved Words, Reserved Extensions, Extensions, PostgreSQL, Procedural Languages, PL/pgSQL, PL/Perl, Contributed Modules]
---

I'm thinking about how to add support for reserved extensions. These are
extensions that one needs to depend on, but aren't distributed via PGXN.
Primarily, this means stuff distributed with the PostgreSQL core, including:

- PostgreSQL itself
- Bundled procedural languages (plpgsql, plperl, etc.)
- All [contributed modules]

Basically, anything that an extension might want to declare as a dependency,
but that isn't on PGXN itself.

There are a number of ways to do this. Which do you think would be the best
approach?

1.  Hard-code a list into the source code
2.  Include an editable list in the runtime configuration file
3.  Add a database table reserved for them and an API to edit it
4.  Create a "pgdev" user and upload a distribution declaring all those
    extensions in a `META.json` file
5.  Same as 4, but actually upload the PostgreSQL source

I'm leaning towards #2, perhaps having it automatically maintain a list in the
database and a metdata file on the mirrors.

But what do you think? Opinions wanted!



  [contributed modules]: https://www.postgresql.org/docs/current/static/contrib.html
