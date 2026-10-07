---
title: PGXN Tools Docker Image Updated
slug: pgxn-tools-v4
date: 2024-01-31T16:52:19Z
aliases: [/post/741049567045468160/pgxn-tools-v4, /post/741049567045468160]
tags: [PGXN, Docker, GitHub Actions, GitHub Workflows, Release, Bundle]
---

Just a quick note to highlight some bug fixes and improvements in the
[pgxn/pgxn-tools Docker image] in the last few weeks.

v1.4.0 adds the `GIT_BUNDLE_OPTS` and `ZIP_BUNDLE_OPTS` environment variables.
The former is passed to `git archive` and the latter to zip, and let
additional options be passed to those commands. For example, the [vectorize
extension]'s [release workflow] sets `GIT_BUNDLE_OPTS: --add-file META.json`
because git archive archives only checked-in files, and `META.json` is not
checked in but [generated from] `META.json.in`.

v1.4.1 [fixes an issue] where `git archive` was never actually used to build a
release zip archive. This changed at some point without noticing due to the
introduction of the `safe.directory` configuration in recent versions of Git.
Inside the container the directory was never trusted, and the `pgxn-bundle`
command caught the error, decided it wasn't working with a Git repository, and
used the `zip` command, instead.

As a result, a number of recent releases have included files they shouldn't,
such as the contents of the `.git` directory. [The fix][fixes an issue]
[disables] `safe.directory` so that the repository directory is always trusted
inside the container, as needed in GitHub actions.

If you use `pgxn-bundle` in a GitHub workflow, be aware that recent releases
(since November 2021 at least) include stuff that it should not, including the
`.git` directory (here's [a list]) and patterns excluded in `.gitattributes`.
Your next release should be cleaner.

v1.4.1 also allows the setting of the `PROFILE` environment variable, so that
the default `PROFILE=--Werror` can be overridden.

v1.4.2 adds (and v1.4.3 improves, see below) `git-archive-all` to the image,
and the `GIT_ARCHIVE_CMD` environment variable to tell `pgxn-bundle` which archive
command to use, either `archive` or `archive-all`. The latter is useful for
repositories that use Git submodules and need to include their contents in the
release, as demonstrated in [this plv8 pull request].

v1.4.2 also adds [`cmake`] and the `libarchive-tools` package to the image.
The latter bundles in [libarchive] tools like `bsdtar` and `bsdcpio`, which
might be useful for editing Zip files in place.

v1.4.3 passes the `--force-submodules` option to `git-archive-all`, because
otherwise submodules weren't included in the bundle.

All of this should just start showing up in your GitHub workflows as long as
they use `container: pgxn/pgxn-tools` or `container: pgxn/pgxn-tools@v1`.

Questions? Post em in the [#extensions] channel on the [Postgres Slack] or to
the attention of the [PGXN Mastodon bot].

  [pgxn/pgxn-tools Docker image]: https://github.com/pgxn/docker-pgxn-tools/
  [vectorize extension]: https://pgxn.org/dist/vectorize/
  [release workflow]: https://github.com/tembo-io/pg_vectorize/blob/v0.9.0/.github/workflows/pgxn-release.yml
  [generated from]: https://github.com/tembo-io/pg_vectorize/blob/v0.9.0/Makefile#L17-L18
  [fixes an issue]: https://github.com/pgxn/docker-pgxn-tools/commit/1114aed
  [disables]: https://stackoverflow.com/a/73100228/79202
  [a list]: https://gist.github.com/theory/93c93571200aad02e93170c6d2c93cbe
  [git-archive-all]: https://github.com/Kentzo/git-archive-all
  [this plv8 pull request]: https://github.com/plv8/plv8/pull/570/files#diff-4e4745c3f39dff4c307faf79bfa1242cc5ba2eb5947bebdcddc8d97eeb40685fR15
  [`cmake`]: https://cmake.org/
  [libarchive]: https://libarchive.org/
  [#extensions]: https://postgresteam.slack.com/archives/C056ZA93H1A
  [Postgres Slack]: https://pgtreats.info/slack-invite
  [PGXN Mastodon bot]: https://mastodon.social/@pgxn
