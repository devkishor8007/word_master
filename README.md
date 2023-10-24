# word_master
made with go | gorilla/mux
- "WordMaster" is a content management system (CMS) designed specifically for bloggers and content creators.
It offers a user intuitive API that is supported by Golang, Mux, PostgreSQL and GORM ORM. This allows
for handling, publishing and distribution of blog articles.

## run locally
```
$ git clone https://github.com/devkishor8007/word_master.git
$ cd word_master
# Rename the file [ .env.example ] to [ .env ] and if you're using docker; you don't have to modify any variables
  otherwise you need to change the variables according to your needs.
  However, if you're not using Docker, make sure to modify the variables based on your specific requirements.

If you're running the api without docker
$ go run main.go

If you're running the api with docker
$ docker-compose up
 
```

## get started [fresh_project]
```
$ mkdir word_master && cd word_master
$ go mod init your-module-name
$ go get -u github.com/gorilla/mux
$ touch main.go
```

## language, tools & packages
|  language | package / library | database | orm
|----------|----------|----------|----------|
| golang | gorilla / mux | Postgres | Gorm
| | dotenv | |
| | bcrypt | |
| | jwt | |
| | dotenv | |
| | rate | |

## API ENDPOINTS
|  base_url | endpoint | http method | payload | add Authorization token in http header [ Authorization  54545454445 ] | describe |
|----------|----------|----------|----------|--------| --------|
| localhost:3002/api/v1/ | | GET | | no | view the test endpoint
| localhost:3002/api/v1/ | signup | POST | { "username": "Peter", "email": "peter@gmail.com", "password": "1234p" } | no | create a new account
| localhost:3002/api/v1/ | signin | POST | { "email": "peter@gmail.com", "password": "1234p" } | no | login your account
| localhost:3002/api/v1/ | profile | GET | | yes | get your profile
| localhost:3002/api/v1/ | category | GET | | no | view all the category
| localhost:3002/api/v1/ | category | POST | { "name": "golang" } | no | add a new category
| localhost:3002/api/v1/ | articles | GET | | no | view all the articles
| localhost:3002/api/v1/ | articles/contributors | GET | | yes | view all the contributors articles  
| localhost:3002/api/v1/ | articles | POST | { "title": "What is GO?", "content": "Go is a statically typed...", "publication_date": "2023-10-14T10:05:00Z", "category_id": 1 } | yes | create a new article by contributors
| localhost:3002/api/v1/ | comment | GET | | no | view all the comments
| localhost:3002/api/v1/ | comment | POST | { "text": "What is NPM?", "comment_date": "2023-10-14T10:05:00Z", "author_id": 1, "article_id": 1 } | yes | add a new comment in a article


