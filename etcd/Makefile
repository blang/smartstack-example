NS = blang
REPO = smartstack-etcd
NAME = smartstack-etcd

PORTS = -p 4001:4001 -p 7001:7001
VOLUMES = -v /tmp/etcddata:/data

.PHONY: build shell run start stop stoprm rm

build:
	docker build -t $(NS)/$(REPO) .

shell:
	docker run --rm --name $(NAME) -i -t $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO) /bin/bash

run:
	docker run --rm --name $(NAME) $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO)

start:
	docker run -d --name $(NAME) $(PORTS) $(VOLUMES) $(ENV) $(NS)/$(REPO)

stop:
	docker stop $(NAME)

rm:
	docker rm $(NAME)

default: build