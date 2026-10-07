---
title: Thoughts on Content Localization
slug: content-localization
date: 2011-04-01T05:33:16Z
aliases: [/post/4252690755/content-localization, /post/4252690755]
tags: [Localization, Internationalization, Translation, Locale::Maketext]
---

The new site is coming along *very* nicely. I'm really excited about it. The
addition of the API server has turned out to be an inspiration. I never
realized how clean the separation between content and presentation could be
until I did this -- even though I've preached that very separation for years
in one of my [day jobs].

But speaking of that separation, I am running into a bit of an annoyance with
content specific to the site. All the content from the network is working
beautifully -- I just fetch it from the API server. But the site has its own
content page independent of the network, including an FAQ, a page on setting
up a mirror, a list of our backers, etc. And the ugly bit has to do with
localization.

I've built the site (and [PGXN Manager]) using the Perl standard localization
library [Locale::Maketext]. And it works great for short labels and such, like
so:

```perl
our %Lexicon = (
    'hometitle' => 'PGXN: PostgreSQL Extension Network',
    'PostgreSQL Extension Network' => 'PostgreSQL Extension Network',
    'About PGXN' => 'About PGXN',
    'User' => 'User',
    'Recent Uploads' => 'Recent Uploads',
    'Blog' => 'Blog',
    'Frequently Asked Questions' => 'Frequently Asked Questions',
    'Release It' => 'Release It',
);
```

Translators just have to copy this code into a subclass and change the strings
on the right-hand side of the `=>`s. Easy, right? It's also very good with
variables and pluralization. I just use it in the templates like this:

```perl
title { T 'PostgreSQL Extension Network' };
```

What sucks, however, is when I need to have longer content pages, especially
those that mix HTML in with the text. The FAQ is particularly relevant here:
most of the answers have at least one link embedded in them. The trouble is,
templating language encodes HTML entities. So I can't easily put HTML in the
localization files, unless I put it in raw, and the output it raw. So the
localization file might have a line like:

```perl
'Send us an <a href="mailto:pgxn@example.com">email</a> with your questions' =>
'Send us an <a href="mailto:pgxn@example.com">email</a> with your questions',
```

And then the template would be told to output the text raw, without encoding
HTML entities:

```perl
p { outs_raw T 'Send us an <a href="mailto:pgxn@example.com">email</a> with your questions' };
```

So now we've lost the benefit of the templating, really: we have to write the
raw HTML ourselves. And besides, for multi-paragraph documents, it gets
annoying to have a bunch of localization lines with names like
`about_paragraph_1`, `about_paragraph_2`, etc. That makes it additionally
difficult for translators. There's just too much crap getting in the way. I've
always liked the idea of having all translations in one place, but I don't
think there's enough of an advantage to that to overcome the annoyances.

So I'd like to find a better way to manage these content pages more like
*documents,* in a format that's easy for translators to manage and isn't so
broken up. I'm open to suggestions. Ideally the content would still be all
kept in one place for a given language, but maybe that's just not feasible.

The one thing that occurs to me is to keep documents for given languages in
separate [Markdown]-formatted files. Each language would have a directory with
the language code ("en", "en-uk", "fr", etc.), and then at build time they
would be converted to HTML for fast serving at run-time. Or maybe at
distribution-packaging time. I've resisted something like this for a while
because it means one has to find stuff to translate in two different places
(the localization library for UI elements and the directory of documents for
content), but maybe it would be best.

How have you solved this problem in your apps? How badly have you annoyed your
translators?

Oh, while I'm thinking about it, where should documentation for the API server
be published? I could include it in the [PGXN::API] source, which would be
useful for anyone else who wanted to run an API server (which will be dead
easy, by the way). But there will also be some stuff that's specific to a
given installation (proxying, downtime, etc.). So maybe it should go on the
main site or something? Or perhaps we need a dedicated wiki server for shit
like this (and some of those content pages like the FAQ, maybe).

Thoughts?

  [day jobs]: https://www.kineticode.com/
  [PGXN Manager]: https://manager.pgxn.org/
  [Locale::Maketext]: https://search.cpan.org/perldoc?Locale::Maketext
  [Markdown]: https://daringfireball.net/projects/markdown/
  [PGXN::API]: https://github.com/theory/pgxn-api/
