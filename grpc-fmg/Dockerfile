FROM alpine:3.6

RUN sed -i 's/dl-cdn\.alpinelinux\.org/mirrors\.aliyun\.com/g' /etc/apk/repositories

RUN apk update --no-cache

RUN mkdir /main
COPY main /main/

WORKDIR /main
ENTRYPOINT ["./main"]
