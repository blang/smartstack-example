NS = blang
REPO = smartstack-web
INSTANCE_ID = httpfrontend1
NAME = smartstack-web

PORTS = -p 2222:22 -p 80:80 -p 3212:3212
VOLUMES = -v /tmp/weblogs:/var/log/supervisor
ENV = -e SMARTS_PUBLIC_IP=192.168.0.32 -e SMARTS_PUBLIC_PORT_80=80 -e SMARTS_ZOOKEEPER=192.168.0.32:2181 -e SMARTS_INSTANCE_ID="$(INSTANCE_ID)"

.PHONY: build shell run start stop stoprm rm

build: buildhttpserver
	docker build -t $(NS)/$(REPO) .

buildhttpserver: httpserver/httpserver

httpserver/httpserver: httpserver/main.go
	cd httpserver && go build .

shell:
	docker run --rm --name $(NAME) -i -t $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO) /bin/bash

run:
	docker run --rm --name $(NAME) $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO)

start:
	docker run -d --name $(NAME) $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO)

stop:
	docker stop $(NAME)

stoprm: stop rm 

rm:
	docker rm $(NAME)

default: build