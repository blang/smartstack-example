FROM ubuntu:14.04
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -q

# Supervisor
RUN apt-get install -y supervisor && \
  mkdir -p /var/log/supervisor && \
  mkdir -p /etc/supervisor/conf.d

RUN apt-get install -y software-properties-common
RUN add-apt-repository -y ppa:saltstack/salt

RUN apt-get install -y ssh-client salt-minion 

# Define working directory.
WORKDIR /


# Add Volumes
VOLUME ["/var/log/supervisor"]

CMD ["supervisord"]