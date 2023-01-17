FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

RUN chmod +x ./tools/wait-for-it.sh ./tools/docker-entrypoint.sh
EXPOSE 8080

ENTRYPOINT ["./tools/docker-entrypoint.sh"]
CMD ["/app/main"]
