# go-calendar-practice

## Quick Start

## FileTree

```
├─README.md
├─
├─go.mod
├─go.sum
├─cmd
├─config            써드파티 라이브러리의 설정 정보등을 담아두는 곳
└─pkg
    ├─db            데이터베이스 설정
    ├─loaders       서버가 실행될 때 load되어야할 모듈들
    ├─middlewares   Auth등 Request가 들어와 Route Handler(Controller) 전에 호출되어야하는 로직
    └─modules       각각의 도메인들에 대한 파일이 모여있는 곳
        ├─base
        └─example
          ├─exmaple.controller.go     해당 도메인 관련 라우터가 정의되는 곳
          ├─exmaple.service.go        해당 도메인의 비즈니스 로직의 흐름이 정의되는 곳
          └─exmaple.model.go          해당 도메인 관련 DB작업이 일어나는 곳

```
