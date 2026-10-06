---
title: PGWest, Search Site Sneak Peak
slug: pgwest-and-sneak-peak
date: 2010-11-03T21:09:00Z
aliases: [/post/1473157261/pgwest-and-sneak-peak, /post/1473157261]
tags: [Search, PgWest, Presentation, Extensions, Postgresql 9.1]
---

I've been working on my [PostgreSQL Conference West presentation], which
heavily features PGXN, of course. I think it's looking good. If you're at
[PGWest] or are in the San Francisco area and free, come see the talk! Should
be a good introduction to creating PostgreSQL extensions and distributing them
on PGXN. The latest bit I added is a section on the modifications needed to
support the forthcoming [`CREATE EXTENSION`] support slated for 9.1.
Fortunately it's dead simple, and will make dealing with extensions in the
database a lot simpler, administratively. Really looking forward to that. Of
course I'll post slides once the talk is over.

As part of preparing for the talk, and because there isn't currently much to
actually *see* of PGXN, I've been mocking up the layout for the new search
site, which as you know from the [status page] is the next part of the project
I'm slated to work on. I've been committing the mockps to the gh-pages branch
of the repository, which means you can see what it looks like live on the net
[right here]. That's the home page, including our sponsor links and tag cloud.
Click the "PGXN Search" button to see a mockup of search results (or get them
[here]). Click on any search result to see the mockup of a documentation page
(or link it [here][1]). The design is based on the [lazydays] open-source Web
design, and I'm quite happy with it. Your thoughts?

As this comes together, I'm gearing up to start hacking on the app to produce
the search site. At this point, I'm thinking that it would become the new home
page for [PGXN], rather than a separate search.pgxn.org site. Thoughts?

I'll post the slides tomorrow.



  [PostgreSQL Conference West presentation]: https://www.postgresqlconference.org/content/building-and-distributing-postgresql-extensions-without-learning-c
  [PGWest]: https://www.postgresqlconference.org/2010/west/
  [`CREATE EXTENSION`]: https://s.coop/pgext
  [status page]: https://pgxn.org/status.html
  [right here]: https://theory.github.com/pgxn/
  [here]: https://theory.github.com/pgxn/results.html
  [1]: https://theory.github.com/pgxn/pgtap.html
  [lazydays]: https://www.oswd.org/design/preview/id/2839
  [PGXN]: https://pgxn.org/
