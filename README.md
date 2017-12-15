# Container Information

Web server used to ack back information from a running container.

```
$ probe -h
Usage of probe:
  -port string
    	Specify the port number you would like to use.  (default "9000")

$ curl localhost:9090
Hello from container: e65f16eb96f7
Local IP of container: 172.17.0.3

```