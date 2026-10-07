---
title: PGXN Utils 0.1.2 Released!
slug: pgxn-utils-012-released
date: 2011-06-24T23:53:00Z
aliases: [/post/6883009649/pgxn-utils-012-released, /post/6883009649]
tags: [Build, Create Extension, Utils, Meta, Release, Bundle]
author:
  name: Dickson S. Guedes
---

Hello everyone!

I'm proud to tell you that a new version of [pgxn_utils] was released!

Now you can release a distribution to [PGXN] in **five steps**!

**First**, install it:

```sh
gem install pgxn_utils
```
**Second**, create your extension, optionally overwrite some defaults:

```sh
mkdir $HOME/extensions
cd $HOME/extensions
pgxn_utils skeleton my_extension --maintainer "Dickson S. Guedes"
```

**Third**, code!

**Fourth**, bundle it:

```console
pgxn_utils bundle my_extension
Extension generated at: /home/guedes/extensions/my_extension-0.0.1.zip
```

**Fifth**, release it:

```console
pgxn_utils release my_extension-0.0.1.zip
Enter your PGXN username: guedes
Enter your PGXN password: ******
Trying to release my_cool_extension-0.0.1.zip ... released successfully!
Visit: https://manager.pgxn.org/distributions/my_cool_extension/0.0.1
```

Ah, you can export `PGXN_USER` and `PGXN_PASSWORD` if you are tired to type
your username and password every time.

[Check this out!][pgxn_utils].

  [pgxn_utils]: https://github.com/guedes/pgxn-utils
  [PGXN]: https://pgxn.org
