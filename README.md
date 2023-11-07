# sysdig-hello-world

repo with a test container image to test vulnerabilities pipeline scanning

Sample Go app that runs a simple Golang webserver on port 8090 - Includes a third party library (logrus) for scanning demonstration purposes.

`docker run --name sysdig-hello-world -p 8090:8090 ghcr.io/bashfulrobot/sysdig-hello-world:latest`

## ToDo

- Pin versions for a better demo.
