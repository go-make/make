FROM golang

ADD . /go/src/gopkg.in/make.v4
RUN \
    apt-get update && \
    apt-get install -y zip unzip gawk python-pip && \
    rm -rf /var/lib/apt/lists/* && \
    curl -sL https://download.docker.com/linux/static/stable/x86_64/docker-17.12.0-ce.tgz | tar -xzf - -O docker/docker > /usr/local/bin/docker && \
    chmod +x /usr/local/bin/docker && \
    pip install docker-compose

RUN make -f /go/src/gopkg.in/make.v4/batteries.mk tools
