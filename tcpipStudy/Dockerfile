FROM ubuntu:latest

RUN apt update &&\
    apt upgrade -y &&\
    apt install iproute2 -y &&\
    apt install curl -y &&\
    apt install git -y &&\
    apt install make -y &&\
    apt install tcpdump -y &&\
    apt install -y wget

RUN wget https://go.dev/dl/go1.21.0.linux-arm64.tar.gz &&\
    tar -C /usr/local -xzf go1.21.0.linux-arm64.tar.gz &&\
    rm go1.21.0.linux-arm64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest &&\
    go install github.com/ramya-rao-a/go-outline@latest &&\
    go install github.com/nsf/gocode@latest &&\
    go install github.com/acroca/go-symbols@latest &&\
    go install github.com/fatih/gomodifytags@latest &&\
    go install github.com/josharian/impl@latest &&\
    go install github.com/haya14busa/goplay/cmd/goplay@latest &&\
    go install github.com/go-delve/delve/cmd/dlv@latest &&\
    go install golang.org/x/lint/golint@latest &&\
    go install golang.org/x/tools/gopls@latest
