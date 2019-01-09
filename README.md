nachricht - Backend
===================

What is nachricht?
------------------

`nachricht` is not much right now, but it's going to be a multi-domain blogging engine.

Another blogging engine? Why?
-----------------------------
Mostly because I want to make one. I used to have my own blogging engine, but the web has changed since I made that. I'm making this to play with some tech I haven't used before. Also, because no blogging engine does exactly what I want without heavy customisation, I might as well spend that time doing something that I consider fun.

What features will nachricht have?
----------------------------------
The features I want to include on the first working release:
* Writing and publishing of articles
* Microservice architecture
* Backend that supports multiple domains
* OpenID authentication/authorisation
* Posting to twitter
* Kubernetes deployment
* REST API
* Tags

Features I want to add to add later:
* A comment system
* Posting to other social media
* GRPC API
* Anything that comes up in me

What technology do you intend to use?
-------------------------------------
The backend will be in go as that's the language I've been using the most lately. The frontend will be done in TypeScript and either vue or react. The persistent store will be PostgreSQL.