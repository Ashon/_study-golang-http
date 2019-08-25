# Simple Web Framework in Go

주말동안 간단한 웹 프레임워크를 만들어 보면서, Go의 기능들을 공부해 보았다.

2일간 빠르게 내용들을 찾아보고 기본적인 Web API를 구현하기 위한 프레임워크를 만들었다.

설계가 조악한 부분들이 보이긴 하지만, 비즈니스 로직 영역과 http 리퀘스트 처리에 대한 영역을
최대한 분리해 보도록 구현해 보면서, 높은 수준의 Go 애플리케이션을 작성하기 위한 방법들을 찾아보는데
중점을 둔다.

## Web Framework

나는 파이썬을 주로 쓰는데, 이번 공부에서는 Golang으로 파이썬에서의
개발 경험들을 Golang으로 만들어 내기 위한 방법들을 찾아보고 구현하였다.

웹 애플리케이션의 코어 로직들은 최대한 숨기면서, 사용자 기능들만 빠르게 구현하기 위한 방법들을
위해, `core`, `app` 디렉토리로 구분하여 간단한 프레임워크와 애플리케이션 구현을 작성해 보았다.

``` txt
$ tree
.
├── app ................ 사용자 애플리케이션의 구현
│   ├── routes.go ...... url route를 관리
│   └── views.go ....... view 함수들
│
├── core ............... net/http에서 API 구현을 위한 프레임워크 로직들
│   ├── config.go ...... 프레임워크로 관리될 설정 모음
│   ├── exceptions.go .. 에러 인터페이스
│   ├── handler.go ..... request / response handler
│   ├── logger.go ...... logger 설정
│   ├── request.go ..... view 구현에 사용되는 request 스트럭처
│   ├── response.go .... response structure
│   └── server.go ...... 서버를 빌드하고 애플리케이션 동작에 대한 코드
│
├── docker-compose.yml
├── Dockerfile
│
├── go.mod
├── go.sum
│
├── main.go ............. 메인 파일
│
└── README.md
```

### Core

웹 개발 프레임워크들의 코어 로직으로서, 비즈니스 로직을 작성하기 위한 간단한 인터페이스를 제공하고
애플리케이션 라이프사이클을 관리하게 된다.

### App

Core에 구현된 인터페이스를 이용해 비즈니스로직을 이해하기 쉽게 만들어 보고,
Go의 특수한 예외처리 방식을 Core로 커버 가능한지를 검증해 본다.

## 무엇을 배웠나

### net/http

### go mod

### panic

### interface

## 더 알아보고픈 내용

### Dependency injection

### Types in go
