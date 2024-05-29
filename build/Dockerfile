# go version
FROM golang:1.22.3-alpine

# install required packages
RUN apk update && apk add git

# prepare the working directory
RUN mkdir -p /go/src/github.com/yuyaprgrm/text2speech

# set the working directory
WORKDIR /go/src/github.com/yuyaprgrm/text2speech

# copy the source code
COPY . /go/src/github.com/yuyaprgrm/text2speech