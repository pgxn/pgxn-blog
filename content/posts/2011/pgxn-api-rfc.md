---
title: PGXN API RFC
slug: pgxn-api-rfc
date: 2011-02-04T04:05:33Z
aliases: [/post/3099288750/pgxn-api-rfc, /post/3099288750]
tags: [Search, API, Mirror, JSON, CPAN, metacpan, JavaScript, Application]
---

Things slowed up a bit over the last couple of months, I admit. There are any
number of reasons for that, not the least were the intrusion of the holidays
and a [little side project] I've been hacking on after-hours (and sometimes
during-hours). But I'm ramping things up again now and need *your* feedback on
my current plans. Here's what I'm working on: the search site.

## Search Sites and APIs

Well, sort of. First of all, I've decided that the "search site" should not be
a separate thing. The [main site] will be the search site. This is following
the example of [JSAN], as well as feedback from [Graham Barr], who created and
maintains [CPAN Search]. Apparently people are often confused that
[search.cpan.org][CPAN Search] is separate from [www.cpan.org]. No point in
adding in confusion from the beginning. And besides, now that the PGXN
fund-raising [is over], I don't know what else would go on the home page.

The other thing that's happened is, just as I was getting my butt in gear on
this stuff, a new CPAN search site came to my attention, [μετα CPAN]. This is
an interesting project. What they did instead of creating a monolithic HTML
search site is to create a [simple API] that serves nothing but JSON. It has
search and displays metadata for CPAN objects (distributions, maintainers,
modules, etc.). The search site, then, is not really a site at all, but a pure
JavaScript application. Once you load it, it just uses the API server to get
all the data. There are a few tricks server-side to proxy the API server so as
to avoid cross-site scripting issues. But otherwise it just works in the
browser.

Now I'm not sure I'll do the same thing, exactly, but there's a lot of appeal
in creating a RESTful API server that's independent of the search site, and
then building the search site to use it. It also has the advantage of being
useful for other projects to just use. Want to create a PGXN search widget for
your blog? Yeah, there's an API for that.

## A Super RESTful Directory

Of course, thanks to the "RESTful Directory" design for the mirrors (described
[here] and revised [here][1]), any mirror is a lightweight API already.
There's a *lot* of metadata one can get just from the static JSON files it
generates. The design is flexible--but designed with a command-line client in
mind. As such, many commands executed in a command-line client would likely
requires multiple requests to a mirror. For example:

    > install pgtap

This would request `/by/extension/pgtap.json` from the server. It would then
parse that file and see that the latest stable version of pgTAP is in the
distribution "pgTAP" at version "0.25.0". So it would then download
`/dist/pgTAP-0.25.0.pgz` to install.

This is great for a command-line client, but wouldn't be so great for a search
site to be responsive. Ideally, a site should send a single request to get all
the data it needs for a particular page.

So here's what I'm thinking for a PGXN API server: It will offer a superset of
the functionality of any other PGXN mirror. That is, all the JSON files in a
mirror will be present, but many of them will have more information than they
would on other mirrors. And then, of course, there will be other URIs to offer
additional API calls.

### Details

So what does that look like? Let's take the pgTAP distribution, which I
released on PGXN earlier this week. To find the pgTAP distribution, one
requests:

> [`/by/dist/pgTAP.json`]

From that, one can see that the latest table release is 0.25.0, and so one can
then request

> [`/dist/pgTAP/pgTAP-0.25.0.json`]

to get all the metadata for that particular release. What I propose, to avoid
the two requests, is to include the contents of the second file in the first.
That would then have all the data necessary to generate the [pgTAP
distribution page] on the PGXN site.

The API would offer similar supersets of data for the [extension] , [owner] ,
and [tag] metadata files, to have the data necessary for the design of the
corresponding [extension][2], [owner][3] and tag layouts of the site.

### Additional Resources

In addition to adding metadata to the existing mirrored JSON files, there
would be other resources available for request from the API server. They would
include:

- Extension Documentation. Each distribution may include documentation for
  included extensions in the `doc` subdirectory. These will go under the
  directory for a specific distribution such as
  `/dist/pgTAP/pgTAP-0.35.0/doc/pgtap.html`. The latest version of each
  document would also be available under `/by/extension`, as in
  `/by/extension/pgtap.html`. This requires that the documentation file have
  the same base name as the extension file itself.

- Other documentation. I'd like to support arbitrary documentation, such as
  for included binary executables, HOWTOs, etc. The canonical copies will go
  under the versioned distribution URL, of course, but I'm not sure about
  permalinks. That might require an extension of the [Meta Spec]; I haven't
  quite figured that out, yet.

- Source code. There will be an interface to browse an unpacked copy of any
  distribution as plain text. This will be under `/src`, as in
  `/src/pgTAP/pgTAP-0.35.0/`.

### Search API

Of course. This is the big one, really. I think it makes sense to have the
`/by` URI respond to search requests. Thus, a request for

    /by?q=testing

would search everything. If you only want to search a certain category of
object, you'd hit the appropriate URI:

    /by/dist?q=tap
    /by/owner?q=clochard
    /by/tag?q=test
    /by/extension?q=gis

The nice thing about this is that it retains the existing entity URLs. The
directory level determines which entities you get.

## Your Thoughts?

So that's my thinking on the search API. I'm going to start hacking on it in
earnest tomorrow, and perhaps next week I can get a very early version out
(basically just another mirror to start with).

But what do you think? Seem like a sane approach? Am I missing anything
obvious or doing anything clearly stupid? Please let me know in the commments!



  [little side project]: https://www.designsceneapp.com/
  [main site]: https://pgxn.org/
  [JSAN]: https://openjsan.org
  [Graham Barr]: https://search.cpan.org/~gbarr/
  [CPAN Search]: https://search.cpan.org/
  [www.cpan.org]: https://www.cpan.org
  [is over]: https://blog.pgxn.org/post/2063677299/goooooooaaaal
  [μετα CPAN]: https://search.metacpan.org/
  [simple API]: https://github.com/CPAN-API/cpan-api/wiki/API-docs
  [here]: https://blog.pgxn.org/post/988613682/restful-directory
    "A RESTful Directory"
  [1]: https://blog.pgxn.org/post/1138292188/arch-and-extension-json
    "Architecture, New Extension JSON Format"
  [`/by/dist/pgTAP.json`]: https://master.pgxn.org/by/dist/pgTAP.json
  [`/dist/pgTAP/pgTAP-0.25.0.json`]: https://master.pgxn.org/dist/pgTAP/pgTAP-0.25.0.json
  [pgTAP distribution page]: https://theory.github.com/pgxn/pgtapdist.html
  [extension]: https://master.pgxn.org/by/extension/pgtap.json
  [owner]: https://master.pgxn.org/by/owner/theory.json
  [tag]: https://master.pgxn.org/by/tag/testing.json
  [2]: https://theory.github.com/pgxn/pgtap.html
  [3]: https://theory.github.com/pgxn/theory.html
  [Meta Spec]: https://pgxn.org/meta/spec.txt
