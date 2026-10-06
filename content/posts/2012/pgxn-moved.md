---
title: PGXN Has a New Home
slug: pgxn-moved
date: 2012-01-12T04:51:04Z
aliases: [/post/15710159951/pgxn-moved, /post/15710159951]
tags: [Server, PostgreSQL, depesz, Community, Migration]
---

Day before yesterday, I finally got all of [PGXN] moved to a new server. I had
been using a small server owned by my company, [Kineticode], and hosted by
[Command Prompt]. That was fine for a while, but CMD was needing its rack
space back, and what with my [new job], I was shutting down Kineticode, too.
It was time to move PGXN elsewhere.

For a while, I got a lot of support and assistance towards moving PGXN to a
[PostgreSQL] community server. [Dave], [Magnus], and [Stefan] kindly spun up a
VM for me, and gave me permission to install Perl modules from [CPAN],
provided I supply them with a script to report to Nagios when Perl modules
were out of date, which of course I did. This was necessary because I built
PGXN with some pretty recent versions of CPAN modules that are not yet
available in Debian stable. I was looking forward to getting things running
and integrating with the community authentication service.

I got the server built, and everything was working reasonably well. Magnus and
I were just working out some issues with the proxy server configuration, and I
was starting to think about how to migrate the data over. But first, I decided
to refactor the Perl module script to use a [more efficient implementation]. I
fired it off and piped its output to the `cpan` utility to just get everything
updated. Unfortunately, unlike my first implementation, which reported only on
CPAN-installed modules, this version of the script also reported when
Debian-installed modules were out-of-date. And since I have my CPAN build
configuration set up to remove previous installations, I upgraded all those
modules, replacing them with new versions.

Well, this was a major fuckup on my part. Turns out there's no simple way to
restore Debian-distributed versions of the modules without rebuilding the
entire system. Worse, this was exactly the sort of thing the community
sysadmins feared. They have to maintain a *lot* of servers. So they naturally
prefer that they all be as similar as possible. The new PGXN server had been
*mostly* similar to what they had before, and Dave and company had been
willing to compromise quite a bit to get PGXN going, but I, unfortunately,
demonstrated how easy it is to ruin the whole thing.

So we decided that a community server isn't the right place for PGXN. At least
not yet. Perhaps in a year or two the Debian distribution will be updated to
have all the prerequisites I need. Better yet, maybe someone create a PGXN
debian distribution! (Volunteers welcomed.) Then I won't have to do anything
special and we can try again (without any `sudo` privileges for me!). But in
the meantime, I still needed to move things.

Fortunately, [depesz] came to the rescue. He has a very nice box hosting his
blog, [explain.depesz.com], and a few other things, and would I like to set
things up there? Depesz used [perlbrew] to set up a Perl install just for the
PGXN system accounts, meaning I could install any Perl modules I needed
without interfering with the system Perl. And each account has its own
privileges to run the services it needs ([Manager], [API], [Site][PGXN])
without the risk of breaking anything else. A few days after getting access,
we had everything set up and ready to go. I pulled the trigger on Monday, and
it went of without a hitch.

My thanks to depesz for the server and all the assistance, not to mention his
[donation]! PGXN now has a very nice home where it can mature.

And as for the future, I have some thoughts about that, too.

- I'd like to blog about the migration itself, and how easy it is (and isn't)
  to build PGXN.
- There are [some][] [bugs] to be fixed and [minor improvements][some] to be
  had. Interested in helping out?
- I'd love to hear your ideas about how to improve PGXN. What would make it
  better? What doesn't work quite right for you now?

And yes, now that this migration is finally done, I expect I'll have more time
to blog and work on PGXN going forward. Please leave your thoughts and ideas in
the comments. This thing is wide open to any kind of idea, and I would greatly
appreciate your feedback.



  [PGXN]: https://pgxn.org/
  [Kineticode]: https://kineticode.com/
  [Command Prompt]: https://commandprompt.com/
  [new job]: https://justatheory.com/autobiographical/iovationeering.html
  [PostgreSQL]: https://www.postgresql.org/
  [Dave]: https://pgsnake.blogspot.com/
  [Magnus]: https://blog.hagander.net/
  [Stefan]: https://www.kaltenbrunner.cc/blog/
  [CPAN]: https://cpan.org
  [more efficient implementation]: https://metacpan.org/module/ExtUtils::Installed
  [depesz]: https://depesz.com/
  [explain.depesz.com]: https://explain.depesz.com/
  [perlbrew]: https://www.perlbrew.pl/
  [Manager]: https://manager.pgxn.org/
  [API]: https://api.pgxn.org/
  [donation]: https://www.pgxn.org/donors/
  [some]: https://github.com/pgxn/pgxn-manager/issues
  [bugs]: https://github.com/pgxn/pgxn-api/issues
