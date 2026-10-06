---
title: New release for the PGXN client
slug: new-release-for-the-pgxn-client
date: 2011-05-02T00:59:37Z
aliases: [/post/5118152273/new-release-for-the-pgxn-client, /post/5118152273]
tags: [PGXN, Client, Release, Uninstall, Drop, sudo]
---

By Daniele Varrazzo

During the last days I've done some lightweight hacking on the PGXN client,
and I've just released the last package on PyPI, with the still very shy
version number of 0.1a4.

First change: the package name. The package is now called `pgxnclient`,
without the point. As [Marti has commented], the point in the name could
create problems to some distributions which may have needed to munge it into a
shape fitting their naming constraints. So, in order to install the last
release you will need the command:

    $ easy_install pgxnclient

The name of the entry point script is still `pgxn` and it shouldn't change.

One of the new features in this release is the addition of commands `drop`, to
remove the the extensions of a distribution from a database (opposite of the
`load` command) and `uninstall`, to remove installed files from the system
(opposite of `install`).

Another change is the ability to work on a local zip or directory instead of
using the API to get metadata and package from PGXN. The main goal of this
feature is to help the packagers to test their packages with the client before
submission, verifying that the data uploaded can be used in an automatized
way.

A third change is in the `install` command: it previously required to be run
as root if, as likely, the PostgreSQL directories are system ones. I was
feeling this need slightly scary as download and build phases would have been
run as root too, so now `sudo` is invoked by the script only when installing,
with command line options allowing its customization.

What I feel mostly missing now is documentation: I know that adding it will
force a big review of the code, adding missing comments, add tests to enforce
what documented... so it's something I will start tomorrow. Once documentation
and code have converged there will also be a more sound basis to discuss about
the client features. And dare version 0.2, maybe :)

Thanks everybody for the feedback. If you want to have a chat about the client
or PGXN in general, the best place is [the PGXN group]. See you there!



  [Marti has commented]: https://blog.pgxn.org/post/5026314153/writing-a-client-for-pgxn#dsq-comments
  [the PGXN group]: https://groups.google.com/group/pgxn-users
