---
title: Writing a client for PGXN
slug: writing-a-client-for-pgxn
date: 2011-04-29T00:15:00Z
aliases: [/post/5026314153/writing-a-client-for-pgxn, /post/5026314153]
tags: []
---

By Daniele Varrazzo.

PostgreSQL 9.1 is quickly rolling to the beta phase, and the release will
bring a new useful feature: [extensions]. Well, actually extensions have
always existed, but they didn't have an identity before: what you used to get
was a set of loose objects and mysterious functions in your namespace,
typically getting in the way in your dump and restore. Now here it is: [CREATE
EXTENSION], in all its uppercase glory!

So, an extension will exist in your database, and this finally gives the
possibility to organize them, this is what [PGXN] is for: a repository for
metadata, documentation and code, neatly packaged and asking to be picked off
the shelf. I've recently [written an extension] myself so I got some original
material to play with PGXN, which has resulted in quite a pleasant experience:
[the infrastructure] is well done thanks to David's design, and [the API] is
ridiculously easy to interact with: I've literally explored it using curl.

There was a big piece missing in the jigsaw that's being composed: a client to
make easy for database developers and administrators to download, build and
install extensions, and with PGXN getting more mature and people starting
packaging their work it was really *begging* to be written! So here it is:
~~pgxn.client~~ [pgxnclient], a command line tool to interact with PGXN. There
is still a lot to be developed, but the idea is to release early, get feedback
for it and improve it quickly to complete the extensions picture. So it's
already [available on PyPI] and very easy to start with it: just use

```console
$ sudo easy_install pgxnclient
```

and in a few seconds you will have a script called `pgxn`, offering several
commands that can be displayed with `pgxn --help` (and currently this is the
most up-to-date documentation you can get, as new commands are getting added
very quickly).

Were you looking for a solution to store hashes? You may try:

```console
$ pgxn search hash
sha 1.0.0
session_hash_tools 1.0.0
semver 0.2.1
```

Uhm, is `sha` a possible solution?

```console
$ pgxn info sha
INFO: best version: sha 1.0.0
name: sha
abstract: This module provides datatypes for storing SHA-1,
SHA-2 and MD5 hashes
maintainer: Alexey Klyukin <a...@commandprompt.com>
license: postgresql
release_status: stable
version: 1.0.0
date: 2011-03-16T10:33:00Z
sha1: 0187b0d261d302605bf8d0a15cdbd809deb245dd
provides: sha: 1.0.0
```

Let's say it is exactly what we were looking for (we could also display the
readme with `pgxn info --readme`). Shall we give it a try?

```console
$ sudo pgxn install sha
INFO: best version: sha 1.0.0
INFO: saving /tmp/tmpj8G6kM/sha-1.0.0.zip
INFO: unpacking: /tmp/tmpj8G6kM/sha-1.0.0.zip
INFO: building extension
[some compiler log]
INFO: installing extension
[files being copied]
```

Now the code is in the right place in the database directory (a specific
`pg_config` can be specified to choose which one): If you want it in a
specific database:

```console
$ pgxn load -U postgres -d test sha
```

This will result in CREATE EXTENSION being invoked, if the target database
supports it, or in the loading of the provided sql file for PostgreSQL
versions up to 9.0.

Easy, wasn't it?

There are still a lot of features to add:

- drop extension from a database and uninstall from the system
- select the target schema
- work with locally downloaded packages

but as `pgxnclient` uses the PGXN server on a side and the [PGXS build
infrastructure] on the other, the features will be easy to add and the client
work is actually an easy one.

Give it a try if you want: your feedback is very welcome. Meanwhile we will
keep on adding features to cover the entire extensions life cycle.

**Edit**: the program has been [renamed to pgxnclient][]: the relevant URLs
have been updated.

  [extensions]: https://developer.postgresql.org/pgdocs/postgres/extend-extensions.html
  [CREATE EXTENSION]: https://developer.postgresql.org/pgdocs/postgres/sql-createextension.html
  [PGXN]: https://pgxn.org/
  [written an extension]: https://pgmp.projects.postgresql.org/
  [the infrastructure]: https://manager.pgxn.org/
  [the API]: https://github.com/pgxn/pgxn-api/wiki/index-api
  [pgxnclient]: https://github.com/dvarrazzo/pgxnclient
  [available on PyPI]: https://pypi.python.org/pypi/pgxnclient/
  [PGXS build infrastructure]: https://developer.postgresql.org/pgdocs/postgres/extend-pgxs.html
  [renamed to pgxnclient]: https://blog.pgxn.org/post/5118152273/new-release-for-the-pgxn-client
