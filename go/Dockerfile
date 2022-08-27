# base stage
FROM golang:1.19-bullseye AS base
RUN apt update && apt install -y \
    build-essential \
    curl \
    groff \
    less \
    python3-pip \
    nodejs \
    npm \
    zip \
    && apt clean \
    && rm -rf /var/lib/apt/lists/*
RUN npm install -g n serverless serverless-localstack \
    && n stable
# ↑のRUNと同じ中で最新版にしようとするとエラーになる
RUN npm install -g npm@latest
ENV TZ="Asia/Tokyo"

WORKDIR /workspace

# work
FROM base as work

RUN useradd -m -d /home/dev -s /bin/bash dev

ENV PATH $PATH:/home/dev/.local/bin

USER dev

RUN pip3 install awscli awscli-local \
    && pip3 cache purge

RUN go install github.com/766b/go-outliner@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install golang.org/x/tools/gopls@latest \
    && go install github.com/dmarkham/enumer@latest \
    && go clean -modcache -cache

RUN echo 'alias aws=awslocal' >> ~/.bashrc

ENV AWS_DEFAULT_REGION="ap-northeast-1"

# build stage
FROM base AS builder
COPY . .
RUN go build -o /go/bin/app -v ./...

# final stage
FROM debian:bullseye-slim
RUN apk --no-cache add ca-certificates
ENV TZ="Asia/Tokyo"
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=hanseiboard Version=0.0.1
EXPOSE 3000
