API Architecture
================

This is a small document describing the architecture of the REST API.
It's not final, and things might be different in the finished project,
but it's a rough draft of how the API should behave.

Authentication/Authorisation
----------------------------

Nachricht will only support OpenID authentication. It will not handle
user registration or authentication itself. To start off, we'll only
support auth for content creators/admins, but in the future we'll also
support normal users to post comments and the scopes should reflect that.

As one of the reasons for making this blog is to play with OpenID, the
names of the scopes aren't set yet but let's define them as `admin`, `writer`,
`user` and `all` for now. The definition of these for this document are:

* `admin`: The admin of all `nachricht` sites, can essentially do everything
on every domain. If not explicitely mentioned, has the same access as `writer`.
* `writer`: The "owner" of a blog, can write stories, moderate, and almost
everything else for a single domain.
* `user`: A logged in user, will be able to comment.
* `all`: An unauthenticated user, will be able to read articles.

Domain support
--------------

I've considered going going for a `/v1/nachricht.example.org/articles/` style
structure, but I found it a bit ugly and wanted to it a bit different.
For the initial design, I'm choosing to let the `Host` header decide what
site is being queried. So, a query to `/v1/articles/` with
`Host: nachricht.example.org` set, will show the articles for `nachricht.example.org`

API versioning
--------------

I want to keep this simple, as the backend is purely an API, the first path
element will be the version. For example: `/v1`.

API
---

Domains
-------

`domain` object attributes:
* `name` (string)(readonly): Name of the domain, this is also the domain
that will be matched on.
* `enabled` (bool): Is the domain enabled.

----------------------

`GET /v1/domains/`

Allowed users:
* `admin`

Returns a list of `domain` objects.

----------------------

`GET /v1/domains/:domainName/`

Allowed users:
* `admin`

Returns a domain if `:domainName` exists, 404 otherwise.

----------------------

`PUT /v1/domain/:domainName/`

Allowed users:
* `admin`

Add a domain with name `:domainName`. Expects a `domain` object in the body.

----------------------

Articles
--------

`article` object attributes:
* `title` (string): The title of the article.
* `content` (string): The content of the article.
* `tags` (list[string]): Tags on this article.
* `publish` (datetime): Time and date when this article should be published.
* `slug` (string)(readonly): Slug for the article.
* `advertised` (bool)(readonly): If this article has been posted to social media.

---------------------

`GET /v1/articles/`

Allowed users:
* `writer`: Can list unpublished articles.
* `all`: Can list published articles.

URL parameters:
* `tag` (string): Filters the articles, returning only the ones tagged with `tag`.

Gets a list of all `article` objects.

----------------------

`GET /v1/articles/:articleSlug/`

Allowed users:
* `writer`: Can get unpublished articles.
* `all`: Can get published articles.

Gets a single `article` object.

----------------------

`PUT /v1/articles/:articleSlug/`

Allowed users:
* `writer`

Adds an article with slug `:articleSlug`, expects an `article` object in the body.

----------------------

Tags
----

`tag` object attributes:
* `name` (string): name of the tag.
* `article_count` (int): number of articles with the tag.

-----------------------

`GET /v1/tags/`

Allowed users:
* `all`

Gets all `tag` objects.

-----------------------