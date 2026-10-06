---
title: "About the Infrastructure: PGXN Manager"
slug: pgxn-manager
date: 2011-04-23T02:54:34Z
aliases: [/post/4854707157/pgxn-manager, /post/4854707157]
tags: [PGXN::Manager, Infrastructure, Code, Perl, SemVer, Json, Normalization, Validation, Zip]
---

The PGXN infrastructure is currently made up of four parts:

- The network of [mirrors], derived from the [master mirror], and synchronized
  on various schedules via `rsync` ([host a mirror]).

- [PGXN Manager] ([code]) is the core of the network. It provides the
  interface for users to upload releases, processes those releases, indexes
  them, and puts them on the master mirror. Details below.

- [PGXN API] ([code][1]) is a mirror server with benefits. Once an hour, it
  `rsync`s from the master mirror, and does extra processing of new and
  modified files, notably full-text indexing. I'll write up some details next
  week.

- [PGXN Site] ([code][2]) powers the main site. It's a thin wrapper around the
  API server, using [WWW::PGXN] to fetch JSON files and convert them to HTML.
  I'll write more about this bit next week, too.

Some details on Manager. This is by far the most complicated part of the
system. Which is funny, because I hadn't anticipated that when I started work
on PGXN (I'd estimated half as many hours as for implementing the site). But
as I worked through design issues and wrote the code, the need for that
complexity became apparent -- and not just because it's the only part that
offers authentication. *The main reason it's complex is so that no other part
needs to be.*

Allow me to explain. People can upload almost anything as a distribution. So
long as the `META.json` adheres to [the spec], the rest can be just about
anything. But the upload doesn't necessarily end up on the network unmodified.
Sure, if you follow the guidelines of the [HOWTO] what you uploaded will be
exactly what ends up on the network. But I didn't want to be that strict about
PGXN Manager would accept. So in addition to verifying the structure of your
`META.json` file, [PGXN::Manager::Distribution] also:

- Extracts the archive. If it's not a zip file, a zip archive is created from
  the contents of the uploaded file. You can upload any kind of archive
  readable by [Archive::Extract]. The currently supported formats are: `.tar`,
  `.tar.gz`, `.gz`, `.Z`, `tar.bz2`, `.tbz`, `.bz2`, `.zip`, `.xz,`, `.txz`,
  `.tar.xz` and `.lzma`. By always converting to a zip file, PGXN client apps
  can be quite simple, not having to worry about the archive format.

- Validates the `META.json`. This part isn't perfect, yet. I fixed a bug today
  where it would die and return a 500 on a missing version number. It'd
  probably be worthwhile to adapt [CPAN::Meta::Validator] to validate PGXN
  `META.json` files at some point, both for Manager and for developers wanting
  to validate before uploading.

- Normalizes all version numbers in the `META.json` into [semantic versions].
  You can specify the distribution version, prerequisite versions, and
  extension versions as simple numbers and they'll be converted to semantic
  versions. A version like "1.20", for example, becomes "1.2.0". See the
  [SemVer documentation] for details on how versions are normalized via the
  `declare()` method. This normalization is done so that client applications
  will get known valid semantic versions to compare when determining
  dependencies. However, it's best that they be semantic versions to begin
  with. Normalized versions will be written back to the archive `META.json`
  file (with the "generated_by" key updated to reflect that PGXN Manager
  regenerated the file). If no versions need validating, the archive
  `META.json` will be left alone.

- Makes sure the zip archive has a directory prefix named `"$dist-$version/"`.
  If the archive has no directory prefix, or if the prefix is not
  `"$dist-$version/"`, the archive is rewritten with that prefix. This ensures
  that the archive will always extract into a directory with the same name as
  the archive and not spray files all over your desktop when you unzip it.

- Copies or writes out a new zip file named `"$dist-$version.pgz"`. Think of
  `.pgz` as "PostgreSQL Zip" or, if you'd rather, "PGXN Zip". Either way, it's
  just a zip archive.

That processing done, with a good `META.json` and zip archive, the JSON,
username, and SHA1 of the zip archive are handed off to the database for more
processing. The [`add_distribution()`] database function does all the heavy
lifting here. It:

- Parses the JSON string, validates that all required keys are present, and
  normalizes version numbers. Yes, this is redundant, but I don't think I need
  to lecture the reads of this blog about database integrity. :-)

