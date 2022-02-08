FROM golang:1.8-jessie

LABEL description="Docker image for running myra-upload Golang application"

# GLIDE
ENV GLIDE_VERSION=0.13.3
ENV GLIDE_DOWNLOAD_URL="https://github.com/Masterminds/glide/releases/download/v${GLIDE_VERSION}/glide-v${GLIDE_VERSION}-linux-amd64.tar.gz"
ENV GIT_SSL_NO_VERIFY=1

# PATH
ENV GOPATH="/go"
ENV APP_NAME="github.com/Myra-Security-GmbH/go-upload-client"
ENV APP_PATH="$GOPATH/src/$APP_NAME"

RUN apt update && \
    apt install -y ca-certificates && \
    apt install -y coreutils && \
    apt install -y upx-ucl

ADD  "${GLIDE_DOWNLOAD_URL}" glide.tar.gz

RUN  tar -xzf glide.tar.gz \
    && mv linux-amd64/glide /usr/bin/ \
    && rm -r linux-amd64 \
    && rm glide.tar.gz

WORKDIR $APP_PATH

#VOLUME . /go/src/github.com/Myra-Security-GmbH/go-upload-client
CMD glide install -v; make;  myra-upload -h; tail -f /dev/null
