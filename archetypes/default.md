---
title: '{{ replace .File.ContentBaseName "-" " " | title }}'
slug: {{ .Name }}
date: {{ now.UTC.Format "2006-01-02T15:04:05Z" }}
description: ~
tags: []
type: {{ .Type }}
# draft: true
# author:
#   name: {{ .Site.Params.PostAuthor.name }}
---

