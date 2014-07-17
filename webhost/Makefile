NS = blang
REPO = smartstack-web
NAME = smartstack-web

PORTS = -p 2222:22 -p 80:80 -p 3212:3212
VOLUMES = -v /tmp/weblogs:/var/log/supervisor

.PHONY: build shell run start stop stoprm rm

build: buildhttpserver
	docker build -t $(NS)/$(REPO) .

buildhttpserver: httpserver/httpserver

httpserver/httpserver: httpserver/main.go
	cd httpserver && go build .

shell:
	docker run --rm --name $(NAME) -i -t $(PORTS) $(VOLUMES) $(NS)/$(REPO) /bin/bash

run:
	docker run --rm --name $(NAME) $(PORTS) $(VOLUMES) $(NS)/$(REPO)

start:
	docker run -d --name $(NAME) $(PORTS) $(VOLUMES) $(NS)/$(REPO)

stop:
	docker stop $(NAME)

stoprm: stop rm 

rm:
	docker rm $(NAME)

default: build