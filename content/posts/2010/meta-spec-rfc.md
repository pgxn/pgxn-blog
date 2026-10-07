---
title: "RFC: The PGXN Metadata Specification"
slug: meta-spec-rfc
date: 2010-08-17T13:00:00Z
aliases: [/post/966989188/meta-spec-rfc, /post/966989188]
tags: [Meta, Meta Spec, Metadata, JSON, CPAN Meta Spec, David Golden, Version Numbers, Prerequisites, License]
---

I've posted a draft of the "PGXN distribution metadata specification," or
[PGXN Meta Spec]." This document specifies the structure and format of the
`META.json` file that PGXN will require in every distribution. In fact, this
is the *only* required file in a distribution. Its job is to describe the
distribution, its extensions, and its dependencies, among other things. This
file is key to the whole thing.

To create it, I've ported the [CPAN Meta Spec], version 2, but with all
deprecated fields removed, and some of the more complex stuff taken out. I
also made a couple of the "required" fields "optional." At its simplest, the
file might look something like this:

```json
{
  "name": "pgTAP",
  "abstract": "Unit testing for PostgreSQL",
  "version": "0.25.0",
  "owner": "David E. Wheeler <theory@pgxn.org>",
  "license": "postgresql",
  "meta-spec": {
    "version": "1.0.0",
    "url": "https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec"
  }
}
```

Not too bad, eh? The URL for the spec might change (might move it to the main
site and/or the mirrors), but otherwise, I think this is pretty solid. Not too
much work to deal with, and reasonably easy to create by hand (which is likely
how we'll all start out).

And additional key that will be really important for the PGXN client is
`prereqs`. This key allows you to identify the prerequisites (from PGXN or the
PostgreSQL core contrib extensions) required to build, test, and/or use a
distribution. For example, if I were to release an [ordered pair] extension,
it of course would include tests written with [pgTAP]. So I'd have something
like:

```json
{
  "name": "pair",
  "abstract": "An ordered pair data type",
  "version": "0.1.0",
  "owner": "David E. Wheeler <theory@pgxn.org>",
  license: "postgresql",
  "meta-spec": {
    "version": "1.0.0",
    "url": "https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec"
  },
  "prereqs": {
    "runtime": {
      "requires": {
        "PostgreSQL": "8.0.0"
      },
      "recommends": {
        "PostgreSQL": "8.4.0"
      }
    },
    "test": {
      "requires": {
        "pgTAP": 0
      }
    }
  }
}
```

That's saying that the "pair" distribution requires PostgreSQL 8.0.0 or higher
and any version of pgTAP to run the test suite. I've also recommended
PostgreSQL 8.4, as that's where it will run best.

Of course, to get the real power of PGXN, you'll also want to use the
`provides` key, which allows you to identify the extensions included in your
distribution. Say that I finally got around to breaking out the schema testing
assertions from the logical testing assertions in pgTAP. I might call the
second module "schematap." So to spell it out, I'd add this to the first
example above:

```json
  "pgtap": {
    "file": "sql/pgtap.sql.in",
    "version": "0.25.0"
  },
  "schematap": {
    "file": "sql/schematap.sql.in"
  }
```
So now the indexer will know that the "pgtap" extension is in
`sql/pgtap.sql.in` and the "schematap" extension is in `sql/schematap.sql.in`.
This is important because it allows other distributions to specify "schematap"
as a prerequisite. It also means that, in the PGXN client, you can type
something like:

```
PGXN> install schematap
```

And, because "schematap" will have been indexed on the network, the client
will be able to find the pgTAP distribution and install it, complete with the
"schematap" extension.

A final note. Version numbers in the Perl community are a [disaster]. Wanting
to avoid that whole morass, I had originally intended to require numeric
version numbers. But [David Golden] -- the current maintainer of the CPAN Meta
Spec -- pointed me to [Semantic Versioning], a version number specification by
GitHub's [Tom Preston-Werner]. This style of version numbering is great for
PGXN for a few reasons:

- It closely matches how PostgreSQL itself is versioned.
- It's very easy to compare version strings.
- Someone else has already dealt with the pain of writing a spec

So this is the standard that PGXN will require. Every version number will be
dotted-integer with three integers (X.Y.Z) and an optional ASCII string at the
end. That's it. PGXN won't invest any special meaning in the version string
the way CPAN does. It will just compare version numbers.

Anyway, please review [the spec][PGXN Meta Spec] itself and leave any comments
or questions below. I expect to start hacking on this stuff this week!

**Update 2010-08-24:** I've just updated [the spec][PGXN Meta Spec] to change
"owner" to "maintainer." I think that the latter term is much better for this
purpose. And then I can use "owner" in PGXN to identify the person who uploads
a distribution.

  [PGXN Meta Spec]: https://github.com/theory/pgxn/wiki/PGXN-Meta-Spec
  [CPAN Meta Spec]: https://search.cpan.org/perldoc?CPAN::Meta::Sepc
  [ordered pair]: https://justatheory.com/computers/databases/postgresql/key-value-pairs.html
  [pgTAP]: https://pgtap.org/
  [disaster]: https://www.dagolden.com/index.php/369/version-numbers-should-be-boring/
    "Version numbers should be boring"
  [David Golden]: https://www.dagolden.com/
  [Semantic Versioning]: https://semver.org/
  [Tom Preston-Werner]: https://tom.preston-werner.com/
