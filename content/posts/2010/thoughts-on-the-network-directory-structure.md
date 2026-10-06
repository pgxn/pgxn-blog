---
title: Thoughts on the Network Directory Structure
slug: thoughts-on-the-network-directory-structure
date: 2010-08-14T23:59:12Z
aliases: [/post/954535657/thoughts-on-the-network-directory-structure, /post/954535657]
tags: [Mirror, Directory, Structure, Metadata, JSON, API, Scalability, REST]
---

I've been thinking about the arrangement of stuff to be distributed to the
mirrors. In doing so, I've kept three goals in mind:

1.  Allow things to scale.
2.  Try to put things in intuitive locations.
3.  Allow for metadata to be stored to act as a Web API.

The first two are a bit mutually-contradicting. Thinking scalability mainly
means minimizing the chances that too many files can be put into a single
directory. [CPAN] started out with all author directories in a single
directory. Given the sheer number of CPAN contributors, this quickly got to be
a bottleneck. To correct for that, they started "hashing" the first two
letters of author names to create subdirectories. My CPAN directory, for
example, is [`D/DW/DWHEELER`]. That doesn't quite make for an intuitive
location, but it's not bad, and the tradeoff seems sufficient to keep things
sane.

The storage of metadata is important, too, as I plan to have the [client] send
requests for JSON files to mirrors in order to find distributions. Basically,
this came down to a naming convention, as well as a recipe for how to find
metadata files for people and extensions. What I've come up with is three
directories:

- `meta` will contain PGXN metadata
- `dist` will contain distributions
- `by` will contain query-able JSON files

The `meta` directory will contain PGXN metadata (`mirrors.json`, timestamp,
other detritus). If I ever create an index file that lists all distributions,
it would go there, too. I'm not going to discuss it any further here, except
to note that it already has `mirrors.json`, which clients will be able to use
to present users with a list of mirrors to choose from.

The `dist` directory will be organized with directories named with "hashes" of
the first two letters of distribution names. Let's say that I'm releasing
[pgTAP] 0.25 on PGXN. To distribute it, the management application will create
the directory (if it doesn't already exist):

    dist/p/pg/pgtap/

Then, for the 0.25 release, it will add three files to that directory:

    dist/p/pg/pgtap/pgtap-0.25.pgz
    dist/p/pg/pgtap/pgtap-0.25.json
    dist/p/pg/pgtap/pgtap-0.25.readme

The `pgtap-0.25.pgz` file will contain the [zipped distribution] ready for
download. `pgtap-0.25.json` will contain metadata about the distribution, such
as the owner's name, the manager's name (more on these folks below), list of
included extensions, location of the `.pgz` and `.readme`, and its SHA1.
`pgtap-0.25.readme` will of course contain the `README` for the distribution
(if it has one).

Every release of pgTAP will have these three files, so after several releases,
the `pgtap` directory might have these files:

    dist/p/pg/pgtap/pgtap-0.23.pgz
    dist/p/pg/pgtap/pgtap-0.23.json
    dist/p/pg/pgtap/pgtap-0.23.readme
    dist/p/pg/pgtap/pgtap-0.24.pgz
    dist/p/pg/pgtap/pgtap-0.24.json
    dist/p/pg/pgtap/pgtap-0.24.readme
    dist/p/pg/pgtap/pgtap-0.25.pgz
    dist/p/pg/pgtap/pgtap-0.25.json
    dist/p/pg/pgtap/pgtap-0.25.readme
    dist/p/pg/pgtap/pgtap.json

The last file there, `pgtap.json`, will actually be a symlink to the JSON file
for latest production release of pgTAP. In this case, it would link to
`pgtap-0.25.json`. The nice thing about this is that, if a client wants to
find the information about the latest release of the pgtap distribution, all
it will have to do is send an HTTP GET request for
`dist/p/pg/pgtap/pgtap.json` to any mirror.

The `by` directory will also contain JSON files for clients to request. The
idea is that you want to find information "by" something. To start with, there
will be three subdirectories:

    by/extension/
    by/manager/
    by/owner/

The first directory, `by/extension/`, will contain links to JSON files for
extensions. Say that the pgTAP distribution offers two extensions to
PostgreSQL named "pgtap" and "schematap". The links would be:

    by/extension/p/pg/pgtap/pgtap-0.23.json
    by/extension/p/pg/pgtap/pgtap-0.24.json
    by/extension/p/pg/pgtap/pgtap-0.25.json
    by/extension/p/pg/pgtap/pgtap.json

Each of these will simply be symlinks pointing to the appropriate distribution
files:

    dist/p/pg/pgtap/pgtap-0.23.json
    dist/p/pg/pgtap/pgtap-0.24.json
    dist/p/pg/pgtap/pgtap-0.25.json
    dist/p/pg/pgtap/pgtap.json

Yes, the last one is a symlink to a symlink. Similarly, these files for
"schematap":

    by/extension/s/sc/schematap/schematap-0.23.json
    by/extension/s/sc/schematap/schematap-0.24.json
    by/extension/s/sc/schematap/schematap-0.25.json
    by/extension/s/sc/schematap/schematap.json

Point to exactly the same files. Essentially, this is a way for extensions to
point to the distributions that contain them.
`by/extension/s/sc/schematap/schematap.json` points to
`dist/p/pg/pgtap/pgtap.json`, which contains path to the `.pgz` file to
download (and lots of other metadata, too).

The idea is that, whatever the name of the extension you want, the client will
be able to easily find the metadata file that tells it where to find the
distribution.

The `by/manager` and `by/owner` directories, on the other hand, contain JSON
files with information about managers and owners and their distributions.
Definitions:

An "owner" is someone who *owns* a distribution. This will often be the
original author of an extension, but may be someone else if maintenance has
been passed on. Basically, the "owner" is the person who should be contacted
with bug reports and the like

A "manager" is someone who manages the release process. This is the user who
will log into the PGXN management application and upload a distribution for
release.

These two people will often be the same person, but not always. I've avoided
the term "author" (the term used by CPAN) because the author of an extension
may no longer maintain it. "Owner" seemed like a better choice (individual
distributions are free to describe their contributors however they wish).

As the *owner* of a few extensions on PGXN, I'd have this file:

    by/owner/t/th/theory.json

This file would contain a list of my distributions and perhaps some other
information (like my full name and blog URL).

As the *manager* of extensions on PGXN (that is, I actually uploaded them), I
would also have:

    by/manager/t/th/theory.json

This file would contain a list of the distributions I've released on PGXN.
This might be exactly the same as the list in my owner file, but may not be.
Perhaps for one release of pgTAP, say 0.26, [Duke Leto] uploaded a release. In
that case, 0.26 would probably be in my owner file, but not in my manager
file: it would be in Duke's manager file, instead.

So these are the basics of the directory structure for the networked mirrors.
Note that I haven't thought much about everything that will go into the JSON
files (or whether or not they'd be versioned). That will likely depend quite a
lot on what the management database ends up looking like. I'll be working on
that next.

But other than that, comments? Questions? Criticisms? Recommendations? Leave a
comment and let me know!



  [CPAN]: https://www.cpan.org/misc/ZCAN.html
    "Zen of Comprehensive Archive Networks (look for “Naming”)"
  [`D/DW/DWHEELER`]: https://www.cpan.org/authors/id/D/DW/DWHEELER/
  [client]: https://pgxn.org/status.html#client
  [pgTAP]: https://pgtan.org/
  [zipped distribution]: https://wiki.postgresql.org/wiki/PGXN#Distribution_Layout
  [Duke Leto]: https://leto.net/dukeleto.pl/
