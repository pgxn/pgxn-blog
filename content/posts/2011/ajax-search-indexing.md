---
title: "Question: Ajax and Search Engine Indexing"
slug: ajax-search-indexing
date: 2011-03-04T20:20:08Z
aliases: [/post/3641963548/ajax-search-indexing, /post/3641963548]
tags: [Ajax, API, Search Engine, Indexing, Indexability, Response Code, Not Found]
---

I've started working on the main (search) site in earnest now. The basic
layout is done, and I'm working on the distribution view ([mockup]). My
thinking so far has been that I would simply serve a page that requested, say,
`/dist/pgTAP/`, and that page would use Ajax requests to fetch the data from
the [API server] and display stuff. I think this will work pretty well except
for one thing: 404s.

That is, if you request `/dist/nonexistent/`, then it will load a page with
the HTTP status code `200 OK`, but then, when the Ajax request 404s, it will
show a "Not found" error message. That's all well and good, but I'm wondering
about the impact of two things:

1.  Since the page itself won't 404, search engines might index links to
    nonexistent extensions. Of course, bad links won't be *that* common, but
    of course they do happen and then tend to live forever.

2.  If the search site uses Ajax to fetch the contents of a page via JSON (or,
    for documentation, as an HTML document it will put into a div), will the
    full content be properly indexed by search engines?

So these are serious questions, in my mind. Do we loose good search engine
indelibility when we load content dynamically?

Of course, I can instead write it so that the back end fetches stuff from the
API server (and perhaps directly from the file system) and get 'round these
issues, but then it's less of a cool example of the use of the API server.

What do you think? Good advice much appreciated!



  [mockup]: https://theory.github.com/pgxn/pgtapdist.html
  [API server]: https://api.pgxn.org/
