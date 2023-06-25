FROM golang:1.19.10-alpine3.18

WORKDIR /app

RUN apk update && apk add --no-cache make nodejs npm sqlite gcc musl-dev

COPY . .

EXPOSE 3005

CMD ["make","run"]