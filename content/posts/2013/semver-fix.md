---
title: PGXN Distribution Prerelease Versions Fixed
slug: semver-fix
date: 2013-06-19T19:24:10Z
aliases: [/post/53377271882/semver-fix, /post/53377271882]
tags: []
---

A couple years ago, the [Semantic Version] specification was [updated] to
require a hyphen between the "patch version" and the "prerelease version."
This despite the fact that v1.0.0 of the spec had been published on the site
for a while (see all the gory issues [here]). Naturally, I had already
implemented that format for PGXN in a [Perl module] and [Postgres data type].
Since we already had some PGXN extensions with the non-hyphenated format, and
I knew supporting the hyphen would break them, I held off updating those
implementations for a year or so.

I finally broke down and updated them, though, to be fully compliant with the
official v1.0.0 spec, and pushed the changes to the PGXN server a few months
ago. And I was right: it *did* break things. Links broke, downloads failed,
and my mailbox filled up with error messages. I triaged the worst of the
issues, but since then, accessing PGXN distributions with prerelease versions
has not worked so well.

Until last night. I finally got some tuits in the last month to dig into this
issue and figure out how to fix it. As a result, there are new releases of
[PGXN Manager], [the API], and [the site], as well as [the `META.json`
validator] that should keep things consistent and working well going forward.
Assuming there are no more backwards-incompatible changes to semantic
versions, of course.

But that still didn't solve the problem of existing distributions with the old
format of semantic version. Alas, I had to download them all, modify them, and
re-index them. This means that, if you had URLs to prerelease versions laying
around, they won't work anymore. Fortunately, they're prerelease versions, so
I wouldn't expect many to have them lying around anyway.

Still in the interest of disclosure (and because their SHA1s have changed),
here's a list of the changed distributions:

- pgmp 1.0.0-b2 changed to [1.0.0-b2]
- pgmp 1.0.0-b3 changed to [1.0.0-b3]
- plv8 1.1.0beta1 changed to [1.1.0-beta1]
- pgvihash 1.0.0dirty changed to [1.0.0-dirty]
- madlib 0.3.0alpha1 changed to [0.3.0-alpha1]
- madlib 0.5.0release1 changed to [0.5.0-release1]
- madlib 0.6.0release1 changed to [0.6.0-release1]
- multicorn 1.0.0beta1 changed to [1.0.0-beta1]
- pg_repack 1.1.8alpha1 changed to [1.1.8-alpha1]
- pg_repack 1.1.8beta1 changed to [1.1.8-beta1]
- pg_repack 1.1.8beta2 changed to [1.1.8-beta2]

Apologies for the dead links, and for taking so long to get this stuff cleaned
up.



  [Semantic Version]: https://semver.org/
  [updated]: https://github.com/mojombo/semver/commit/05a00df02db3d4ba83e0caaff6b31c10c77f7d3d
  [here]: https://github.com/mojombo/semver/issues/49#issuecomment-19705479
  [Perl module]: https://metacpan.org/module/SemVer
  [Postgres data type]: https://pgxn.org/extension/semver/
  [PGXN Manager]: https://manager.pgxn.org/
  [the API]: https://api.pgxn.org/
  [the site]: https://pgxn.org/
  [the `META.json` validator]: https://metacpan.org/module/PGXN::Meta::Validator
  [1.0.0-b2]: https://pgxn.org/dist/pgmp/1.0.0-b2/
  [1.0.0-b3]: https://pgxn.org/dist/pgmp/1.0.0-b3/
  [1.1.0-beta1]: https://pgxn.org/dist/plv8/1.1.0-beta1/
  [1.0.0-dirty]: https://pgxn.org/dist/pgvihash/1.0.0-dirty/
  [0.3.0-alpha1]: https://pgxn.org/dist/madlib/0.3.0-alpha1/
  [0.5.0-release1]: https://pgxn.org/dist/madlib/0.5.0-release1/
  [0.6.0-release1]: https://pgxn.org/dist/madlib/0.6.0-release1/
  [1.0.0-beta1]: https://pgxn.org/dist/multicorn/1.0.0-beta1/
  [1.1.8-alpha1]: https://pgxn.org/dist/pg_repack/1.1.8-alpha1/
  [1.1.8-beta1]: https://pgxn.org/dist/pg_repack/1.1.8-beta1/
  [1.1.8-beta2]: https://pgxn.org/dist/pg_repack/1.1.8-beta2/
