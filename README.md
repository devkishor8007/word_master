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

## API DOCS - Swagger 
After running the 'word_master' api, you can get the url: `http://localhost:3002/swagger/index.html` and hit on any browser to view the api docs.
If you want to change or add any thing on the swagger docs, you can and after modify on the code; do `$ swag init` and re-run the api.

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
| | swagger | |
