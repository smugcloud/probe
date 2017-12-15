# Container Information

Web server used to ack back information from a running container.

Current information returned:

* Hostname of the container
* Private IP of the container

## Usage

```
$ probe -h
Usage of probe:
  -port string
    	Specify the port number the server will listen on. (default "9000")

$ curl localhost:9090
Hello from container: e65f16eb96f7
Local IP of container: 172.17.0.3

```