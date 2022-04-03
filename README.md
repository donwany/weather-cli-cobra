### Install
```shell
https://github.com/spf13/cobra

go get -u github.com/spf13/cobra@latest
```
### build
```shell
 go run main.go
 go build -o wectl
 go install
```
### run
```shell
./wectl --version

./wectl city plano

  Response - Temperature at plano is 25 degrees celcius

./wectl city --city plano

  Response - Temperature at plano is 25 degrees celcius
```
### deploy
```shell
sudo mv ./wectl /usr/local/bin/

$ wectl city plano 
$ wectl city --city plano
```

## Docker
```shell
docker build -t go-app-img .

docker run -d -p 3333:3000 \ 
--name go-app-container go-app-img

docker build -t worldbosskafka/go-app .

docker push worldbosskafka/go-app:latest
```

## Kubernetes
```shell
kubectl create deployment my-go-app \
    --image=worldbosskafka/go-app:latest \
    --port=3000
    
kubectl create -f deployment.yaml
kubectl get deployments
kubectl get pods
kubectl get svc
minikube ip

kubectl expose deployment my-go-app \ 
   --type=NodePort \ 
   --name=go-app-svc \ 
   --port=3000
   --target-port=3000
```