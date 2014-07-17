NS = blang
REPO = smartstack-zookeeper
NAME = smartstack-zookeeper

PORTS = -p 2181:2181 -p 2888:2888 -p 3888:3888

.PHONY: build shell run start stop stoprm rm

build:
	docker build -t $(NS)/$(REPO) .

shell:
	docker run --rm --name $(NAME) -i -t $(PORTS) $(NS)/$(REPO) /bin/bash

run:
	docker run --rm --name $(NAME) $(PORTS) $(NS)/$(REPO)

start:
	docker run -d --name $(NAME) $(PORTS) $(NS)/$(REPO)

stop:
	docker stop $(NAME)

stoprm: stop rm 
	
rm:
	docker rm $(NAME)

default: build