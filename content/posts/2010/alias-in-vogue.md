---
title: ALIAS Back in Vogue 
slug: alias-in-vogue
date: 2010-09-02T13:00:00Z
aliases: [/post/1053165383/alias-in-vogue, /post/1053165383]
tags: [Alias, PostgreSQL 9, PL/pgSQL, PostgreSQL, Function, Parameter, Named Parameter]
---

I've been hard at work on [PGXN Manager], the app for users to upload
distributions to PGXN. I am of course following my own dictum: "the database
*is* the model." As a result, I've been creating an API for creating,
updating, and deleting entities, as well as generating JSON (more on that
later).

I'm also using PostgreSQL 9. The motivation to make the jump to 9.0 was to to
try to use the [JSON data type patch], but I abandoned it when I couldn't get
it to compile. (I might come back to it later, but right now I'm trying to
practice [YAGNI] and [JFDI] so that I don't end up owning a yak farm). But
there are other reasons to stick with 9.0, like improved [hstore] support, the
[DO] statement, and [named parameters].

Actually, this last one is very nice, as it allows me to use SQL syntax to
specify function parameters instead of using an `hstore` value to hack it. For
example, I have this function for update a user record:

    CREATE OR REPLACE FUNCTION update_user(
        nick  LABEL,
        name  TEXT  DEFAULT NULL,
        email EMAIL DEFAULT NULL,
        uri   URI   DEFAULT NULL
    ) RETURNS BOOLEAN LANGUAGE plpgsql SECURITY DEFINER AS $$
    DECLARE
        _email ALIAS FOR email;
        _uri   ALIAS FOR uri;
    BEGIN
        UPDATE users
           SET full_name  = COALESCE(name,   full_name),
               email      = COALESCE(_email, users.email),
               uri        = COALESCE(_uri,   users.uri),
               updated_at = NOW()
         WHERE nickname   = nick
           AND status     = 'active';
        RETURN FOUND;
    END;
    $$;

The nice thing about named parameters is that I can call this function like
so:

    SELECT update_user(
        nick  := 'theory',
        name  := 'David E. Wheeler',
        email := 'justatheory@pgxn.org',
        uri   := 'https://www.justatheory.com/'
    );

Hell, since the parameters are named, I can specify them in any order. And
because there are defaults, I can omit one or more of them:

    SELECT update_user(
        name  := 'David E. Wheeler',
        email := 'justatheory@pgxn.org',
        nick  := 'theory'
    );

And it will just work. Nice!

One thing I did discover, though: Now that I'm using named parameters, I don't
want to use [Hungarian notation] for the parameter names. Back before 9.0,
when callers never saw argument names when calling a function, I could call
them whatever I wanted. This was especially important when executing queries
in PL/pgSQL, because one needs to be aware of parameter (and variable!) names
that conflict with SQL identifiers, especially columns. (Another nice feature
in 9.0 is that it will throw an exception if you create a PL/pgSQL function
with conflicting parameter and identifier names.) But now that the parameter
names are more likely to be exposed and *used* to the function caller, I want
them to be meaningful.

I messed with this for a while. I was okay with using "nick" for "nickname"
and "pass" for "password", but was annoyed with the alternatives for "email"
and "uri," especially since they're used with those names elsewhere (queries
against the table, generated JSON). I really wanted to use parameter names
that were the same as the columns. I tried just using the dollar variables
(`$1`, `$2`, etc.), but got errors for them, too: they seem to be compiled
into the parameter names.

Then, on a guess, I tried creating aliases for the variables. From the example
above, its:

        _email ALIAS FOR email;
        _uri   ALIAS FOR uri;

Then I used `_email` and `_uri` in my `UPDATE` statement. And what do you
know, it worked! I found this somewhat humorous, given the history of `ALIAS`.

It used to be, long ago, that you couldn't use the parameter names in the body
of PL/pgSQL functions, so `ALIAS` was there to let you alias the dollar
variable names to other names. But somewhere around, oh, 8.0 or so, we were
blessed with the ability to use the parameter names directly. Suddenly `ALIAS`
seemed superfluous. I've hardly ever used it myself. It's just been sitting
there, like my appendix, waiting for another use.

And now there is one. A really good one! I can safely use parameter names that
are the same as column names in my PL/pgSQL functions as long as I alias them.
And it just works!

Well, almost. It seems that `ALIAS` means what it says: the parameter names
are still around can can be used. So sometimes you might run into an error
like

    ERROR:  column reference "email" is ambiguous

Even though you're not using the variable. I ran into this in the update
function where I was using the column names in the left-hand side of the `SET`
expressions. The solution, fortunately, is simple: table-qualify the column
names as appropriate:

    UPDATE users
       SET full_name  = COALESCE(name,   full_name),
           email      = COALESCE(_email, users.email),
           uri        = COALESCE(_uri,   users.uri),

Note the use of `users.email` instead of just `email` in the `COALESCE()`
function. Seems like a reasonable workaround in exchange for the ability to
have parameter names that match column names. I'm sold!

Now just to consider whether to change the `nick` and `pass` parameter names
to `nickname` and `password` in order to be completely consistent. I guess
it's a good idea.

More next week. I've been doing lots of hacking and have much to share, but
have another project that will take up my time between now and Monday, so I'll
have to come back to it.



**Update 2010-08-07:** I turns out that there is a much better way to do this:
Just function-name-qualify parameter names where they might conflict with
database object names:

    UPDATE users
       SET full_name  = COALESCE(name,   full_name),
           email      = COALESCE(update_users.email, users.email),
           uri        = COALESCE(update_users.uri,   users.uri),

No need for the aliases at all! I had no idea about this feature. Many thanks
to Colin 't Hart for the comment below about how this is available in Oracle
and to Tom Lane for [smacking me upside the head] with [the fine manual] (look
for the "note" at the bottom) when I asked about it.



  [PGXN Manager]: https://github.com/theory/pgxn-manager
  [JSON data type patch]: https://commitfest.postgresql.org/action/patch_view?id=351
  [YAGNI]: https://en.wikipedia.org/wiki/YAGNI
  [JFDI]: https://www.urbandictionary.com/define.php?term=JFDI
  [hstore]: https://developer.postgresql.org/pgdocs/postgres/hstore.html
  [DO]: https://developer.postgresql.org/pgdocs/postgres/sql-do.html
  [named parameters]: https://developer.postgresql.org/pgdocs/postgres/sql-syntax-calling-funcs.html#SQL-SYNTAX-CALLING-FUNCS-NAMED
  [Hungarian notation]: https://en.wikipedia.org/wiki/Hungarian_notation
  [smacking me upside the head]: https://archives.postgresql.org/pgsql-hackers/2010-09/msg00404.php
  [the fine manual]: https://www.postgresql.org/docs/9.0/static/plpgsql-structure.html
