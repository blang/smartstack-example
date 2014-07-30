FROM ubuntu:14.04
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -q

RUN apt-get install -qy build-essential curl git
RUN curl -s https://storage.googleapis.com/golang/go1.3.src.tar.gz | tar -v -C /usr/local -xz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1
ENV PATH /usr/local/go/bin:$PATH
RUN cd /opt/ && git clone https://github.com/coreos/etcd.git etcd && cd ./etcd && ./build

WORKDIR /data
VOLUME ["/data"]

EXPOSE 4001 7001
CMD ["/opt/etcd/bin/etcd", "-data-dir=/data", "-name=master"]