---
title: PGXN Utils
slug: pgxn-utils
date: 2011-05-14T01:05:00Z
aliases: [/post/5465631144/pgxn-utils, /post/5465631144]
tags: [Build, Makefile, Control File, README, Meta, Create Extension, Skeleton, Generator]
---

by Dickson S. Guedes

Do you ever have problems with copy and paste? I often did, and that is why I
create custom templates for often used files that match certain patterns.

With files from the structure of PostgreSQL's extensions was the same thing.

I was tired of creating the files and edit the META, controlfile, READMEs,
etc. every time I start a new extension and felt that I need something that
made me more productive, so I decided to create an automatic generator and
share it with the world.

I called it [pgxn-utils], and you should give it a try: it is easy to install,
easy to use and will help you to start hacking quickly!

**How?**

1.  First install it:

    ```sh
    gem install pgxn_utils
    ```

2.  Then start a new extension:

    ```sh
    pgxn_utils skeleton my_cool_extension
    ```

Thats all! It will create the initial skeleton for you and you can start
coding! But, if you don't want to install it, [see it in action]

Good hack!

  [pgxn-utils]: https://github.com/guedes/pgxn-utils/
  [see it in action]: https://pgcasts.com/media/pgxn_utils-usage-example.mpeg
