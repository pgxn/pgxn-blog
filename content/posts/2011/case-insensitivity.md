---
title: Case Insensitivity
slug: case-insensitivity
date: 2011-04-26T04:09:08Z
aliases: [/post/4948135198/case-insensitivity, /post/4948135198]
tags: [Lowercase, Case, Case Sensitive, Case Insensitive, Meaning]
---

I pushed out new versions of PGXN::Manager, PGXN::API, and PGXN::Site today,
all with the aim of resolving some issues with case-sensitivity. I was alerted
to this when Daniele Varrzzo released [italian_fts] with the [Italian tag] and
the API sync blew up. Oops.

So here's the deal. The following objects have case-insensitive names on PGXN:

- Distributions
- Extensions
- Versions
- Users (nickname)
- Tags

These names are case-preserving, so they should show up with the proper
capitalization in JSON files and whatnot. But to avoid issues with URLs, I
decided that all file names should be lowercase only. Fortunately, nearly
everything was lowercase already: only a few nickname JSON files were not
lowercased, plus the files for [pgTAP]. So once I got the code updated, I
renamed those files appropriately. In the case of the pgTAP download, I
actually re-indexed it, so that the prefix in the zip file would properly be
"pgtap-0.25.0" instead of "pgTAP-0.25.0", since the file name is now
`pgtap-0.25.0.pgz`.

So why make these objects case-insensitive? Well, because I tend to think of
the names of things as having *meaning*, and really, there is no difference in
meaning between "pgtap," "pgTAP," or "PGTap." I've relased pgTAP; I think it
would be as confusing as hell if someone else released a completely different
distribution named "PGTap." (And I went through a bit of hell when the
[Apache::test] CPAN module was replaced with [Apache::Test][1], a completely
new module. It really messed things up on the case-insensitive file system I
use).

But as I said, the case is preserved in the name I use, just not in the names
of files. So if you look at the [pgTAP distribution page][pgTAP], you'll see
that "This Release" is labeled "pgTAP 0.25.0" (I need to fix the `<h1>`
element still).

  [italian_fts]: https://pgxn.org/dist/italian_fts/
  [Italian tag]: https://pgxn.org/tag/italian/
  [pgTAP]: https://pgxn.org/dist/pgtap/
  [Apache::test]: https://search.cpan.org/perldoc?Apache::test
  [1]: https://search.cpan.org/perldoc?Apache::Test
