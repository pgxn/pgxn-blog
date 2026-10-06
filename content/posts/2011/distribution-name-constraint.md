---
title: Distribution Name Constraint
slug: distribution-name-constraint
date: 2011-05-04T20:42:37Z
aliases: [/post/5197276864/distribution-name-constraint, /post/5197276864]
tags: [Distribution Name, Constraint, Domain, ASCII, File Name]
---

Following [a pgxn-users discussion], I'm looking at adding a constraint on the
names distributions can have. I'm thinking this:

    CEATE DOMAIN namespace AS CITEXT
    CHECK ( VALUE ~ '^[-a-z0-9_]{2,}$' );

So, just ASCII numbers and letters, dash, and underscore, with a minimum
length of two characters. Maybe `,` and `+` should be allowed, too? What do
you think? What constraints do you expect on distribution download file names?



  [a pgxn-users discussion]: https://groups.google.com/group/pgxn-users/browse_thread/thread/4528b2a02e8886ff
