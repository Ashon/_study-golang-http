# Simple Web Framework in Go

주말동안 간단한 웹 프레임워크를 만들어 보면서, Go의 기능들을 공부해 보았다.

설계가 일부 조악한 부분들이 보이긴 하지만, 비즈니스 로직 영역과 http 리퀘스트 처리에 대한 영역을
최대한 분리해 보도록 구현해 보면서, 높은 수준의 Go 애플리케이션을 작성하기 위한 방법들을 찾아보는데
중점을 둔다.

## ToC

- [Simple Web Framework in Go](#simple-web-framework-in-go)
  - [ToC](#toc)
  - [Web Framework Overview](#web-framework-overview)
    - [Core](#core)
    - [App](#app)
  - [Experiences in python](#experiences-in-python)
    - [Simple DI in python](#simple-di-in-python)
      - [Core module](#core-module)
      - [User Application](#user-application)
    - [Managing Views](#managing-views)
  - [Implementation in Go](#implementation-in-go)
    - [Features](#features)
      - [View, Route Management](#view-route-management)
      - [Middleware Management](#middleware-management)
      - [Request Flow Control](#request-flow-control)
  - [소감](#%ec%86%8c%ea%b0%90)
  - [무엇을 배웠나](#%eb%ac%b4%ec%97%87%ec%9d%84-%eb%b0%b0%ec%9b%a0%eb%82%98)
    - [net/http](#nethttp)
    - [go mod](#go-mod)
    - [panic](#panic)
    - [interface](#interface)
    - [package system](#package-system)
  - [더 알아보고픈 내용](#%eb%8d%94-%ec%95%8c%ec%95%84%eb%b3%b4%ea%b3%a0%ed%94%88-%eb%82%b4%ec%9a%a9)
    - [Dependency injection](#dependency-injection)
    - [Types in go](#types-in-go)
    - [Other Go Web Framework](#other-go-web-framework)

## Web Framework Overview

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

애플리케이션으로 들어오는 요청과 응답을 다루기 위한 방법들을 관리할 수 있게 한다.

### App

Core에 구현된 인터페이스를 이용해 비즈니스로직을 이해하기 쉽게 만들어 보고,
Go의 특수한 예외처리 방식을 Core로 커버 가능한지를 검증해 본다.

## Experiences in python

나는 주로 파이썬의 `Flask`, `Django`, `Sanic`, `Vibora` 등의 웹 프레임워크들을
이용해서 간단한 프로젝트들을 만들었다.

`Django`의 경우는 애플리케이션을 구현하기 위한 많은 기능들이 프레임워크에서 제공되기 때문에 프레임워크의 사상을 따라 애플리케이션을 만들 수 있다.

그외 `Flask` 류의 프레임워크들은 뷰를 구현하고 라우트를 관리하기 위한 다양한 기능들을 제공해 주긴 하지만, Django같은 DI(Dependency Injection) Framework가 내장되어 있지 않아, 큰 규모의 애플리케이션을 만들기 위해서는 DI를 위한 코드들을 별도로 작성해야 할 수 있다.

### Simple DI in python

파이썬은 모듈들을 간단한 방법으로 Dynamic import할 수 있게 만들어져 있기 때문에, DI를 위한 기능도 손쉽게 구현할 수 있다.

자바 류의 언어에서 구현된 `DI` 라이브러리들의 사용 경험을 그대로 따라가기 보다는 간단한 방식으로 application에 내가 설정을 통해 원하는 미들웨어나 의존성들을 설정으로 주입할 수 있도록 구현해서 쓰고 있다.

아래는 파이썬에서 내가 주로 사용했던 DI를 위한 dynamic import 로직과 injection의 예시이다.

#### Core module

``` python
# file: utils.py
import importlib


def get_module(module_path: str):
    """module_path를 인자로 받아 해당 모듈을 리턴한다.
    """

    module_path, _, child_name = module_path.rpartition('.')

    module = importlib.import_module(module_path)
    child = getattr(module, child_name)

    return module, child


def instantiate(classpath, constructor):
    """module_path로부터 획득한 클래스를 인스턴스로 생성해 인스턴스를 리턴한다.
    """

    _, instance_cls = get_module(classpath)
    instance = instance_cls(**constructor)

    return intsance
```

``` python
# file: core.py
from flask import Flask
from utils import instantiate


def build_application(settings):
    app = Flask(__name__)
    app.config.from_object(settings)

    for name, spec in settings.MIDDLEWARES.items():
        middleware = instantiate(**spec)

        # 적절한 방법으로 middleware와 app 컨텍스트를 연결.
        middleware.init_app(flask_app)

        setattr(app, name, middleware)

    return app
```

#### User Application

정의된 코어 소스코드들로 애플리케이션 라이프사이클을 제어하도록 하고, 사용자는
필요한 모듈만을 정의하여 애플리케이션을 작성할 수 있게 된다.

``` python
# file: settings.py
MIDDLEWARES = {
    'user_api_client': {
        'classpath': 'app.middlewares.user_api.Client',
        'constructor': {
            'api_url': 'http://blabla.api.com:8080',
            'api_user': 'testuser',
            'api_key': 'blabla'
        }
    }
}
```

``` python
# file: app.py
from core import build_application
import settings


flask_app = build_application(settings)
flask_app.run(...)
```

이런식으로 관리하게 되면, Django 에서의 DI를 이용해 미들웨어를 관리하는 것 처럼
application의 네임스페이스에 설정한 미들웨어를 주입하고 런타임에서 사용할 수 있게 해 준다.

미들웨어 로직과 애플리케이션 코어와의 루즈한 커플링을 제공해 줌으로써,
넓은 범위로의 재사용성을 보장해 줄 수 있다.

### Managing Views

Flask를 예로 들면 간단한 API View는 데코레이터를 이용해 정의할 수 있다.

``` python
# file: app.py
from flask import Flask


app = Flask(__name__)


@app.route()
def hello():
    return 'hello world'
```

실제 비즈니스 로직을 처리하는 거대한 애플리케이션을 만들기 위해서는 애플리케이션 컨텍스트와
강하게 연결되는 데코레이터 형식 보다는, 구현된 비즈니스 로직을 가지고 사용할 주체에서 빌드업 해서 쓰는게
옳다고 생각한다.

``` python
# file: views.py
def hello():
    return 'hello world'
```

``` python
# file: app.py
from flask import Flask
from views import hello


app = Flask(__name__)
app.route('/', methods=['GET'])(hello)
```

이런식으로 작성하면 비즈니스 로직을 애플리케이션 로직으로부터 분리해서 작게 관리할 수 있다.

위에서 설명한 DI와 조합하면, 애플리케이션 설정으로 사용자가 필요한 API들을 하나의 `suite`로서
관리할 수 있게 되는 장점이 있다.

## Implementation in Go

글에 적은 파이썬 경험들을 Go로 표현해 보면서, Go의 사용성이나 Go의 철학들을 이해해 보면 좋을 것 같다.

유저 애플리케이션을 작성하기 위해 API View들을 비교적 쉽게 만들 수 있게 해 보고, API Route 관리,
프레임워크의 리퀘스트 처리 흐름을 사용자가 일부 제어 가능하도록 하는 미들웨어 등을 작성할 수 있는
기능을 구현 해 보고자 했다.

### Features

#### View, Route Management

#### Middleware Management

#### Request Flow Control

## 소감

- `tab`은 영 적응이 안되지만 `vscode`의 코딩 어시스턴트 기능들에 의존했기 때문에,
  별다른 불편함 없이 코드를 작성할 수 있었다.
  - 타입 시스템으로 코드를 작성하면서 바로 코드 스펙들을 알 수 있는게 장점이기도 하다고 생각하지만..
    너무 IDE 의존적이 되어, 이런 부분들에 무뎌질 수 있겠다는 생각이 들었다.
- 파이썬에 비해 일부 비즈니스 로직들의 표현이, verbose 하다는 느낌은 여전히 지울 수 없다.
  - 성능을 위해선 어쩔 수 없는 부분인 걸까..
- 프레임워크에서 할 일들이 표현되는 방식을 보면, 결국 `Go`, `Python` 할 것 없이 비슷하게 표현됨을
  알 수 있었다.

## 무엇을 배웠나

### net/http

### go mod

### panic

### interface

### package system

## 더 알아보고픈 내용

### Dependency injection

### Types in go

### Other Go Web Framework
