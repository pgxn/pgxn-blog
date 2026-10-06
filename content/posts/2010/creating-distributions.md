---
title: Creating Distributions
slug: creating-distributions
date: 2010-10-19T22:28:00Z
aliases: [/post/1353947079/creating-distributions, /post/1353947079]
tags: [Distribution, Makefile, Pgxn, Pgxs, pg_config, pg_regress, Directory Structure]
---

Following the [upload of `pair`] to PGXN, I wanted to take a few minutes to
write about how to structure a PGXN distribution.

## OMG Distribution WTF?

First of all, what is a "distribution" in the PGXN sense? Basically, it's a
collection of one or more PostgreSQL extensions. That's it.

So why allow more than one extension? Maybe no PGXN distribution will ever
have more than one extension. After all, the goal should be many focused,
minimalist tools that folks can combine in their apps. But sometimes it
doesn't work out that way.

As an example, I've been planning to break [pgTAP] up into two parts for a
while: one for scalar and relational testing, the other for schema testing.
Often one needs only the scalar and relational testing, while the schema
testing is more often needed only for testing replication and whatnot. Whether
or not I choose to distribute both parts in one package I have yet to
determine, but it could well make sense to keep them in one distribution.

Besides, I've tried to write [PGXN::Manager] in such a way that it's not
specific to PostgreSQL. So that if someone wanted to create a Drupal XN with
it or something, they could. Or PyXN. Or, hell, even if [CPAN] wanted to
switch someday, they could. (Note that I've registered myxn.org for fun. I may
or may not do anything with that, but see [MyTAP]. Yes, I am insane.)

## That's So Meta

Anyway, back to the structure of distributions. At its simplest, the only
thing PGXN requires is a single file, `META.json`, which describes the
package. This is (currently) the only file that PGXN Manager uses to index a
distribution, so it's important to get it right. The [PGXN Meta Spec] has a
rather complete example of a hypothetical pgTAP distribution `META.json`.

If you have only one .sql file for your extension and it's the same name as
the distribution (and I expect this would be common), then you can make it
pretty simple. For example, the `pair` distribution has only one SQL file. So
the `META.json` could be:

    {
       "name": "pair",
       "abstract": "A key/value pair data type",
       "version": "0.1.0",
       "maintainer": "David E. Wheeler <david@justatheory.com>",
       "license": "postgresql",
       "meta-spec": {
          "version": "1.0.0",
          "url": "https://pgxn.org/meta/spec.txt"
       },
    }

That's it. The only thing that may not be obvious from this example is that
all version numbers in a `META.json` *must* be [semantic versions]. If they're
not, PGXN will make them so. So "1.2" will become "1.2.0", and so would
"1.02". So do try to use semantic versions and not worry about it.

In the short run, you won't need anything more in your `META.json` file. But
once I get to creating the search site and the command-line client for PGXN,
you're probably going to want to do more. Other useful keys to include are:

- [`tags`][]: An array of tags to associate with a distribution. Will help
  with searching.
- [`prereqs`][]: A list of prerequisite extensions or PostgreSQL contrib
  moules (or PostgreSQL itself).
- [`provides`][]: A list of included extensions. Useful if you have more than
  one or the one has a different name that the distribution (silly, but it
  happens). It also will index such extension names such that you are the
  owner, if you're the first to update one with that name.
- [`release_status`][]: To label a distribution as "stable," "unstable," or
  "testing." Useful for uploading distributions for people to test but that
  clients won't install by default.
- [`resources`][]: A list of related links, such as to an SCM repository or
  bug tracker. The search site will output these links.

Have a look at [the `META.json` in the `pair` distribution] for a more
extended example.

## New Order

For PGXN, the general idea is that you'll use [PGXS] to create your PostgreSQL
extensions. I'm hoping to encourage a slight modification of the directory
layout for PGXN distributions, but as I hope I've made clear so far, PGXN
itself doesn't really care how you structure things, or if you use PGXS. That
said, the proposed [download and installation client] will assume the use of
PGXS (unless and until the PostgreSQL core adds some other kind of
extension-building support), so it's probably the best choice.

Most PGXS-powered distributions have the code files in the main directory,
with documentation in a `README.extension_name` file. What I'd like to see
instead, and will encourage via the forthcoming [search site], is that things
be organized into subdirectories:

- `src` for any C source code files
- `sql` for SQL source files. These usually are responsible for installing an
  extension into a database
- `doc` for documentation files (the search site will likely look there for
  Markdown, Textile, HTML, and other document formats)
- `test` for tests

