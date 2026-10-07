---
title: Extension Makefiles for PostgreSQL 9.1
slug: extension-makefiles
date: 2011-04-20T19:29:00Z
aliases: [/post/4783001135/extension-makefiles, /post/4783001135]
tags: [Extension, Makefile, make, Pgxs, pg_config]
---

In order to keep distribution packaging as simple as possible, I worked up
this `Makefile` some time ago:

```makefile
DATA = $(wildcard sql/*.sql)
DOCS = $(wildcard doc/*.txt)
TESTS = $(wildcard test/sql/*.sql)
REGRESS = $(patsubst test/sql/%.sql,%,$(TESTS))
REGRESS_OPTS = --inputdir=test --load-language=plpgsql
MODULES = $(patsubst %.c,%,$(wildcard src/*.c))

PG_CONFIG = pg_config
PGXS := $(shell $(PG_CONFIG) --pgxs)
include $(PGXS)
```

The nice thing about this code is that it has nothing specific to a
distribution in it. It figures out what SQL files there are, what doc files
there are, and what C files need compiling by just looking in the `sql`,
`doc`, and `src` directories, respectively. It also specifies that tests are
in the `test` directory. About the only thing I've customized here is adding
`--load-language=plpgsql` to `REGRESS_OPTS`, as the tests for the distribution
I've copied this from require PL/pgSQL to run. Simple, and anyone can use it
with very little need to tweak it, as long as they don't mind storing their
files in the specified directories.

Today, I'm updating my distributions to support PostgreSQL 9.1's new
`CREATE EXTENSION` syntax, but I want to continue supporting older versions of
PostgreSQL, as well. Basically, this means that the files listed in the `DATA`
variable vary based on the version of PostgreSQL you're installing against.
Here are the additional things the `Makefile` needs to do:

- If installing against PostgreSQL less than version 9.1, exclude files in
  `DATA` that contain `--`. Such files are are migration scripts, which aren't
  supported before 9.1.
- If installing against PostgreSQL greater than or equal to 9.1:
  - Copy the files `sql/$EXTENSION.sql` and `sql/$EXTENSION--unpackaged.sql`
    to `sql/$EXTENSION--$EXTVERSION.sql` and
    `sql/$EXTENSION--npackated--$EXTVERSION.sql`, respectively.
    `CREATE EXTENSION` requires that the version string be in migration file
    name. I'd rather not have to rename the file in my repository before every
    release (and I'd rather keep it without the version for  9.1 anyway), so
    it needs to be copied.
  - Add the new `sql/$EXTENSION--$VERSION.sql` file to `EXTRA_CLEAN`.
  - Include only files in `DATA` that contain `--`. There's no need to install
    the original file without the version number, or any uninstall file,
    either, since it's not needed on 9.1 anymore.

I've been trying to figure out how to modify my standard `Makefile` to support
these changes, without requiring a lot of tweaking, so that other folks can
easily use it in the future. Thanks to help from [Andrew Dunstan], this is
what I've come up with:

```makefile
EXTENSION=semver
EXTVERSION=0.2.2

DATA = $(filter-out $(wildcard sql/*--*.sql),$(wildcard sql/*.sql))
DOCS = $(wildcard doc/*.txt)
TESTS = $(wildcard test/sql/*.sql)
REGRESS = $(patsubst test/sql/%.sql,%,$(TESTS))
REGRESS_OPTS = --inputdir=test --load-language=plpgsql
MODULES = $(patsubst %.c,%,$(wildcard src/*.c))

PG_CONFIG = pg_config

VERSION     = $(shell $(PG_CONFIG) --version | awk '{print $$2}')
PGVER_MAJOR = $(shell echo $(VERSION) | awk -F. '{ print ($$1 + 0) }')
PGVER_MINOR = $(shell echo $(VERSION) | awk -F. '{ print ($$2 + 0) }')

ifeq ($(PGVER_MAJOR), 9)
ifneq ($(PGVER_MINOR), 0)
all: sql/$(EXTENSION)--$(EXTVERSION).sql sql/$(EXTENSION)--unpackaged--$(EXTVERSION).sql

sql/$(EXTENSION)--$(EXTVERSION).sql: sql/$(EXTENSION).sql
    cp $< $@

sql/$(EXTENSION)--unpackaged--$(EXTVERSION).sql: sql/$(EXTENSION)--unpackaged.sql
    cp $< $@

DATA = $(filter-out sql/$(EXTENSION)--unpackaged.sql,$(wildcard sql/*--*.sql)) sql/$(EXTENSION)--$(EXTVERSION).sql
EXTRA_CLEAN = sql/$(EXTENSION)--$(EXTVERSION).sql sql/$(EXTENSION)--unpackaged--$(EXTVERSION).sql
endif
endif

PGXS := $(shell $(PG_CONFIG) --pgxs)
include $(PGXS)
```

This is not exactly ideal, but not *too* bad. It's not quite the drop-in
version we had before, because now the first line needs to name the extension
we're distributing, and the second needs to specify the version (and would
then need to be updated for every release). Maybe they could be read from the
control file somehow? Other than that, you should be able to just forget the
rest of the file (mostly). Here's how it addresses the above requirements:

