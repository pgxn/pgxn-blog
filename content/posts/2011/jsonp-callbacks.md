---
title: Using JSONP Callbacks
slug: jsonp-callbacks
date: 2011-05-13T03:58:33Z
aliases: [/post/5441748625/jsonp-callbacks, /post/5441748625]
tags: [JSONP, Callback, Blog, Json]
---

Using the new [JSONP callback support] in the PGXN API, I just added a sidebar
to [my personal blog] that shows a list of all my PGXN distributions. The code
is simple:

    <div class="links">
        <h2>PGXN Code</h2>
        <script type="text/javascript">
        function pgxn_distros(data) {
            document.write('<dl>');
            for (dist in data.releases) {
                document.write(
                    '<dt><a href=https://pgxn.org/dist/' + dist +
                    '>' + dist + '</a></dt>' +
                    '<dd>' + data.releases[dist].abstract + '</dd>'
                );
            }
            document.write('</dl>');
        }
      </script>
      <script type="text/javascript"
         src="https://api.pgxn.org/user/theory.json?callback=pgxn_distros">
      </script>
    </div>

And the output looks like this:

![PGXN Code]

Not bad, eh? AS you can see, JSONP is dead easy to use with the JSON files
served by the JSON API. Try that, CPAN!



  [JSONP callback support]: https://github.com/pgxn/pgxn-api/wiki/JSONP
  [my personal blog]: https://www.justatheory.com/ "Just a Theory"
  [PGXN Code]: ../../media/1205386689_3.png
