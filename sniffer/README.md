# sniffer

## docker image
There are individual make targets to build the go binary (`build`) and create the docker image (`image`) but we combine these two in the make target `build-image`.

To compile your code and create a docker image with this code, run;

    VERSION=v1.0 make build-image

## run
To run the sniffer container with port 9090 exposed, in the `HAProxy` docker network (created earlier), run;

    VERSION=v1.0 make run

The `HAProxy` docker network allows all containers to talk to each other

## stop
To stop a running simple server container, run;

    make stop
