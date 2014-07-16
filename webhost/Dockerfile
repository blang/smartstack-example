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

EXPOSE 80


# Define working directory.
WORKDIR /

# Add web application
ADD httpserver/httpserver /usr/local/sbin/

# Add supervisor config
ADD conf/supervisor /etc/supervisor/conf.d/

# Add nerve config
ADD conf/nerve/nerve.conf.json /etc/

CMD ["supervisord"]
