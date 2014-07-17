NS = blang
REPO = smartstack-service
NAME = smartstack-service

PORTS = -p 2223:22 -p 9000:9000 -p 3215:3215
VOLUMES = -v /tmp/weblogs:/var/log/supervisor

.PHONY: build shell run start stop stoprm rm

build: buildhttpservice
	docker build -t $(NS)/$(REPO) .

buildhttpservice: httpservice/httpservice

httpservice/httpservice: httpservice/main.go
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