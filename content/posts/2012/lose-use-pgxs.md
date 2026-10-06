---
title: Lose USE_PGXS in your Makefiles
slug: lose-use-pgxs
date: 2012-04-11T16:36:00Z
aliases: [/post/20908326485/lose-use-pgxs, /post/20908326485]
tags: [Extension, USE_PGXS, Makefile, make, pg_config, contrib]
---

Traditionally, folks have created extensions for PostgreSQL by copying one of
the [contrib modules] and hacking it into something new. One of the things
that comes along for the ride is the `Makefile` ([example]). As a result,
there are a lot of third-party extensions that use the `USE_PGXS` variable.

A bit of background. The core contrib extensions generally rely on a relative
path to include the core `Makefile`s needed to build the extension. Because
they ship with the core distribution, they can generally expect that the core
has already been compiled, the necessary `Makefile`s have been created, and
that they should be built against them. All the assumptions are that the
extensions should be built against the source tree in which they are
distributed. So there is no need to use [`pg_config`] to find [PGXS]; it
already knows where to find what it needs.

But as extensions, there is still the possibility that one might want to build
them against an existing installation of PostgreSQL, or an older version than
the source with which they're distributed. So the core hackers provided the
`USE_PGXS` variable so that one can in effect tell `make`, "Don't build
against the local source tree, but find PGXS for some other install and build
against that, instead." It was expected to be exceptional, since most folks
would build against the local source tree, and not a big deal to make anyone
else build it with:

    make USE_PGXS=1

Today things are different. There is a growing ecosystem of third party
extensions on [PGXN], [pgFoundry], [GitHub], and [Bitbucket], and obviously
they're not distributed with the PostgreSQL core. For these extensions, there
is no surrounding PostgreSQL source code to automatically include, so they
*must* use `pg_config` to find PGXS in order build.

Yet, there are quite a few third-party extensions that nevertheless assume
that they are in the `contrib` directory of the PostgreSQL source code
distribution, and so still have the `USE_PGXS` variable. The [twitter_ftw
1.0.0 `Makefile`] is a recent example. Just like core extensions, it has this
code:

    ifdef USE_PGXS
    PG_CONFIG = pg_config
    PGXS := $(shell $(PG_CONFIG) --pgxs)
    include $(PGXS)
    else
    subdir = contrib/twitter_fdw
    top_builddir = ../..
    include $(top_builddir)/src/Makefile.global
    include $(top_srcdir)/contrib/contrib-global.mk
    endif

Because Hitoshi-san originally copied the `Makefile` from a core extension, it
still assumes it will be distributed in core by default. And as I said, there
are quite a few third-party extensions that exhibit this pattern.

And now the [PSA][]: **Please don't use `USE_PGXS` in PostgreSQL extension
`Makefile`s.**

Not only is it unnecessary, it makes no sense for third-party extensions. They
should *always* assume that they need to use `pg_config` to find PGXS. If you
have an extension `Makefile` with `USE_PGXS` like twitter_ftw 1.0.0 did, you
should change it to something like this (as Hitoshi-san did in the
[twitter_ftw 1.0.1 `Makefile`]):

    PG_CONFIG = pg_config
    PGXS := $(shell $(PG_CONFIG) --pgxs)
    include $(PGXS)

That's it. I am asking you to *make your `Makefile` simpler.*

------------------------------------------------------------------------------

*An Aside*

There is *one* situation in which you might need to include the core contrib
`Makefile`s. And it's pretty unusual. If you need to support PostgreSQL 8.1 or
earlier, `pg_config` will not be able to tell you where to find PGXS. So users
will have to copy the extension source directory into the PostgreSQL source
`contrib/` directory and build from there. They will need a way to tell `make`
*not* to use PGXS. In this one unusual case, I suggest you add a `NO_PGXS`
variable. [pgTAP's `Makefile`] provides an example. But honestly, very few
extensions need to support PostgreSQL 8.1 (the oldest release currently
supported by the core hackers is *8.3!*), so make use of this pattern only if
absolutely necessary.

Otherwise, please don't use `USE_PGXS`.

------------------------------------------------------------------------------

If you want a complete guide to creating your extension `Makefile`, have a
look at the [PGXN Howto], which includes some detailed examples that include
support for pre- and post-[`CREATE EXTENSION`] support. The [PGXS] docs
contain additional details about all the `Makefile` variables you can use to
simplify extension configuration and installation. Check 'em out.



  [contrib modules]: https://www.postgresql.org/docs/current/static/contrib.html
  [example]: https://git.postgresql.org/gitweb/?p=postgresql.git;a=blob;f=contrib/isn/Makefile;hb=HEAD
  [`pg_config`]: https://www.postgresql.org/docs/current/static/app-pgconfig.html
  [PGXS]: https://www.postgresql.org/docs/current/static/extend-pgxs.html
  [PGXN]: https://pgxn.org/
  [pgFoundry]: https://pgfoundry.org/
  [GitHub]: https://github.com/
  [Bitbucket]: https://bitbucket.org/
  [twitter_ftw 1.0.0 `Makefile`]: https://api.pgxn.org/src/twitter_fdw/twitter_fdw-1.0.0/Makefile
  [PSA]: https://en.wikipedia.org/wiki/Public_service_announcement
  [twitter_ftw 1.0.1 `Makefile`]: https://api.pgxn.org/src/twitter_fdw/twitter_fdw-1.0.1/Makefile
  [pgTAP's `Makefile`]: https://api.pgxn.org/src/pgtap/pgtap-0.90.0/Makefile
  [PGXN Howto]: https://manager.pgxn.org/howto
  [`CREATE EXTENSION`]: https://www.postgresql.org/docs/current/static/sql-createextension.html