- To exclude files with `--` in them on  9.1, the first `DATA` line filters
  them out:

```makefile
DATA = $(filter-out $(wildcard sql/*--*.sql),$(wildcard sql/*.sql))
```
- Next, we need to know if we're on 9.1 or higher. So we use
  `pg_config --version` to get the version number and some `awk` stuff to get
  the major and minor parts. Then, if the major version is 9 and the minor is
  *not* 0, we:

  - Add `sql/$(EXTENSION)--$(EXTVERSION).sql` and
    `sql/$(EXTENSION)--unpackaged--$(EXTVERSION).sql` as dependencies of the
    `all` rule (which is the default PGXS rule).
  - Add the `sql/$(EXTENSION)--$(EXTVERSION).sql` and
    `sql/$(EXTENSION)--unpackated--$(EXTVERSION).sql` rules, which copy
    `sql/$(EXTVERVERSION).sql` and `sql/$(EXTVERVERSION)--unpackaged.sql`
    files. Of course this assumes that such files exist.
  - Add `sql/$(EXTENSION)--$(EXTVERSION).sql` and
    `sql/$(EXTENSION)--unpackaged--$(EXTVERSION).sql` to `EXTRA_CLEAN`, so
    that they'll be deleted by `make clean`.
  - Set `DATA` again, this time to include *only* files with `--` in them,
    except for `sql/$(EXTENSION)--unpackaged.sql`.

And with that, it works. But it has some disadvantages over the previous, very
simple `Makefile` I've been using up to now:

- One must modify it for every release, if only to set `EXTVERSION`.
- It relies more on Unix tools than previously, specifically `awk`. I'm not
  sure how big a deal this is in practice; the [pgTAP Makefile] has been
  relying on `awk` an even worse gymnastics for some time and no one has
  complained.
- It's not perfectly future-proof. It will need to be modified when PostgreSQL
  10.0.0 is released, and again when 11.0.0 is released.
- It won't handle a distribution with multiple extension in it. I don't even
  want to think about that.

Suggestions for ways to eliminate these shortcomings would be greatly
appreciated, especially if it allows use extension authors to get back to
something as simple as the first example at the top of this post.

**UPDATE:** Added the `unpackaged` bits I didn't realize I needed until after
I'd released a new version of [semver] and discovered that the "unpackaged"
script needs to always be tied to the default version.

  [Andrew Dunstan]: https://people.planetpostgresql.org/andrew/
  [pgTAP Makefile]: https://github.com/theory/pgtap/blob/master/Makefile
  [semver]: https://pgxn.org/dist/semver/
