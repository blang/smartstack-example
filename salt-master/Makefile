NS = blang
REPO = smartstack-saltmaster
INSTANCE_ID = defaulinstance
PUBLIC_PORT = 9000
NAME = smartstack-saltmaster-$(INSTANCE_ID)
PORTS = -p $(PUBLIC_PORT):9000
#VOLUMES = -v /tmp/weblogs:/var/log/supervisor
ENV = -e SMARTS_PUBLIC_IP=192.168.0.32 -e SMARTS_PUBLIC_PORT_9000=$(PUBLIC_PORT) -e SMARTS_ZOOKEEPER=192.168.0.32:2181 -e SMARTS_INSTANCE_ID=$(INSTANCE_ID)

.PHONY: build shell run start stop stoprm rm

build: 
	docker build -t $(NS)/$(REPO) .

shell:
	docker run --rm --name $(NAME) -i -t $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO) /bin/bash

run:
	docker run --rm --name $(NAME) $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO)

start:
	docker run -d --name $(NAME) $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO)

stop: rm
	docker stop $(NAME)

rm:
	docker rm $(NAME)

default: build