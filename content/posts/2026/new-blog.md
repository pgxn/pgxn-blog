---
title: Welcome to the New PGXN Blog
slug: new-blog
date: 2026-10-08T15:39:56Z
description: The PGXN Blog has a new home.
tags: [Blog, Hugo, Tumblr]
type: posts
---

After a couple years putting off, I finally made some time to migrate the
[PGXN Blog] --- the canonical site where this very post lives --- from the
original Tumblr blog set up in 2010. It now lives [in Git] as Markdown files
and relies on [Hugo] to build the site and GitHub Pages to serve it. I even
migrated the old Twitter cross-posts; their titles all end in "ⓣ" to
distinguish them from normal posts.

I've made quite a few Hugo sites over the years, including my [personal blog]
and a couple of internal blogs for present and past employers. It lets me
write posts in Markdown as I prefer, without the limitations of Tumblr's
hobbled formatting. It also allows others to contribute simply by making a
[pull request].

I started the layout with the minimalist [hugo-blog-awesome] theme (its only
JavaScript operates the dark/light mode toggle), but modified it to suit my
whims, including a tag list at the end of each post and proper formatting of
titles. The site consists entirely of static files, so should load quite
quickly. I also borrowed the [Atom feed] code from my [personal blog] to
publish full text posts for syndication to [Planet PostgreSQL] and elsewhere.

Each post has a new URL, `/{year}/{slug}`, but the old URLs remain as simple
pages with a `<meta http-equiv="refresh">` element to redirect to the new URL.
The new feed also reuses the IDs of the old Tumblr posts, so subscriptions
should never show dupes; [file an issue] if you notice any problems.

Thanks for continuing to tune in to PGXN news and updates over the years. I've
got a couple of interesting things coming up shortly, so watch out for more!

  [PGXN Blog]: https://blog.pgxn.org "News and Updates from the PostgreSQL Extension Network"
  [in Git]: https://github.com/pgxn/pgxn-blog "The PGXN Blog on GitHub"
  [Hugo]: https://gohugo.io "Hugo: The world’s fastest framework for building websites"
  [personal blog]: https://justatheory.com/ "Just a Theory"
  [pull request]: https://github.com/pgxn/pgxn-blog/pulls "PGXN Blog Pull Requests"
  [hugo-blog-awesome]: https://github.com/hugo-sid/hugo-blog-awesome
    "Fast, minimal blog with dark mode support"
  [Atom feed]: https://en.wikipedia.org/wiki/Atom_(web_standard) "Wikipedia: Atom (web standard)"
  [Planet PostgreSQL]: https://planet.postgresql.org 
  [file an issue]: https://github.com/pgxn/pgxn-blog/issues/new "Create new PGXN Blog Issue"