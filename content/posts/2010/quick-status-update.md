---
title: Quick Status Update
slug: quick-status-update
date: 2010-09-28T17:29:48Z
aliases: [/post/1205127396/quick-status-update, /post/1205127396]
tags: [Status Update, Manager, Contribute, Sponsor]
---

Sometimes I wish I weren't such a perfectionist.

But not often.

A quick update:

So far, I've spent 43 hours on the [PGXN Manager] database. I had estimated 24
hours. Ha ha ha ha ha!

I've spent 45 hours on the implementation of the app itself. I had estimated
40 hours.

How will I make up the difference? I'm not sure, really. I've already
under-recorded my hours quite a lot, basically donating my time to create
[SemVer] and to create the basic HTML layout for the site (I suck at design,
but found [a good template] to base it on). That will likely continue. This is
an OSS project, and while I'm getting paid to work on it thanks to our
[generous contributors], I'm also contributing quite a lot of time to it.

Speaking of contributors, we've still not quite met our fundraising goals.
There's enough there for me to finish PGXN Manager so that you can start
releasing your extensions on PGXN, and for me to start on the search and
documentation site, but not finish it. If you know any organizations that
would like to sponsor this work and get their name and link on the PGXN site
in perpetuity, please [send them over!].

*Ahem.*

So what have I been doing? I'll write up some notes in a few other blog posts,
including how I'm generating JSON in the database, using [Plack] for the app,
and nicely-degrading Ajaxification with [jQuery] and RESTful URLs (I hope!).
The main page of the app works, as does authentication via basic auth. There's
a screen to request a user account, and a UI for PGXN admins to accept or
reject such requests. To make it actually usable, I just need to add the
upload feature for users, and then I can do a first release. That will allow
people to sign up, get approved, and start uploading extensions for
distribution on PGXN.

I *will* get that that done this week.

After that, I'll add screens for users to edit their account information,
change passwords, reset passwords, and edit permissions. I expect each of
those to take less time, as they'll use a lot of the infrastructure I've
already built.

Watch this space, and thanks for your patience!

Oh, and if you want to help out, please do fork [PGXN::Manager][PGXN Manager]
and ping me in #pgxn on Freenode for the deets on getting it built.

  [PGXN Manager]: https://github.com/theory/pgxn-manager
    "PGXN Manager repository on GitHub"
  [SemVer]: https://search.cpan.org/perldoc?SemVer
  [a good template]: https://andreasviklund.com/templates/andreas07/
  [generous contributors]: https://pgxn.org/contributors.html
  [send them over!]: https://pgxn.org/
  [Plack]: https://plackperl.org/
  [jQuery]: https://jquery.org/
