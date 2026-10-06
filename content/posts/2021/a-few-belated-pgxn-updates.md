---
title: A Few Belated PGXN Updates
slug: a-few-belated-pgxn-updates
date: 2021-05-15T03:16:44Z
aliases: [/post/651216661677064192/a-few-belated-pgxn-updates, /post/651216661677064192]
tags: [PGXN, Upgrade, TLS, Retina]
---


The last couple weeks I've returned to PGXN and made a few updates. Nothing
huge, but all long overdue.

First up, [PGXN Manager] is all TSL now. No public HTTP-only site. What
started as a convenient division of labor (http for the public site and https
for the authenticated site) turned out to be quite irksome --- especially
since some browsers wouldn't load the non-TLS site at all anymore. So I did
away with it, and now it's all TLS, authenticated and not. (The API and search
sites have been TLS for a while now.)

I also fixed a few long-standing bugs in PGXN Manager, most of which weren't
visible to end-users, but have annoyed me over the years. A few silly server
errors, some instances of uploads failing for anything other than zip files,
that sort of thing.

Oh, and if you're a extension author, PGXN Manger now allows updates to old
distribution versions to be uploaded. Previously it only allowed a new version
to be greater than all previous versions. Now it will allow a new X.Y.Z
version if X.Y previously existed and the new .Z is greater, and a new X.Y
version if X previously existed and the new .Y is greater than any previous
X.Y. To get this to work properly, I also dropped the check for versions of
extensions in the uploaded files. It would just be too complicated to add a
bunch of rules more likely to annoy than not. So it now only enforces patterns
for distribution release versions. I trust extension authors not to then lower
extension versions on new releases --- that would just be silly, and not
helpful to your users.

As part of this work, I also revamped the management of the PGXN server. Back
in 2010 I wrote Capistrano files to manage the server, but have long ceased to
use them, as they ceased to work. I've now removed all that detritus from the
PGXN Manager, API, and Site repositories, and replaced them all with a single
new repository, [pgxn-ops], which contains Ansible playbooks to manage all the
services. They don't deal with a lot of the server-side stuff, which depesz
handles separately, but they now make it much easier to build and deploy a new
release, restart it remotely, manage passwords, etc.

And finally, I've updated the [PGXN search site]. In addition to fixing a few
long-standing minor annoyances (borders around images linked in documentation,
broken links, etc.), I also updated most of the graphics to be
retina-friendly. So it should look a lot sharper on your hi-res screens now.
Check it out!

Next up I think I'd like to make the search site more mobile-friendly, and
then perhaps I'll finally go back and attack the terrible search provided by
the API. I'll try to do it in less than five years this time.



  [PGXN Manager]: https://manager.pgxn.org
  [pgxn-ops]: https://github.com/pgxn/pgxn-ops
  [PGXN search site]: https://pgxn.org/
