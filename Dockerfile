FROM python:2.7

RUN apt-get update && apt-get install -y curl unzip

# Download & Install Go
RUN curl -s https://go.googlecode.com/files/go1.1.1.linux-amd64.tar.gz | tar -v -C /usr/local -xz
ENV PATH /usr/local/go/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin
ENV GOPATH /go
ENV CGO_ENABLED 0
RUN cd /tmp && echo 'package main' > t.go && go test -a -i -v

# Download & Install Go AppEngine
RUN curl --remote-name https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.17.zip
RUN unzip go_appengine_sdk_linux_amd64-1.9.17.zip -d /usr/local
ENV PATH /usr/local/go_appengine:$PATH

# Add local files to the Docker container
ADD . /go/src
WORKDIR /go/src