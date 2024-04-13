# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

#COPY * ./app
COPY ./ /app

RUN go get github.com/gin-gonic/gin@v1.8.1

ENV MYSQL_USER=root
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes
ENV MYSQL_PASSWORD=''
ENV MYSQL_USER=root

RUN go build -o main

EXPOSE 3306
EXPOSE 2040

CMD ["PWD"]
RUN ./main prod_config
CMD [ "/app/main" ]