---
title: PGXN Client 1.0 Released!
slug: pgxn-client-10-released
date: 2011-11-28T00:34:34Z
aliases: [/post/13427961249/pgxn-client-10-released, /post/13427961249]
tags: [PGXN Client, Release]
author:
  name: Daniele Varrazzo
---

Finally, here it is. Well tested, documented, and pampered. With the [PGXN
Client] installing extensions from the [PGXN Network] is a breeze:

```console
$ pgxn install semver
$ pgxn load semver
$ psql
=# select 'foo'::semver;
ERROR:  bad semver value 'foo': expected number at foo
LINE 1: select 'foo'::semver;
```

Error! Meaning: success!

The number of extensions on PGXN is steadily increasing, so we hope the client
will make their adoption even easier.

The client is now extensible: either writing in Python to reuse some of the
other commands implementation or writing new self-contained scripts in any
language to be invoked by the pgxn commands dispatcher. The first client
extensions are already available: [PGXN Utils] and [PGXN:Meta:Validator] are
targeted for easier development of new extensions.

The client is [released on PyPI], so installing is just:

```console
$ sudo easy_install pgxnclient
```

Complete documentation and further links are available from the [project
homepage][PGXN Client].

Any feedback is welcome; you can contact me and all the other people behind
PGXN on the [PGXN User Group]. See you there!

  [PGXN Client]: https://pgxnclient.projects.postgresql.org/
  [PGXN Network]: https://pgxn.org/
  [PGXN Utils]: https://github.com/guedes/pgxn-utils
  [PGXN:Meta:Validator]: https://github.com/pgxn/pgxn-meta-validator/
  [released on PyPI]: https://pypi.python.org/pypi/pgxnclient
  [PGXN User Group]: https://groups.google.com/group/pgxn-users/