I've tried to make the `pair` distribution a good [example of this]. To make
it all work, The [Makefile] is written like so:

    DATA = sql/pair.sql sql/uninstall_pair.sql
    TESTS = $(wildcard test/sql/*.sql)
    REGRESS = $(patsubst test/sql/%.sql,%,$(TESTS))
    REGRESS_OPTS = --inputdir=test
    DOCS = doc/pair.txt

    ifdef NO_PGXS
    top_builddir = ../..
    include $(top_builddir)/src/Makefile.global
    include $(top_srcdir)/contrib/contrib-global.mk
    else
    PG_CONFIG = pg_config
    PGXS := $(shell $(PG_CONFIG) --pgxs)
    include $(PGXS)
    endif

The `DATA` variable identifies the files containing the extension, while
`TESTS` loads a list of all the tests, which are in the `test/sql` directory.
Note that I'm using `pg_regress` for tests. It expects that tests be named and
that there be corresponding "expected" files to compare against. With the
`REGRESS_OPTS = --inputdir=test` line, I'm telling `pg_regess` to find the
test files in [`test/sql`] and the expected output files in [`test/expected`].
And finally, the `DOCS` variable points to a single file with the
documentation, [`doc/pair.txt`]. If this extension had required any C code
(like [pgTAP] or [PostGIS] do), I would have pointed the `MODULES` variable at
files in a `src` directory.

After that we just have build instructions. If called with `make NO_PGXS=1`,
it assumes that the unzipped distribution directory has been put in the
"contrib" directory of the PostgreSQL source tree used to build PostgreSQL.
That's probably only important if one is installing on PostgreSQL 8.1 or
lower. Otherwise, it assumes a plain `make` and uses the [`pg_config`] in your
path to find PGXS to do the build.

For more on PostgreSQL extension building support, please consult [the
documentation][PGXS].

## Zip Me Up

Once you've got your extension developed and well-tested, and your
distribution just right and the `META.json` file all proof-read and solid,
it's time to upload the distribution to PGXN. What you want to do is to zip it
up to create a distribution archive. Here's what I did for `pair`, exporting
it from Git:

    git checkout-index -af --prefix ~/Desktop/pair-0.1.0/
    cd ~/Desktop/
    rm pair-0.1.0/.gitignore
    zip -r pair-0.1.0.zip pair-0.1.0

Then the `pair-0.1.0.zip` file was ready to upload. Simple, eh?

Now, one can upload any kind of archive file to PGXN, including a tarball, or
bzip2...um...ball? Basically, any kind of archive format recognized by
[Archive::Extract]. You can upload a `.pgz` if you like, in which case PGXN
will assume that it's a zip file. A zip file is best because then
PGXN::Manager won't have to rewrite it. It's also preferable that everything
be unpacked from an archive into a directory with the name
`$distribution-$version`. If not, PGXN will rewrite it to do so. But it saves
the server some effort if all it has to do is move a .zip file that's properly
formatted, so it would be appreciated if you would upload stuff that's already
nicely formatted for distribution in a zip archive.

## Release It!

And that's it! Not too bad, eh? Just please do be very careful cutting and
pasting examples; I initially uploaded the `pair` distribution thinking that
it contained pgTAP. It was kind of a PITA to fix. Hopefully we'll be able to
build things up to the point where a lot of this stuff can be automated
(especially the creation of the `META.json`), but for now it's done by hand.
So be careful out there, and good luck!

Oh, and if you have an extension that you'd like to release on PGXN now, I am
running a limited beta for interested extension developers. Please hit the
[mail list] for the details to be posted shortly.



  [upload of `pair`]: https://blog.pgxn.org/post/1352326020/first-upload
  [pgTAP]: https://pgtap.org/
  [PGXN::Manager]: https://github.com/theory/pgxn-manager
  [CPAN]: https://cpan/org/
  [MyTAP]: https://theory.github.com/mytap/
  [PGXN Meta Spec]: https://pgxn.org/meta/spec.html
  [semantic versions]: https://semver.org/
  [`tags`]: https://pgxn.org/meta/spec.html#tags
  [`prereqs`]: https://pgxn.org/meta/spec.html#prereqs
  [`provides`]: https://pgxn.org/meta/spec.html#provides
  [`release_status`]: https://pgxn.org/meta/spec.html#release_status
  [`resources`]: https://pgxn.org/meta/spec.html#resources
  [the `META.json` in the `pair` distribution]: https://github.com/theory/kv-pair/blob/master/META.json
  [PGXS]: https://www.postgresql.org/docs/9/static/xfunc-c.html#XFUNC-C-PGXS
  [download and installation client]: https://wiki.postgresql.org/wiki/PGXN#PGXN_Client
  [search site]: https://wiki.postgresql.org/wiki/PGXN#Search_Site
  [example of this]: https://github.com/theory/kv-pair/blob/
  [Makefile]: https://github.com/theory/kv-pair/blob/master/Makefile
  [`test/sql`]: https://github.com/theory/kv-pair/tree/master/test/sql/
  [`test/expected`]: https://github.com/theory/kv-pair/tree/master/test/expected/
  [`doc/pair.txt`]: https://github.com/theory/kv-pair/blob/master/doc/pair.txt
  [PostGIS]: https://postgis.org/
  [`pg_config`]: https://www.postgresql.org/docs/current/static/app-pgconfig.html
  [Archive::Extract]: https://search.cpan.org/perldoc?Archive::Extract
  [mail list]: https://groups.google.com/group/pgxn-users/
