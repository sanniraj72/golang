# Process of Deploying a simple api service (GoLang) using Kubernetes 

**1.** I have created a simple api-service using go. That will show /home page and /countries.

**2.** Create Docker file to contanerize the app.

```
FROM golang:latest
RUN mkdir app && cd app
WORKDIR app/
COPY . .
RUN go build main.go
CMD ["./main"]
EXPOSE 8080
```

**3.** Build this image using 
            ```docker build -t <imagename:tag> <path of Dockerfile>```

### RUN IMAGE USING DOCKER:
   3.1 Run image and expose the port using command
            ```docker run -p 8080:8080 -td <imagename:tagname>```
   3.2 To verify the api, enter into the bash of container using command
            ```docker exec -it <container-id> bash```

### PUSH IMAGE TO DOCKER HUB:
**4.** Login to the docker hub in terminal using ```docker login``` command

**5.** Before pushing to docker hub we have to tag it using -
       ```docker tag <container ID> <reponame:tag> ```
    e.g. - docker tag <container-id> sannidocker/myfirstrepository:myimagev1.0

**6.** After tag now you can push the image using command 
       ```Docker push <image>```

### RUN OR DEPLOY IMAGE USING KUBERNETES:
**7.** Create a YAML file 

```
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
        image: sannidocker/myfirstrepository:myimagev1.0
        ports:
        - containerPort: 80
```

**8.** To create pods with replication controller use command - 
        ```kubectl create -f <filename.yaml>```

**9.** To verify the api, enter into the bash of container using command
        ```docker exec -it <container ID> bash```

**10.** Inside container use curl command to see the result - 
        curl [localhost:8080/home]
        curl [localhost:8080/countries]
        
**NOTE:** When you are giving imagename in pod.yaml file, then it will search locally. But if image not found locally             then to get from your "dockerhub" repository, you should login to dockerhub from terminal and imagename must be             prefix with the repo name.
