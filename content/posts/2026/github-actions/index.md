---
title: Build, Test and Release Extensions with the New PGXN GitHub Actions
slug: github-actions
date: 2026-11-16T20:14:17Z
description: |
  Introducing three new GitHub actions for more efficient, multi-platform
  testing of extensions on a variety of Postgres versions. Release, too!
tags: [GitHub Actions, Postgres, Matrix, Release, Workflow]
type: posts
---

I've made three new GitHub actions in the past few months to replace the
venerable Docker-based [pgxn/pgxn-tools], which has become bloated and
requires more work to maintain than I'd like. The new actions require much
less maintenance, support all GitHub runner OSes and architectures, and run
more quickly. What they do:

*   [pgxn/postgres-action] starts a local, specific version of Postgres server
*   [pgxn/matrix-action] generates a matrix of Postgres versions to test
*   [pgxn/release-action] releases an extension on PGXN

## Make a Release

Starting simple, here's how to use the release action:

```yaml
name: 🚀 Release
on:
  push:
jobs:
  release:
    name: Release on PGXN
    runs-on: ubuntu-latest
    steps:
        uses: actions/checkout@v7
        uses: pgxn/release-action@v0
        with:
          username: ${{ secrets.PGXN_USERNAME }}
          password: ${{ secrets.PGXN_PASSWORD }}
          dry-run: ${{ !startsWith( github.ref, 'refs/tags/v' ) }}
```

Yep, that's it. The [pgxn/release-action] action will use `git archive` to
create a PGXN bundle and release it on PGXN. (Want to omit some files, such as
GitHub workflows, from the release? [Use gitattributes] to exclude them!) The
value passed for the `dry-run` option prevents publishing the bundle on PGXN
except when pushing a Git tag starting with `v`.

The action also outputs the `archive` name, so you can use it elsewhere, such
as a GitHub release on the last line of this workflow:

```yaml
name: 🚀 Release
on:
  push:
jobs:
  release:
    name: Release on GitHub and PGXN
    runs-on: ubuntu-latest
    steps:
      - name: Check out the repo
        uses: actions/checkout@v7
      - name: Release on PGXN
        uses: pgxn/release-action@v0
        id: pgxn
        with:
          username: ${{ secrets.PGXN_USERNAME }}
          password: ${{ secrets.PGXN_PASSWORD }}
          dry-run: ${{ !startsWith( github.ref, 'refs/tags/v' ) }}
      - name: Create GitHub Release
        uses: softprops/action-gh-release@v3
        if: startsWith( github.ref, 'refs/tags/v' )
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          name: "Release ${{ github.ref_name }}"
          files: ${{ steps.pgxn.outputs.archive }}
```

It again checks `github.ref` to only release on a tag starting with `v`. Then
it includes the `archive` output from the `pgxn` step in the release. Simple,
right? Want to include a changelog in the release? Consider using a [Keep a
Changelog] file and [theory/changelog-version-notes-action] to generate a file
to pass to the [softprops/action-gh-release] `body_path` option.

## Test Supported Versions & Platforms

This example, [from pg-hostname], uses the [pgxn/matrix-action] action to
generate the test matrix, and [pgxn/postgres-action] to start Postgres. Then
it builds and tests the extension:

```yaml {linenos=table}
name: ✅ CI
on:
  push:
defaults:
  run: { shell: bash }
jobs:
  matrix:
    name: 📋 Generate Matrix
    runs-on: ubuntu-latest
    outputs:
      matrix: ${{ steps.matrix.outputs.matrix }}
    steps:
      - name: Generate Matrix
        id: matrix
        uses: pgxn/matrix-action@v0
        with:
          min-version: 9.0
          beta: true
          deprecated: true
          no-os: windows

  build:
    needs: matrix
    strategy:
      fail-fast: false
      matrix:
        pg: ${{ fromJson(needs.matrix.outputs.matrix) }}
    name: ${{ matrix.pg.emoji }} ${{ matrix.pg.os }}/${{ matrix.pg.arch }} 🐘 v${{ matrix.pg.version }}
    runs-on: ${{ matrix.pg.runner }}
    steps:
      - uses: actions/checkout@v7
      - uses: pgxn/postgres-action@v0
        with: { version: "${{ matrix.pg.version }}" }
      - run: make -j $(nproc)
      - run: sudo make install
      - run: make installcheck
        id: test
      - name: Show Diffs
        if: failure() && steps.test.outcome == 'failure'
        run: find . -name regression.diffs -exec cat {} +
      - uses: actions/upload-artifact@v7
        if: failure() && steps.test.outcome == 'failure'
        with:
          name: regression-${{ matrix.pg.os }}-${{ matrix.pg.arch }}-pg${{ matrix.pg.version }}
          path: |
            regression.diffs
            results
```

{{% figure
    src       = "jobs.png"
    alt       = "Screenshot of list of passing GitHub Workflow jobs for a variety of OSes, architectures, and Postgres versions."
    caption   = "Workout."
    class     = "float-right quarter"
%}}

The resulting output resembles the image on to the right.

I admit it seems like a lot; let's break it down:

*   The `matrix` job starting on line 7 generates the test matrix
    *   Lines 10-11 define the output for the matrix
    *   Lines 13-20 use [pgxn/matrix-action] to generate the matrix
    *   The options on line 17-20 set the minimum Postgres version to 9.0,
        allow the inclusion of beta and deprecated releases (e.g, those
        provided by but deprecated by the packager), and exclude `windows`
