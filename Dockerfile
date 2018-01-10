FROM golang

ADD . /go/src/gopkg.in/make.v4
RUN \
    apt-get update && \
    apt-get install -y unzip gawk && \
    rm -rf /var/lib/apt/lists/*

RUN make -f /go/src/gopkg.in/make.v4/batteries.mk tools
