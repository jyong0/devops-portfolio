FROM postgres:16-alpine

# 필요한 패키지 설치
RUN apk add --no-cache curl bash

# goose 설치
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

# goose 와 psql 모두 포함됨
ENTRYPOINT [ "bash", "-c" ]
