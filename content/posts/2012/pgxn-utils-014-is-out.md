---
title: pgxn-utils 0.1.4 is out!
slug: pgxn-utils-014-is-out
date: 2012-04-15T21:44:12Z
aliases: [/post/21171605888/pgxn-utils-014-is-out, /post/21171605888]
tags: [PGXN, Utils, Automation, Release]
---

*by Dickson S. Guedes*

Hi everybody!

I'm proud to tell you that a new version of [PGXN Utils] was released with
this new features:

- Git support
- Templates and custom templates
- [PGXN Client] integration

## Git support

You can start a new extension with or without version control. By default
`pgxn-utils` supports [git] but it will not create a repository unless you use
`--git` option in the skeleton task.

You can try:

```console
$ pgxn-utils skeleton my_cool_versioned_extension --git
```

When you create a new extension with git support in addition to creating the
skeleton, `pgxn-utils` will initialize a git repository and create the initial
commit.

Once you have your extension in a git repository your `bundle` will use only
the committed files to create the archive, but if your repository is dirty
then `pgxn-utils` will suggest that you to commit or stash your changes before
bundling.

You must be careful with new files not added to repository, because they will
**not** be archived.

## Default templates

There are three default templates: `sql`, `c` and `fdw`. If you call
`skeleton` without specifying a template, `sql` is the default. But if your
extension will supply some C modules or you will create a FDW, you can create
the extension calling `skeleton` with a `--template` option.

Try:

```console
$ pgxn-utils skeleton my_cool_c_extension --template=c
$ pgxn-utils skeleton my_cool_fdw_extension --template=fdw
```

The templates contain example code and some links to PostgreSQL documentation
that will try to help you to start coding. SQL and C templates contains some
test examples, and the example code will compile and pass `make installcheck`.
However, this code is intended to be an example, and you must write your own
tests and code.

## Custom templates

If you don't like the templates provided by `pgxn-utils` you can create you
own. Just create a directory with at least a `META.json` or `META.json.tt`
file and then use your directory as argument to the `--template` option.

To know how create your own template, see the examples in the [templates
directory].

## PGXN Client integration

If you have [PGXN client] installed you can change the command line from
`pgxn-utils some_task` to `pgxn some_task` and this will save you some typing.

See:

```console
$ cd /tmp
$ pgxn skeleton --help
PGXN Utils version: 0.1.4
Usage:
  pgxn skeleton extension_name

Options:
      [--git]                            # Initialize a git repository after create the extension
  -a, [--abstract=ABSTRACT]              # Defines a short description to abstract
  -p, [--target=TARGET]                  # Define the target directory
                                          # Default: .
      [--template=TEMPLATE]              # The template that will be used to create the extension. Expected values are: sql, c, fdw
                                          # Default: sql
  -r, [--release-status=RELEASE_STATUS]  # Initial extension's release status
  -d, [--description=DESCRIPTION]        # A long text that contains more information about extension
  -b, [--generated-by=GENERATED_BY]      # Name of extension's generator
  -l, [--license=LICENSE]                # The extension license
  -t, [--tags=one two three]             # Defines extension's tags
  -v, [--version=VERSION]                # Initial version
  -m, [--maintainer=MAINTAINER]          # Maintainer's name <maintainer@email>

Creates an extension skeleton in current directory

$ pgxn skeleton test
      create  test
      create  test/test.control
      create  test/.gitignore
      create  test/.template
      create  test/META.json
      create  test/Makefile
      create  test/README.md
      create  test/doc/test.md
      create  test/sql/test.sql
      create  test/sql/uninstall_test.sql
      create  test/test/expected/base.out
      create  test/test/sql/base.sql

$ cd test/
$ pgxn bundle
          run  make distclean from "."
      create  /tmp/test-0.0.1.zip
```

I hope you enjoy this version. ":)

  [PGXN Utils]: https://github.com/guedes/pgxn-utils
  [PGXN Client]: https://pgxnclient.projects.postgresql.org/
  [git]: https://git-scm.org
  [templates directory]: https://github.com/guedes/pgxn-utils/tree/master/lib/pgxn_utils/templates
