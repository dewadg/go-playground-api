FROM golang:1.18-alpine AS server-build

RUN apk add --no-cache git

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o go-playground

FROM node:lts-alpine AS web-build

WORKDIR /build

COPY . .

RUN cd web && npm i && npm run build

FROM golang:1.18-alpine

WORKDIR /usr/bin/go-playground

COPY --from=server-build /build/go-playground .
COPY --from=web-build /build/web/dist web/dist

RUN chmod a+x go-playground
RUN mkdir /opt/go-playground

EXPOSE 8000

ENV APP_ENV production
ENV TEMP_DIR /opt/go-playground
ENV WEB_DIR /usr/bin/go-playground/web/dist

ENTRYPOINT ["/usr/bin/go-playground/go-playground"]
