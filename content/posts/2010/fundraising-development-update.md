---
title: Fundraising and Development Update
slug: fundraising-development-update
date: 2010-10-19T18:43:57Z
aliases: [/post/1352594371/fundraising-development-update, /post/1352594371]
tags: [Fundraising, Development, PGXN, Contribute, Perl, Database, PgWest, Conference]
---

Yesterday was a busy day. In addition to making the [first PGXN release], I
updated the fundraising spreadsheet and then the thermometer displayed on the
[main site]. The good news is that things are coming along nicely. Thanks to
recent contributions from [Command Prompt], [Marchex], Hitoshi Harada, and
[25th-floor], we are now just $2500 short of our goal of $25,000. Thank you
all!

Can you help us get to our goal in time for [PgWest 2010], which is November
2-4 in San Francisco? I'll be giving a talk there, "[Building and Distributing
PostgreSQL Extensions Without Learning C]", in which PGXN will of course be
featured. Would be great to announce that the fundraising was successful.

As for the time I've put in so far, I'm happy to have PGXN Manager up and
working, but of course it has taken more hours than I expected. 76.5 so far.
I'd estimated 40. Meanwhile, the database design is up to 43 hours from the
estimated 24. And that doesn't count the hours I spent chasing shiny yaks and
shaving them, like [SemVer] and [Router::Resource]. Those of you who estimate
development projects, take heed! I clearly need to double all my estimates
before I submit them.

Still, with the fundraising nearly done, I'm committed to finishing this
project. I view it as a project budget, and so that's what it will be, whether
or not it takes me twice or four times as many hours as I'd estimated.

That said, you could help. Right? PGXN Manager is in good shape, but it's not
done. If you'd like to roll up your sleeves and contribute some code, please
[fork it], build it, and hack what you can. A few things on the to-do list:

- Add user admin interface. Should include checkbox to make and remove
  administrative permission
- Add mirror admin interface
- Add extension ownership permissions interface

The database API is there for these bits already; the code would mainly be in
Perl. Hit me on #pgxn on [Freenode] if you want to help.

If documentation is your thing, contributions there would be appreciated, as
well. In particular, the [About PGXN] page is a bit thin. Other interfaces
will need help, too. More on that as we add users.

Thanks everyone for your support!

  [first PGXN release]: https://blog.pgxn.org/post/1352326020/first-upload
  [main site]: https://pgxn.org/contributors.html
  [Command Prompt]: https://www.commandprompt.com/
  [Marchex]: https://www.marchex.com/
  [25th-floor]: https://www.25th-floor.com/
  [PgWest 2010]: https://www.postgresqlconference.org/2010/west/
  [Building and Distributing PostgreSQL Extensions Without Learning C]: https://www.postgresqlconference.org/content/building-and-distributing-postgresql-extensions-without-learning-c
  [SemVer]: https://search.cpan.org/perldoc?SemVer
  [Router::Resource]: https://search.cpan.org/perldoc?Router::Resource
  [fork it]: https://github.com/theory/pgxn-manager/
  [Freenode]: https://webchat.freenode.net/
  [About PGXN]: https://manager.pgxn.org/about
