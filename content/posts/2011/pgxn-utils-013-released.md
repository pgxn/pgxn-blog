---
title: PGXN Utils 0.1.3 Released!
slug: pgxn-utils-013-released
date: 2011-09-08T06:39:00Z
aliases: [/post/9950473714/pgxn-utils-013-released, /post/9950473714]
tags: [PGXN Utils, Development, Utils, Build, Bundle Extension, Meta]
author:
  name: Dickson S. Guedes
---

Hello everyone!

I'm proud to tell you that a new version of [pgxn_utils] was released!

In this version some errors with OpenSSL was fixed (thanks @theory to report
then), and now you can release an extension to [PGXN][] [in **five steps**]
even using Ruby 1.8!

Another change was the executable's name that changed from `pgxn_utils` to
`pgxn-utils` for a close integration with next version of [PGXN Client], but
some work need to be done, yet.

I used `pgxn-utils` to release itself to [PGXN][1]!

```console
$ pgxn-utils release pgxn_utils-0.1.3.zip 
Enter your PGXN username: guedes
Enter your PGXN password: ***********************
Trying to release pgxn_utils-0.1.3.zip ... released successfully!
Visit: https://pgxn.org/dist/pgxn_utils/0.1.3/
```

Cool, eh? So, since the PGXN's mirrors were synced and you have `pgxn` client,
you could install `pgxn_utils` using:

```sh
pgxn install pgxn_utils
```

If you don't have `pgxn` client you can install it using rubygems

```sh
gem install pgxn_utils
```

Have fun!

  [pgxn_utils]: https://github.com/guedes/pgxn-utils
  [PGXN]: https://pgxn.org
  [in **five steps**]: https://blog.pgxn.org/post/6883009649/pgxn-utils-0-1-2-released
  [PGXN Client]: https://pgxnclient.projects.postgresql.org
  [1]: https://pgxn.org/dist/pgxn_utils/