- Creates a new metadata structure and stores all the required and many of the
  optional meta spec keys, as well as the SHA1 of the distribution file, the
  date, and the user's nickname.

- Sets the "release_status" to "stable" if there was no status in the original
  JSON.

- Adds a "provides" section to the metadata if none was included in the
  original JSON. In such a case, it assumes that the distribution contains one
  extension and that it has the same name and version as the distribution
  itself.

- Validates that the uploading user is owner or co-owner of all provided
  extensions. If no one is listed as owner of one or more included extensions,
  the user will be assigned ownership. If the user is not owner or co-owner of
  any included extensions, an exception will be thrown.

- Records the distribution, extensions, and tags in the database.

Once all this work is done, `add_distribution()` returns all the JSON that
needs to be written to the mirror. These files make up the "index" on the
network, and include:

- The `META.json` file ([example]). This file is derived from the `META.json`
  included in the archive ([example][3]), but reflects all the normalization
  changes and added keys outlined above.
- A short JSON file summarizing all releases of the distribution
  ([example][4]).
- JSON files describing all releases of all extensions included in the
  distribution ([example][5]).
- A JSON file for the user, which includes data about all releases made by
  that user ([example][6]).
- JSON files for each tag associated with the distribution. Each of these
  files lists all the distributions the tag is associated with ([example][7]).

It also returns JSON for network statistics files. These are updated every
time a new release is uploaded:

- [`dist.json`] lists the 56 most recent releases and has a count of all
  distributions and of releases of those distributions.
- [`extension.json`] lists the 56 most recent extension releases and a count
  of all extensions on the network.
- [`user.json`] lists of the 56 most prolific users (based on the number of
  distributions) and a count of distributions and releases for each.
