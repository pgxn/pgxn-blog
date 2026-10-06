---
title: "It's Finally Up: The New PGXN Site!"
slug: new-pgxn-site
date: 2011-04-14T06:32:06Z
aliases: [/post/4601750614/new-pgxn-site, /post/4601750614]
tags: [Site, Launch, Final, Awesome, API]
---

I'm pleased to announce that the new [PGXN SITE] went live last night. Some of
the things it does:

- Full text search of all documentation, distribution metadata and `README`s,
  extensions metadata, and users
- list of the 56 most recent uploads
- Find users by first letter of nickname (mostly so search spiders can drill
  down to all site content)
- A tag cloud with the 56 most commonly-used tags (they're all the same size
  and color right now because each one is used only once at the moment)
- Pages for users ([example])
- Pages for distributions, with inlined `README` ([example][1])
- Browse unzipped distributions ([example][2])
- Pages for documentation ([example][3])
- Permalinks for extensions [example][4])

The entire site is backed by the [API Server], which is mainly composed of
static JSON and HTML files read by the site back end from the local file
system. So the site is quite fast. Please browse around, let me know what you
think, and please [report any bugs] you find.

There are a few more things left to do on my list:

- Add pages for users who don't yet have distributions on the network.
- Convert the localization interface to `gettext` and solicit translations.
- Document the API; might lead to some changes as I run across
  inconsistencies.
- Add an Atom feed of recent releases.
- Add aliases for `$nickname@pgxn.org` addresses, to forward to personal
  addresses; this way we can just publish pgxn.org addresses.
- Write documentation on how to optimize distributions for optimal exposure on
  PGXN.
- Consider changing the Permalink URL to something shorter. Suggestions?

If you'd liked to help out with any of these tasks, just let me know. Better
yet, grab the source from [GitHub] and just hack!

Oh, and speaking of the source, I would *really* appreciate a code review or
four. Although I have solicited advice on key questions fro people more
knowledgeable than I (thanks Aristotle, Miyagawa, Andreas, Graham), the code
is entirely my own. More eyes will be a huge help!

Tomorrow I'll blog a bit about the architecture for the network. I'm quite
happy with it.



  [PGXN SITE]: https://pgxn.org/
  [example]: https://pgxn.org/user/alexk
  [1]: https://pgxn.org/dist/explanation/
  [2]: https://api.pgxn.org/src/semver/semver-0.2.0/
  [3]: https://pgxn.org/dist/explanation/doc/explanation.html
  [4]: https://pgxn.org/extension/pair
  [API Server]: https://api.pgxn.org/
  [report any bugs]: https://github.com/pgxn/pgxn-site/issues
  [GitHub]: https://github.org/pgxn/
