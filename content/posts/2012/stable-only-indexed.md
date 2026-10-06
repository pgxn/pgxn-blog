---
title: Only Stable Releases are Indexed
slug: stable-only-indexed
date: 2012-04-06T17:02:19Z
aliases: [/post/20596200256/stable-only-indexed, /post/20596200256]
tags: [Release, Status, Release Status, Stable, Testing, Unstable, Full Text Index, FTI, Search Results, Adaptive Estimator]
---

I've had a few reports over the last few months that folks uploaded new
extensions to [PGXN] but they failed to show up in search results. For
example, as of right now, if you [search for "distinct"], you get two results,
[OmniPITR] and [pgTAP]. This despite the fact that the recently-released
[adaptive_estimator] extension is [tagged "distinct"]. Author [Tomas Vondra]
reported this issue, and it took me a couple of days to realize why it wasn't
showing up.

Here's the reason: Only "stable" releases are indexed. The [first] and
[second] adaptive_estimator releases both have their release status set to
"testing". The PGXN API only indexes stable releases. The idea behind that is
that you want most folks to continue using stable releases while you work on
testing new versions. So when users search the site, only the latest stable
release will appear in search results. Similarly, installing an extension via
the [PGXN client], prefers the latest stable release by default. If you want
the most recent, you have to [specify the `--unstable` or `--testing` option].

So, the upshot is, if you want your extension to appear in the full text
search results on the PGXN, site, release a stable version.

That said, I first noticed this issue a while ago, and filed [an issue] with
the idea that, if an extension is uploaded that is not stable, but there are
no stable versions, then the extension should be indexed, anyway. The idea
here is that, when you first upload it, you don't have any existing users to
keep on a stable release anyway, because there *is no stable release.* So
perhaps we should go ahead and index it. A non-stable release would only be
omitted from the index if there was an existing stable release.

Thoughts?



  [PGXN]: https://pgxn.org/
  [search for "distinct"]: https://pgxn.org/search?q=distinct&in=docs
  [OmniPITR]: https://pgxn.org/dist/omnipitr/doc/internals.html
  [pgTAP]: https://pgxn.org/dist/pgtap/doc/pgtap.html
  [adaptive_estimator]: https://pgxn.org/dist/adaptive_estimator/
  [tagged "distinct"]: https://pgxn.org/tag/distinct/
  [Tomas Vondra]: https://pgxn.org/user/tomasv
  [first]: https://pgxn.org/dist/adaptive_estimator/1.0.0/
  [second]: https://pgxn.org/dist/adaptive_estimator/1.1.0/
  [PGXN client]: https://pgxnclient.projects.postgresql.org/
  [specify the `--unstable` or `--testing` option]: https://pgxnclient.projects.postgresql.org/usage.html#pgxn-install
  [an issue]: https://github.com/pgxn/pgxn-api/issues/2
