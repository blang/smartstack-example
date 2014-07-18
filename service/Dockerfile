FROM ubuntu:14.04
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -q

# Supervisor
RUN apt-get install -y supervisor && \
  mkdir -p /var/log/supervisor && \
  mkdir -p /etc/supervisor/conf.d

# Install ruby
RUN apt-get install -y ruby1.9.1-dev patchutils build-essential

# Install nerve
RUN gem install nerve

# Install HAProxy
RUN apt-get install -y haproxy

EXPOSE 3215

# Install synapse
RUN gem install synapse

EXPOSE 9000

# SSH
RUN apt-get install -y openssh-server && \
  mkdir /var/run/sshd && \
  echo 'root:root' |chpasswd && \
  sed -ri 's/^session\s+required\s+pam_loginuid.so$/session optional pam_loginuid.so/' /etc/pam.d/sshd && \
  sed -i 's/without-password/yes/' /etc/ssh/sshd_config && \
  sed -i 's/UsePAM yes/UsePAM no/' /etc/ssh/sshd_config

EXPOSE 22

ADD conf/ssh /root/.ssh/

# Define working directory.
WORKDIR /

# Add web application
ADD httpservice/httpservice /usr/local/sbin/

# Add supervisor config
ADD conf/supervisor /etc/supervisor/conf.d/

# Add nerve config
ADD conf/nerve/nerve.conf.json /etc/

# Add synapse config
ADD conf/synapse/synapse.conf.json /etc/

# Add haproxy config
ADD conf/haproxy /etc/haproxy/

# Add scripts
ADD bin /usr/local/bin/

# Add Volumes
VOLUME ["/var/log/supervisor"]

CMD ["supervisord"]
