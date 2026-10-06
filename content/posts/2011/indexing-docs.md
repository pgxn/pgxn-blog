---
title: Thoughts on Indexing and Documentation
slug: indexing-docs
date: 2011-03-22T05:13:00Z
aliases: [/post/4018670551/indexing-docs, /post/4018670551]
tags: [Indexing, Full Text Search, README, Documentation, Search Results, Search]
---

So I'm designing the full text indexing for the PGXN search site. I'm modeling
it on [CPAN Search], which has been great. There are four search options:

- Full documentation search. This is the most common. Includes doc title and
  body.
- User search. Search on names, nicknames, email addresses, URIs, etc.
- Distribution search. Search on distribution name, abstract, description,
  tags, and the README.
- Extension search. Search on extension name and abstract.
- Tag search. Search on tag name only.

The documentation search is the one I'm perhaps least sure about. It assumes
that each extension in a distribution will have documentation. But so far that
has not really been the practice for PostgreSQL extensions. Most folks seem to
stick the documentation in the README. And even then it can be [almost
nothing]. So a search for "count nulls" probably would not find "countnulls"
extension, because there is no documentation. What should I do about this? I'm
thinking one of:

- Encourage folks to write documentation. I'm going to do this anyway, because
  the docs will really help the visibility of an extension on the site. It
  looks [like this]. If you have no docs for an extension, your extension will
  not appear in the search results (or perhaps it might, but link to the
  distribution).

- If there is no documentation for an extension in a distribution, index the
  README as the documentation. I'm not really keen on this idea, because the
  README should describe the distribution, how to install it, etc. I'm
  planning to use it in the distribution-specific index. Documentation of the
  extension should be more about how the extension works, what it's interface
  is, etc. Or so it seems to me, at least (I'm admittedly biased to this
  practice among CPAN modules). But at least with this approach there would be
  a link to "documentation" for an extension on the search site.

Erm, not really thinking of any other options. I feel pretty strongly that
folks should write docs for their extensions, as much as possible, and I've
set things up so that, from PGXN's point of view, at least, you can write
documentation in whatever format you like (assuming the format is supported by
or added to [Text::Markup]), as long as they're in a `doc/` or `docs`
directory. I want it to be as easy as possible. But I also want there to be
decent search results ASAP.

Comments?



  [CPAN Search]: https://http//search.cpan.org
  [almost nothing]: https://master.pgxn.org/dist/countnulls/1.0.0/README.txt
  [like this]: https://theory.github.com/pgxn/pgtap.html
  [Text::Markup]: https://search.cpan.org/perldoc?Text::Markup
