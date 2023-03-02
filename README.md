# RaulGoBackEndProjeto
GO + Fiber + GORM + Docker + Swagger +Migrations + Air + Jenkins 

Fiber is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.

Docker is an open source platform that enables developers to build, deploy, run, update and manage containers—standardized, executable components that combine application source code with the operating system (OS) libraries and dependencies required to run that code in any environment.

GORM provides a migrator interface, which contains unified API interfaces for each database that could be used to build your database-independent migrations.

Projeto em GO utilizando containerização com docker+docker-compose . CRUD realizado com GORM e Postgres. Swagger para GO utilzado para documentar as rotas da API.

Hot reloading with a package called air to rebuild the app when it has changes. The command runs an installation through the DockerFile and ,then, it builds the container through docker-compose.yml

Install docker 

docker compose up

Swagger for mapping all the API endpoints

Jenkins for integration and deploy
