FROM golang:1.18 AS server-build

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -ldflags '-extldflags "-static"' -o go-playground

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
RUN mkdir /opt/go-playground/temp
RUN mkdir /opt/go-playground/data

EXPOSE 8000

ENV APP_ENV production
ENV TEMP_DIR /opt/go-playground/temp
ENV WEB_DIR /usr/bin/go-playground/web/dist

ENTRYPOINT ["/usr/bin/go-playground/go-playground"]
