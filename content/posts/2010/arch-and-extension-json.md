---
title: Architecture, New Extension JSON Format
slug: arch-and-extension-json
date: 2010-09-17T17:33:00Z
aliases: [/post/1138292188/arch-and-extension-json, /post/1138292188]
tags: [PGXN Manager, PGXN::Manager, Plack, Catalyst, Dancer, Jifty, Extension, Metadata]
---

I'm making good progress on [PGXN Manager]. Hopefully I can start alpha
testing it next week. As I mentioned [previously], I had estimated 40 hours of
work to create it, but was hoping to get it done in around 30 (because I spent
10 extra hours on the database design). So far I'm at 23 hours, so it's
looking pretty good.

Architecturally, I've gone for a very minimal [Plack]-based app. No
[Catalyst], [Jifty], or even [Dancer]. Just a very simple [Plack app] that
uses [Router::Simple::Sinatraish] to route URIs to the appropriate
[controller] actions (which are just class methods). The controller just
dispatches to [Template::Declare]-based [templates] for the HTML rendering. I
guess I've kind of created my own framework here, but really, there ain't much
to it. This app is simple enough that I couldn't see the use of adding all the
overhead of a framework.

Meanwhile, I've been hacking on the [distributon class]. This will be the core
class of the app. It takes a [Plack upload] object and a username and does all
the rest, analyzing an uploaded archive, normalizing it if necessary,
registering it with the database, and indexing it by updating all the
appropriate JSON files on the mirrors. It's nearly finished, but I have one
other thing to do in the database, first.

In my [last update][previously], I asked for advice on whether or not PGXN
should allow an extension with the same version number to appear in multiple
distributions. And thanks to a [comment from Aristotle], I'm changing it to
allow that. But it also means that my [original extension json spec] needs to
change.

Here's an example of what I'm thinking. Say that there are three versions of
an extension named "trip", and that they appear in distributions as follows:

    trip 0.2.6
      pair-0.3.0

    trip 0.2.5
      trip-0.2.2
      pair-0.2.2rc
      pair-0.2.1
      trip-0.1.1

    trip 0.2.4
      pair-0.1.1rc
      pair-0.1.0

So sometimes it's in the "trip" distribution and other times it's in the
"pair" distribution. My thought is that, for a given version, it would list
the distributions it's in in reverse chronological order (by upload date). So
the format would be:

    {
        "latest": "stable",
        "stable": { "dist": "pair", "version": "0.3.0" },
        "testing": { "dist": "pair", "version": "0.2.2rc" },
        "distributions": {
           "0.2.6": [
              { "dist": "pair", "version": "0.3.0" }
           ],
           "0.2.5": [
              { "dist": "trip", "version": "0.2.2" },
              { "dist": "pair", "version": "0.2.2rc", "status": "testing" },
              { "dist": "pair", "version": "0.2.1" },
              { "dist": "trip", "version": "0.1.1" }
           ],
           "0.2.4": [
              { "dist": "pair", "version": "0.1.1rc", "status": "testing" },
              { "dist": "pair", "version": "0.1.0" }
           ]
        }
    }

This way, every distribution it's included in is listed, and clients can
quickly tell where to find the latest stable, testing, and unstable versions,
and which of those is the most recent. This is a bit more convoluted than [the
original][original extension json spec], but I think is a good choice, in that
it's comprehensive but also easy to figure out what's the latest.

Unless you can think of a better format, this is what I'm going with.
Comments?

Look for a post next week announcing an alpha program!



  [PGXN Manager]: https://github.com/theory/pgxn-manager
  [previously]: https://blog.pgxn.org/post/1082188310/db-status-update
  [Plack]: https://plackperl.org/
  [Catalyst]: https://www.catalystframework.org/
  [Jifty]: https://jifty.org/
  [Dancer]: https://perldancer.org/
  [Plack app]: https://github.com/theory/pgxn-manager/blob/master/lib/PGXN/Manager/Router.pm
  [Router::Simple::Sinatraish]: https://search.cpan.org/perldoc?Router::Simple::Sinatraish
  [controller]: https://github.com/theory/pgxn-manager/blob/master/lib/PGXN/Manager/Controller.pm
  [Template::Declare]: https://search.cpan.org/perldoc?Template::Declare
  [templates]: https://github.com/theory/pgxn-manager/blob/master/lib/PGXN/Manager/Templates.pm
  [distributon class]: https://github.com/theory/pgxn-manager/blob/master/lib/PGXN/Manager/Distribution.pm
  [Plack upload]: https://search.cpan.org/perldoc?Plack::Request::Upload
  [comment from Aristotle]: https://blog.pgxn.org/post/1082188310/db-status-update#comment-75931198
  [original extension json spec]: https://blog.pgxn.org/post/988613682/restful-directory#extension-json