- [`tag.json`] lists the 56 most popular tags (measured by the number of
  distributions they're associated with) and a count of all tags on the
  network.
- [`summary.json`] has basic summary information about the network, which is
  just counts of distributions, releases, extensions, users, tags, and
  mirrors.

If you think that's a lot of data to be updated, you're right! But since
releases are relatively infrequent (a couple a day at the moment), it's best
to generate all this stuff as static files that are `rsync`ed to all mirrors.
In this way, *any* mirror can function as a very simple, lightweight REST API.
And indeed, that's just how the planned [PGXN client] will behave. The
[WWW::PGXN][8] already provides the interface it will use.

I guess that was a lot of information. Let this be a reference document for
interested hackers, then. The core functionality is all there, but there's a
[lot more to be done][]:

- Add a UI for users to assign extensions to co-owners (when more than one
  user is allowed to upload an extension) and to transfer ownership ([#6]).
- Add a user administration interface ([#7]).
- Add a UI where admins can transfer ownership of extensions when an owner
  disappears ([#8]).
- Add a callback architecture to trigger other actions on successful upload.
  These would include tweeting new releases (currently just handled in the
  controller) and running arbitrary command-line utilities (such as syncing
  the PGXN API server so it can be up-to-date as quickly as possible) ([#9]).
- Add support for "instant mirroring" via [`rrr`] ([#10]).
- Add an Atom feed for recent releases (basically a dupe of the distribution
  stats file) ([#17]).
- Add a command-line interface to the API. The HTTP server is written in such
  a way that it can act as an API serving JSON, but it's not well-tested and
  nothing uses it yet (that I'm aware of) ([#1]).
- Add UI for re-indexing a distribution ([#11]).
- Add ability to delete distributions ([#12])?
- Add an email gateway so that email addressed to `$nickname@pgxn.org` will be
  forwarded to a user's actual address. This would allow us to remove literal
  email addresses from the JSON files and the site (the site obfuscates them,
  but still...). Anyone got some good postfix chops for this? The [`users`
  table] is quite simple ([#13]).

Want to help out? Fork [PGXN Manager][9] and have at it. Hell, at this point
I'd really appreciate a code review, as I'm pretty sure there's only been one
set of eyes on this code so far.

Next week, I plan to blog about

- Project status: where the hours went and where they're going.
- [URI templates] and how they provide a method index for the mirror API
- Mirrors and mirroring via `rsync`
- [PGXN API][10]: how it provides a superset of the mirror API, including
  full-text search.
- [PGXN Site][11]: how it's just a very thing wrapper over the API (but can
  read the API directly from the file system!)
- Plans for the PGXN client

But given how these things go, and how I need to start writing mirror API and
API server documentation, it might take me a longer to get to them all.



  [mirrors]: https://api.pgxn.org/mirror/meta/mirrors.json
  [master mirror]: https://master.pgxn.org/
  [host a mirror]: https://pgxn.org/mirroring/
  [PGXN Manager]: https://manager.pgxn.org/
  [code]: https://github.com/pgxn/pgxn-manager/ "Fork PGXN::Manager on GitHub"
  [PGXN API]: https://api.pgxn.org
  [1]: https://github.com/pgxn/pgxn-api "Fork PGXN::API on GitHub"
  [PGXN Site]: https://pgxn.org
  [2]: https://github.com/pgxn/pgxn-site "Fork PGXN::Site on GitHub"
  [WWW::PGXN]: https://github.com/pgxn/www-pgxn "Fork WWW::PGXN on GitHub"
  [the spec]: https://pgxn.org/spec/
  [HOWTO]: https://manager.pgxn.org/howto
  [PGXN::Manager::Distribution]: https://github.com/pgxn/pgxn-manager/blob/master/lib/PGXN/Manager/Distribution.pm
    "PGXN::Manager::Distribution on GitHub"
  [Archive::Extract]: https://search.cpan.org/perldoc?Archive::Extract
  [CPAN::Meta::Validator]: https://search.cpan.org/perldoc?CPAN::Meta::Validator
  [semantic versions]: https://semver.org/
  [SemVer documentation]: https://search.cpan.org/~dwheeler/SemVer-v0.2.0/lib/SemVer.pm#Usage
  [`add_distribution()`]: https://github.com/pgxn/pgxn-manager/wiki/DB-API#add_distribution
  [example]: https://api.pgxn.org/mirror/dist/semver/0.2.1/META.json
  [3]: https://api.pgxn.org/src/semver/semver-0.2.1/META.json
  [4]: https://api.pgxn.org/mirror/dist/semver.json
  [5]: https://api.pgxn.org/mirror/extension/semver.json
  [6]: https://api.pgxn.org/mirror/user/theory.json
  [7]: https://api.pgxn.org/mirror/tag/datatype.json
  [`dist.json`]: https://api.pgxn.org/mirror/stats/dist.json
  [`extension.json`]: https://api.pgxn.org/mirror/stats/extension.json
  [`user.json`]: https://api.pgxn.org/mirror/stats/user.json
  [`tag.json`]: https://api.pgxn.org/mirror/stats/tag.json
  [`summary.json`]: https://api.pgxn.org/mirror/stats/summary.json
  [PGXN client]: https://github.com/pgxn/pgxn-client/
    "Fork PGXN::Client on GitHub"
  [8]: https://github.com/pgxn/www-pgxn/ "Fork WWW::PGXN on GitHub"
  [lot more to be done]: https://github.com/pgxn/pgxn-manager/issues/
  [#6]: https://github.com/pgxn/pgxn-manager/issues/6
  [#7]: https://github.com/pgxn/pgxn-manager/issues/7
  [#8]: https://github.com/pgxn/pgxn-manager/issues/8
  [#9]: https://github.com/pgxn/pgxn-manager/issues/9
  [`rrr`]: https://search.cpan.org/perldoc?File::Rsync::Mirror::Recent
  [#10]: https://github.com/pgxn/pgxn-manager/issues/10
  [#17]: https://github.com/pgxn/pgxn-manager/issues/17
  [#1]: https://github.com/pgxn/pgxn-client/issues/1
  [#11]: https://github.com/pgxn/pgxn-manager/issues/11
  [#12]: https://github.com/pgxn/pgxn-manager/issues/12
  [`users` table]: https://github.com/pgxn/pgxn-manager/blob/master/sql/03-users.sql
  [#13]: https://github.com/pgxn/pgxn-manager/issues/13
  [9]: https://manager.pgxn.org/ "For PGXN::Manager on GitHub"
  [URI templates]: https://tools.ietf.org/html/draft-gregorio-uritemplate-04
  [10]: https://api.pgxn.org/
  [11]: https://pgxn.org/
