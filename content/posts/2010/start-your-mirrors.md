---
title: Start Your Mirrors!
slug: start-your-mirrors
date: 2010-08-13T23:45:00Z
aliases: [/post/949562598/start-your-mirrors, /post/949562598]
tags: [Getting Started, JFDI, Mirror, Progress, rsync, HTTPD]
---

Work has finally started. We now have [mirroring]. This was the first task for
the project, and it's now checked off on the [status] page. Hurrah!

This turned out to be a pretty simple task, of course. For now I'm using the
[Kineticode] (my other employer) server to host both the [PGXN site] and the
master mirror. This is temporary, just a way to get things up and going as
quickly as possible. The master mirror runs `rsyncd` with read-only access to
the `pgxn` path. So all you have to do to mirror it is set up a `cron` job
like:

    rsync -az --delete rsync://master.pgxn.org/pgxn /path/to/pgxn

If the destination directory is under an HTTP root, you're done. Otherwise,
throw a web server over it and *then* you're done.

Of course, the web server isn't really necessary yet. Soon it will be the
interface for downloading distributions and metadata from the mirrors. Right
now, the mirrors just have two files in them (a `README` and an `index.html`).
This is just the start of things. The next step is to figure out the directory
structure for the mirrors. I'm working on that right now and will soon be
asking for feedback. Watch this space for details!



  [mirroring]: https://pgxn.org/mirroring.html
  [status]: https://pgxn.org/status.html
  [Kineticode]: https://www.kineticode.com/
  [PGXN site]: https://www.pgxn.org/