*   The `build` job starting on line 22 builds and tests the extension
    *   Line 23 requires that the `matrix` job finish before this one starts
    *   Line 27 generates the `pg` matrix, using `fromJson()` to extract the
        output of the `matrix` job
    *   Various matrix fields construct the `name` on line 28
    *   Line 29 specifies the runner provided by the matrix
    *   Lines 32-33 use the [pgxn/postgres-action] to start the
        matrix-specified version of Postgres on the runner --- no more Docker
        container!
    *   Lines 34-37 define the typical [PGXS] pattern to build and test the
        extension; note the use of `sudo` to install the extension, required
        because the Linux runner does not own the Postgres install (and is
        harmless on other OSes)
    *   The `Show Diffs` step starting on line 38 shows the contents of
        `regression.diffs` if the test failed
    *   And the upload action from line 41 on upload `regression.diffs` and
        the test result files if the test failed

A daily [pgxn/matrix-action] workflow updates its [list of all candidate
configurations][platforms.json], so it should always be up-to date.
Practically that means you automatically start testing the next Postgres
version within a day of its first beta release, unless you set `max-version`.

## Install, Build, Then Start

The [pgxn/matrix-action] and [pgxn/postgres-action] actions feature a number
of input options to define the runners, Postgres versions, and various
configurations.

Let's say that your extension requires loading a module in
[`shared_preload_libraries`], but you have to install Postgres to build the
module first. Use `no-start` to install but not start Postgres, then build the
extension, and finally use [pgxn/postgres-action] again with the desired
configuration:

```yaml
# Install but don't start Postgres
- uses: pgxn/postgres-action@v0
  with:
    version: ${{ matrix.pg.version }}
    no-start: true
# Build and install the extension.
- run: make -j $(nproc) && sudo make install
# Now start Postgres with the module in shared_preload_libraries
- uses: pgxn/postgres-action@v0
  with:
    version: ${{ matrix.pg.version }}
    config: |
      shared_preload_libraries = 'my_module'
# Continue testing
- run: make installcheck
```

## To Dos

I've been using these actions for a few months with a variety of extensions.
Some examples:

*   [semver]
*   [pg_clickhouse]
*   [chdb]
*   [hostname]

But there remain a few things left to do:

*   Proper Windows support. The [pgxn/postgres-action workflow] successfully
    tests a SQL-only extension on Windows Postgres 17-18, but not earlier
    versions due to a [path issue]. C extensions simply don't work at all. I'm
    not sure whether the EDB Windows Postgres binaries installed by
    [Chocolatey] need adjusting or if the issues can be fixed by further
    configuration. Help wanted!
*   I'd like to add a [pgrx] action, probably separate from the others. I'll
    likely get to it sometime in the next few months unless someone beats me
    to it.

## Thanks & Enjoy

Thanks for reading, I hope you find the new actions useful.

Props to [Peter Eisentraut] for demonstrating the use of the [PGDG Apt] repo
for starting Postgres on Linux runners, to the
[ikalnytskyi/action-setup-postgres] action for how to install and start
Postgres on macOS and Windows runners, and to Apt maintainer [Christoph Berg]
for patiently answering my questions about starting and running Postgres on
Linux.

  [pgxn/pgxn-tools]: https://github.com/pgxn/docker-pgxn-tools
    "Test image for PostgreSQL & PGXN extensions"
  [pgxn/matrix-action]: https://github.com/pgxn/matrix-action/
    "Generate a matrix of Postgres versions on Workflow OSes and architectures"
  [pgxn/postgres-action]: https://github.com/pgxn/postgres-action/
    "Install, configure, and start a local Postgres cluster"
  [pgxn/release-action]: https://github.com/pgxn/release-action/
    "Release a source code package on PGXN"
  [from pg-hostname]: https://github.com/theory/pg-hostname/pull/2
    "theory/pg-hostname#2 Use new PGXN actions"
  [Use gitattributes]: https://justatheory.com/2025/05/release-on-pgxn/#exclude-files-from-release
    "Auto-Release PostgreSQL Extensions on PGXN: Exclude Files from Release"
  [Keep a Changelog]: https://keepachangelog.com/en/1.1.0/
  [theory/changelog-version-notes-action]: https://github.com/theory/changelog-version-notes-action
    "Extract Changelog Release Notes"
  [softprops/action-gh-release]: https://github.com/softprops/action-gh-release
    "GitHub Action for creating GitHub Releases"
  [PGXS]: https://www.postgresql.org/docs/current/extend-pgxs.html
    "PostgreSQL Docs: Extension Building Infrastructure"
  [platforms.json]: https://github.com/pgxn/matrix-action/blob/main/platforms.json
  [`shared_preload_libraries`]: https://www.postgresql.org/docs/18/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES
  [semver]: https://github.com/theory/pg-semver/actions
  [pg_clickhouse]: https://github.com/ClickHouse/pg_clickhouse/actions
  [chdb]: https://github.com/ClickHouse/pg_chdb/actions
  [hostname]: https://github.com/theory/pg-hostname/actions
  [pgxn/postgres-action workflow]: https://github.com/pgxn/postgres-action/actions/runs/41246481078
  [path issue]: https://github.com/EnterpriseDB/edb-installers/issues/719
    "EnterpriseDB/edb-installers#719 `pg_config --pgxs` outputs long windows path < Postgres 17"
  [Chocolatey]: https://community.chocolatey.org/packages/postgresql
    "Chocolatey PostgreSQL"
  [pgrx]: https://github.com/pgcentralfoundation/pgrx
    "Build Postgres Extensions with Rust!"
  [Peter Eisentraut]: https://peter.eisentraut.org
  [PGDG Apt]: https://wiki.postgresql.org/wiki/Apt "PostgreSQL packages for Debian and Ubuntu"
  [ikalnytskyi/action-setup-postgres]: https://github.com/ikalnytskyi/action-setup-postgres/
    "Setup a PostgreSQL for Linux, macOS and Windows runner machines"
  [Christoph Berg]: https://www.df7cb.de
