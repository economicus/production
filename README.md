# Economicus
이코노미쿠스 백엔드 서버 API입니다.\
환경변수에 대한 내용은 공개하지 않으니 따로 연락주시길 바랍니다.
> .env 파일 설정 없이 서버 구동 불가합니다.\
> 자체적인 .env 파일을 만들거나 다운 받은 후에 프로젝트 루트 폴더에 넣어 사용하십시오.

## License
작성 예정

## Overview
작성 예정

## Requirements
도커로 사용 시,

- [docker](https://www.docker.com/products/docker-desktop)
- [docker-compose](https://docs.docker.com/compose/)

도커를 사용하지 않을 시,

- [mysql](https://www.mysql.com/)
- [golang](https://go.dev/)

## Installation
도커 사용 시,

```shell
$ git clone https://github.com/economicus/economicus-be.git
$ cd economicus-be
# set up .env file in root folder
$ docker-compose up -d --build
```

도커 사용하지 않을 시,

```shell
$ git clone https://github.com/economicus/economicus-be.git
$ cd ./economicus-be/app
# set up environments
$ make build_and_run
```

## Go Packages
- Routing: [gin](https://github.com/gin-gonic/gin.git)
- ORM: [gorm](https://gorm.io/)
- Logging: [logrus](https://github.com/sirupsen/logrus)
- JWT: [jwt-go](https://github.com/dgrijalva/jwt-go.git)


## Usage
깃 위키에 담을 예정

## Layout

```text
economicus
│
├── README.md
├── .gitignore
├── .env
├── docker-compose.yml
│
├── volumes
│   │-- mysql
│   └── nginx
│
├── quant
│   └── 
│
└── main
    ├── go.mod
    ├── go.sum
    ├── Makefile
    ├── Dockerfile
    ├── cmd
    │   └── economicus
    │       └── main.go
    │
    ├── commons
    │   ├── converter
    │   │   ├── converter_test.go
    │   │   ├── interface-converter.go
    │   │   └── string-converter.go
    │   ├── bcrypt.go
    │   └── bcrypt_test.go
    │
    ├── config
    │   ├── app.go
    │   ├── aws.go
    │   ├── database.go
    │   ├── jwt.go
    │   └── util.go
    │
    │
    ├── internal
    │   ├── api
    │   │   ├── handler
    │   │   │   ├── auth.go
    │   │   │   ├── comment.go
    │   │   │   ├── helper.go
    │   │   │   ├── quant.go
    │   │   │   ├── reply.go
    │   │   │   └── user.go
    │   │   ├── hateos
    │   │   │   └── hateos.go
    │   │   ├── repository
    │   │   │   ├── auth.go
    │   │   │   ├── comment.go
    │   │   │   ├── comment_test.go
    │   │   │   ├── quant.go
    │   │   │   ├── quant_test.go
    │   │   │   ├── reply.go
    │   │   │   ├── reply_test.go
    │   │   │   ├── repository.go
    │   │   │   ├── repository_test.go
    │   │   │   ├── user.go
    │   │   │   └── user_test.go
    │   │   ├── routes
    │   │   │   ├── auth.go
    │   │   │   ├── comment.go
    │   │   │   ├── quant.go
    │   │   │   ├── reply.go
    │   │   │   └── user.go
    │   │   ├── service
    │   │   │   ├── auth.go
    │   │   │   ├── comment.go
    │   │   │   ├── quant.go
    │   │   │   ├── reply.go
    │   │   │   └── user.go
    │   │   ├── token
    │   │   │   └── token.go
    │   │   ├── api.go
    │   │   ├── logger.go
    │   │   └── middleware.go
    │   ├── drivers
    │   │   ├── aws.go
    │   │   └── database.go
    │   └── models
    │       ├── comment.go
    │       ├── error.go
    │       ├── models.go
    │       ├── object.go
    │       ├── profile.go
    │       ├── quant.go
    │       ├── quant-option.go
    │       ├── quant-result.go
    │       ├── query-option.go
    │       ├── reply.go
    │       └── user.go
    │   
    ├── script
    │   ├── run.sh
    │   └──
    │
    └── tests
        ├──
        └──
```