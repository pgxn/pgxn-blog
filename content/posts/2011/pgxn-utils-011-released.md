---
title: PGXN Utils 0.1.1 Released!
slug: pgxn-utils-011-released
date: 2011-05-23T04:53:30Z
aliases: [/post/5758832725/pgxn-utils-011-released, /post/5758832725]
tags: [Build, Create Extension, Utils, Bundle Extension, Meta, README, Skeleton]
author:
  name: Dickson S. Guedes
---

Hello everyone!

This was a productive weekend that allowed me to work on some new features in
[pgxn_utils] and I'm proud to tell you that a new version was released!

Trying to simplify your extension-development life I've added two tasks to
`pgxn_utils`: `change` and `bundle`. The first one is just a convenient way to
change META information about extension, incrementally, the second one is an
easy way to archive your extension in a zip file well named.

To install it just type:

```sh
gem install pgxn_utils
```

Or, if you don't want to install it yet, see it in action on this [screencast]
3:05.

**Work in progress...**

I'm working now to simplify the release, creating a task to send bundled file
to [PGXN].

There are a lot of work to do yet so, please, [tell me] if you found a bug or
have suggestions.

Have a nice code! ":)

  [pgxn_utils]: https://github.com/guedes/pgxn-utils
  [screencast]: https://blip.tv/pgcasts/pgxn_utils-0-1-1-released-5194610
  [PGXN]: https://pgxn.org
  [tell me]: https://github.com/guedes/pgxn-utils/issues
