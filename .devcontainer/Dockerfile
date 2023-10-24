ARG GO_VERISION="1.19-bullseye"

# goのツールはビルドは重いけどワンバイナリで1箇所にまとまるので別のステージでビルドしてまとめてコピーする
FROM golang:${GO_VERISION} as go-tool-builder
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go install github.com/766b/go-outliner@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install golang.org/x/tools/gopls@latest \
    && go install github.com/dmarkham/enumer@latest

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${GO_VERISION}

# [Choice] Node.js version: none, lts/*, 18, 16, 14
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
RUN apt update -y \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt autoremove -y \
    && apt update -y \
    && apt -y install --no-install-recommends \
    bash-completion \
    ca-certificates \
    curl \
    git \
    gnupg \
    groff \
    less \
    lsb-release \
    python3 \
    python3-pip \
    zip

# awscli
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm aws awscliv2.zip -rf \
    && pip3 install awscli-local

# CDK
RUN npm i -g aws-cdk aws-cdk-local

# Docker
RUN mkdir -p /etc/apt/keyrings \
    && curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null \
    && apt update -y && apt autoremove -y \
    && apt install -y docker-ce-cli docker-compose-plugin

# docker-composeは古いので指定したバージョンに更新
ARG DCVERSION="2.10.2"
ENV DCURL="https://github.com/docker/compose/releases/download/v${DCVERSION}/docker-compose-linux-x86_64"
RUN curl -L ${DCURL} -o /usr/local/bin/docker-compose && chmod +x /usr/local/bin/docker-compose

# docker関連のbash-completion
ARG DCOMPVERSION="1.29.2"
ENV DCOMPURL="https://raw.githubusercontent.com/docker/compose/${DCOMPVERSION}/contrib/completion/bash/docker-compose"
RUN curl -L ${DCOMPURL} -o /etc/bash_completion.d/docker-compose

# goのツール系
COPY --from=go-tool-builder /go/bin/* /go/bin

# [Optional] Uncomment the next lines to use go get to install anything else you need
RUN groupadd docker && usermod -aG docker vscode
USER vscode
RUN echo "complete -C '/usr/local/bin/aws_completer' aws" >> /home/vscode/.bashrc \
    && echo "complete -C '/usr/local/bin/aws_completer' awslocal" >> /home/vscode/.bashrc

COPY resources/home/ /home/vscode
RUN echo . /home/vscode/.aliases >> /home/vscode/.bashrc
