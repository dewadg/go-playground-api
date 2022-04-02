# Go Playground API

Backend part of [Go Playground Web](https://github.com/dewadg/go-playground-web)

Basically just another Go playground that I build for learning purpose.

## How to run

Create new file `.env` by copying from `.env.example`. Set required variables in there.

```
go get ./...

set -a && . ./.env

make serve
```

## GraphQL

GQL endpoint is at `/graphql`
