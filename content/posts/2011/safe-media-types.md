---
title: What Media Types for Browsing?
slug: safe-media-types
date: 2011-04-21T16:56:43Z
aliases: [/post/4809393274/safe-media-types, /post/4809393274]
tags: [MIME, Media Type, MIME Type, Browse, Source, User Submitted, Content, Text, Plain Text, HTML, C, JSON]
---

Thanks to a comment from "Anon," I'm auditing the media types used to serve
files from the "Browse" links on the site. The browse interface is managed by
[Plack::App::Directory], which gets its media type mapping from [this file].
What PGXN::API does is simply change the mappings of some of those types to
`text/plain`. As of [this commit], files are served as plain text if:

- The strings "html", "x-c", "xml", "calendar", or "vcard" appear in the
  default media type; or
- The file name extension is one of `.bat`, `.css`, `.eml`, `.js`, `.json`,
  `.mime`, or `.swf`.

Are there other [media types][this file] that should be disabled for safe
browsing of user-submitted content?

  [Plack::App::Directory]: https://search.cpan.org/perldoc?Plack::App::Directory
  [this file]: https://github.com/miyagawa/Plack/blob/master/lib/Plack/MIME.pm
  [this commit]: https://github.com/pgxn/pgxn-api/commit/0157c18cbc0835b627fa2e42b2433337f5f3fff5
