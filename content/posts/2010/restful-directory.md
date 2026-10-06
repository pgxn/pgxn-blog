---
title: A RESTful Directory
slug: restful-directory
date: 2010-08-21T18:52:00Z
aliases: [/post/988613682/restful-directory, /post/988613682]
tags: [Mirror, Directory, Structure, Meta, Metadata, JSON, API, Scalability, REST]
---

Following my post outlining a possible network [directory structure],
[Aristotle Pagaltzis] saw fit to bug me via email about a different approach.
I couldn't understand WTF he was talking about until today. Then it lit my
brain on fire. As a result, I now think that there is a much better way to
organize the metadata files for the PGXN -- one that happens not to include
any symbolic links (which is something that [Andreas König] has been flagging,
via email, as a possible bottleneck).

First, the `/dist` directory will be the same as before. Releases of pgTAP
would be in:

    dist/p/pg/pgtap/pgtap-0.23.pgz
    dist/p/pg/pgtap/pgtap-0.23.json
    dist/p/pg/pgtap/pgtap-0.23.readme
    dist/p/pg/pgtap/pgtap-0.24.pgz
    dist/p/pg/pgtap/pgtap-0.24.json
    dist/p/pg/pgtap/pgtap-0.24.readme
    dist/p/pg/pgtap/pgtap-0.25.pgz
    dist/p/pg/pgtap/pgtap-0.25.json
    dist/p/pg/pgtap/pgtap-0.25.readme

The only change is that the `pgtap.json` symlink is gone.

Now, the new stuff. In the root directory will be a file, `index.json`, that
contains templates for URIs. It will look something like this:

    {
        "dist":   "/dist/$a/$ab/$dist/$dist-$version.pgz",
        "readme": "/dist/$a/$ab/$dist/$dist-$version.readme",
        "meta":   "/dist/$a/$ab/$dist/$dist-$version.json",

        "by-dist":      "/by/dist/$a/$ab/$dist.json",
        "by-extension": "/by/extension/$a/$ab/$extension.json",
        "by-owner":     "/by/owner/$a/$ab/$owner.json",
        "by-manager":   "/by/manager/$a/$ab/$manager.json",
    }

The PGXN client will always fetch this file before it does anything else,
because the file tells it how to find stuff. The advantage here is that the
client doesn't have to know anything about how the directory is actually
organized, just what the template variables might be. They are:

- `$dist`: A [distribution name]
- `$version`: A [version number]
- `$extension`: An [extension name]
- `$owner`: An [owner's name]
- `$manager`: A release manager's name (managers are the people who upload
  distributions to PGXN)
- `$a`: The first letter of a distribution, extension, owner, or manager name.
- `$ab`: The first two letters of a distribution, extension, owner, or manager
  name.

I'm not thrilled about using prefix-staggering to avoid having too many files
in a directory. But the truth is that this approach allows me to punt. I could
also make sure the client supports, for example, `$bc` and `$cd`, so that one
could stagger things differently. And then the nice thing is that I don't have
to use those at all. The templates will tell the client exactly how to
construct the URIs for things, and the templates needn't include those
staggering variables if they're not appropriate. The client won't care because
it will have no built-in knowledge of how things are organized. It will have
to find out from `index.json`.

From the URI templates, you can now see where the other metadata will be
stored. For extension names, a hypothetical pgTAP distribution with two
extensions will have a JSON file for each extension:

    /by/extension/p/pg/pgtap.json
    /by/extension/s/sc/schematap.json



The `pgtap.json` file will look something like this:

    "stable":   "0.25.0",
    "testing":  "0.26.0b1",
    "unstable": "0.30.0u",
    "versions": {
        "0.26.0b1": { "dist": "pgtap", "version": "0.26.0b1", "status": "testing" },
        "0.30.0u":  { "dist": "pgtap", "version": "0.30.0u",  "status": "unstable" },
        "0.25.0":   { "dist": "pgtap", "version": "0.25.0",   "status": "stable" },
        "0.24.0":   { "dist": "pgtap", "version": "0.24.0",   "status": "stable" },
        "0.25.0":   { "dist": "pgtap", "version": "0.23.0",   "status": "stable"  }
    }

Right at the top, it would always list the most recent stable, testing, and
unstable version number, and then it would have a list metadata for all
versions. Said metadata would include the associated distribution name,
version, and release status.

Here's how it would work. Say I ask the client to install pgtap:

    PGXN> install extension pgtap

The client would first fetch `/index.json`, then look for the URI template for
"by-extension", which is `/by/extension/$a/$ab/$extension.json`. Filling in
the template, it would know to request `/by/extension/p/pg/pgtap.json`. With
that file, it would see that the most recent stable version is in the "pgtap"
distribution version 0.25.0. Using the `dist` URI template, which is
`/dist/$a/$ab/$dist-$version.pgz`, it would then fetch
`/dist/p/pg/pgtap/pgtap-0.25.0.pgz`.

The advantage here is that there are no symbolic links and no knowledge of the
directory structure built into clients. The client just knows to fetch
`/index.json` and then to use the templates in that file to fetch other
information. That's the whole interface. Very [REST]ful.

The structure of the other `/by` files would be similar. For

    PGXN> install dist pgtap

the client would use the "by-dist" URI template to construct the URL
`/by/dist/p/pg/pgtap.json`. That file would have something like:

    "stable":   "0.25.0",
    "testing":  "0.26.0b1",
    "unstable": "0.30.0u",
    "versions": {
        "0.26.0b1": "testing",
        "0.30.0u":  "unstable",
        "0.25.0":   "stable",
        "0.24.0":   "stable" ,
        "0.23.0":   "stable"
    }

So then the client would know that "0.25.0" was the most recent version, and
use the `dist` URI template to request `/dist/p/pg/pgtap/pgtap-0.25.0.pgz`.

If The client command had been:

    PGXN> readme dist pgtap

It would use the `readme` URI template. And the command:

    PGXN> meta dist pgtap

Would use the `meta` URI template to fetch the metadata for the distribution.

If the client had requested a specific version:

        PGXN> install dist 0.23.0

It could either use the `by-dist` URI template to download the list of all
versions to see if 0.23.0 was valid, or just use the `dist` URI template to
try to download the distribution itself.

And finally, the owner and manager JSON files, such as

    /owner/t/th/theory.json

Would look something like:

    "full_name": "David Wheeler",
    "email": "theory@pgxn.org",
    "uri": "https://justatheory.com",
    "distributions": {
        "pgtap": [ "0.25.0", "0.24.0", "0.23.0" ]
        "pair": [ "0.2.0", "0.1.0", "0.0.5" ]
    }

With that, the client can be asked to fetch metadata for a given owner name
and use it to figure out what distributions and versions the the owner, um,
owns. One could then fetch the metadata, readme, or distribution file for any
of those distributions and versions.

Overall, I think that this is a much better solution than I outlined
[before][directory structure]. If only I could figure out something more
elegant that the prefix-staggering/hashing stuff, it would be just about
perfect.

Thoughts?



  [directory structure]: https://blog.pgxn.org/post/954535657/thoughts-on-the-network-directory-structure
  [Aristotle Pagaltzis]: https://plasmasturm.org/
  [Andreas König]: https://search.cpan.org/~andk/
  [distribution name]: https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec#name
  [version number]: https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec#version
  [extension name]: https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec#provides
  [owner's name]: https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec#owner
  [REST]: https://en.wikipedia.org/wiki/REST
