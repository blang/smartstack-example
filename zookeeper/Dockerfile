FROM ubuntu:14.04
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -q

RUN apt-get install -y openjdk-7-jre-headless wget
RUN wget -q -O - http://mirror.dkd.de/apache/zookeeper/zookeeper-3.4.6/zookeeper-3.4.6.tar.gz | tar -xzf - -C /opt
RUN cp /opt/zookeeper-3.4.6/conf/zoo_sample.cfg /opt/zookeeper-3.4.6/conf/zoo.cfg

ENV JAVA_HOME /usr/lib/jvm/java-7-openjdk-amd64

EXPOSE 2181 2888 3888

ADD conf/zookeeper /opt/zookeeper-3.4.6/conf/

CMD ["/opt/zookeeper-3.4.6/bin/zkServer.sh","start-foreground"]
