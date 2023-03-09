# http-packet-proxy

This repository is an example of running HAProxy with traffic mirroring of HTTP traffic.

The repository is made up of three components;

- A generic server
- HAProxy
- A server that receives mirrored traffic, which we are calling `sniffer`.

These three are all running in their own docker containers, exposing ports to be accessed on.

## Prerequisites
Before you run any of the containers, you will need to create a docker network by running;

    docker network create --driver=bridge haproxy


## Running

Once all three containers are up and running, watch the logs on all the containers. 

Next, send traffic to your server. 

You should see data appearing on all logs but in particular the `sniffer` pod should be printing out data stored in its in-memory "cache". 

Request and response data is stored together based on a `unique-id` passed in by HAProxy