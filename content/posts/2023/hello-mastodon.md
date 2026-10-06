---
title: Hello Mastodon 🐘
slug: hello-mastodon
date: 2023-02-18T22:53:46Z
aliases: [/post/709635160523620352/hello-mastodon, /post/709635160523620352]
tags: [PGXN, Mastodon, Twitter, PostgreSQL, LISTEN/NOTIFY]
---

Hey all you Postgres people out there! Just wanted to make a quick post
regarding the PGXN [twitter bot]. In light of the recent announcements by
Twitter to charge for API use, I updated [PGXN Manager] to also post to
Mastodon! If you're have joined the exodus to the Fediverse, give it a follow:

[@pgxn@mastodon.social]

In fact, I made some time to rewrite the notification bits of the manager
service. Previously there was just some code jammed into the upload controller
that would make an API call to Twitter. I ripped that out, and added a trigger
to the distributions table that posts a a [LISTEN/NOTIFY] message.

Then I wrote a little service that checks for new notifications every 5
seconds and dispatches to one mor more configured consumers. Today there are
just two: Twitter and Mastodon; in the future the might be more, especially
since I also added triggers to the users and mirrors tables, so we could have
the bot announce new mirrors and users. Should be pretty easy to do, might put
in the time in the next few weeks.

Oh, and for fun, I also added some emoji to the Mastodon posts, as well as the
abstract from new releases. Makes the messages more informative than on
Twitter. Of course, we could probably do the same, there, especially since the
post length was extended to 280 characters sometime after the original bot. I
have in mind to make a little library that allows the customization of
messages via configuration.



  [twitter bot]: https://twitter.com/pgxn
  [PGXN Manager]: https://manager.pgxn.org
  [@pgxn@mastodon.social]: https://mastodon.social/@pgxn
  [LISTEN/NOTIFY]: https://www.postgresql.org/docs/current/sql-notify.html
