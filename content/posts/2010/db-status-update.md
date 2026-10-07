---
title: "Status Update: DB API, Extension Versions RFC"
slug: db-status-update
date: 2010-09-07T18:46:38Z
aliases: [/post/1082188310/db-status-update, /post/1082188310]
tags: [Status Update, Database, API, Documentation, Extensions, Versions, RFC]
---

I spent most of the time I had to work on PGXN the last two weeks creating the
database for [PGXN Manager]. I had estimated 24 hours of work to design the
database. So far I've logged 34 hours. But I think that will come out in the
wash, really, because I did a lot of work to generate JSON from database
functions. This is code I had originally expected to do in the app layer. I'm
starting work on that this week, with an estimated 40 hours. Hope I can do it
in 30. :-)

I'm pretty happy with how the database API is turning out. See the nifty
[database API documentation] I whipped up using [gendoc] and embedded
[MultiMarkdown]. There are still a few tweaks I need to make. I just checked
in a change to eliminate all the `ALIAS`es I [blogged about] last week. Turns
out that one can just [function-name-qualify] the function parameter names.
Who knew? I sure didn't.

But I do have a question for you, dear readers. Right now, the database
requires that extensions have unique version numbers in every distribution. So
if, for example, you uploaded the distribution foo 1.2.2 with the extension
bar 1.2.2, when you next uploaded foo 1.2.3, bar could not be 1.2.2. I
designed this this way by following my own practice of always incrementing the
version numbers of all modules included in my [CPAN distributions]. So all
modules have unique version numbers, with never a duplicate.

However, CPAN itself only cares that distributions have unique version
numbers. It doesn't care about the version numbers of included modules. See,
for example, [HTML::Mason]. Note that the various included modules have all
sorts of different version numbers, and many have no version numbers at all.
So while the core module has a version number (1.45 at the time of this
writing), if you click to look at older versions, say [Mason 1.44], you'll see
that no other modules have their version numbers changed.

I don't think I want to allow extensions without version numbers on PGXN.
That's been a recipe for many annoyances with CPAN. But maybe I should allow
extension version numbers to repeat? The upside is less maintenance for
multi-extension distribution authors. The downside is potentially less
accuracy in the determination of prerequisites.

For example, say that we have two versions of distribution foo:

```
  -----------------------------------------------------------------------------
  Distribution                           Extensions
  -------------------------------------- --------------------------------------
  foo 1.2.2                              foo 1.2.2\
                                         bar 1.2.2

  foo 1.2.3                              foo 1.2.3\
                                         bar 1.2.2
  -----------------------------------------------------------------------------
```

Note how the version number of extension bar has not changed between releases.
But if I was a user of both foo and bar, and I specified that I required bar
1.2.2 but was actually using stuff in foo 1.2.3, I could end up with errors
because I would only have foo 1.2.3.

This is an issue periodically faced by Perl hackers. I might require
HTML::Mason::ApacheHandler 1.69 but really what I need is HTML::Mason 1.44. Of
course, as a Perl hacker, it's my responsibility to make sure I require
exactly what I need, but sometimes it's not clear what's the primary module in
a distribution. And if I don't realize that the version number of
HTML::Mason::ApacheHandler doesn't change with every release, I might make
mistakes.

Maybe that's okay. Maybe PGXN should be less rigid like this. But I'm leaning
toward keeping things as they are and requiring new versions of every
extension in every upload of a distribution. It's a bit tougher for those
(few?) hackers who will create multi-extension distributions, but probably
more reliable for everyone else.

What do you think?

Anyway, I'm getting to work on the Web app this week while this issue
percolates. Been studying up on [Plack]. So nice!

  [PGXN Manager]: https://github.com/theory/pgxn-manager/
  [database API documentation]: https://github.com/theory/pgxn-manager/wiki/DB-API
  [gendoc]: https://github.com/theory/pgxn-manager/blob/master/bin/gendoc
  [MultiMarkdown]: https://fletcherpenney.net/multimarkdown/users_guide/multimarkdown_syntax_guide/
  [blogged about]: https://blog.pgxn.org/post/1053165383/alias-in-vogue
  [function-name-qualify]: https://archives.postgresql.org/pgsql-hackers/2010-09/msg00404.php
  [CPAN distributions]: https://search.cpan.org/~dwheeler/
  [HTML::Mason]: https://search.cpan.org/dist/HTML-Mason/
  [Mason 1.44]: https://search.cpan.org/~drolsky/HTML-Mason-1.44/
  [Plack]: https://plackperl.org/
