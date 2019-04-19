1. Created a go application file.

2. create Docker file using that application 
$ cat Dockerfile
FROM golang:latest
RUN mkdir app && cd app
WORKDIR app/
COPY . .
RUN go build main.go
CMD ["./main"]
EXPOSE 8080

3.Build this image using "docker build -t <imagename:tag> <path of Dockerfile>"

4.Login the docker hub in terminal using "docker login"

5. Our own Docker image we need to tag using "docker tag <container ID> <reponam:tag> 

6.Docker push <image>

7.Create a YAML file 

$cat pod.yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: goapi
spec:
  replicas: 3
  selector:
    app: goapi
  template:
    metadata:
      name: goapi
      labels:
        app: goapi
    spec:
      containers:
      - name: web-server
        image: myimage:v1.0
        ports:
        - containerPort: 80

8.kubectl create -f <filename.yaml>

9. To verify the running container 
docker exec -ti <container ID> bash

Using curl command we can verify the application running inside the container
curl localhost:8080/home
