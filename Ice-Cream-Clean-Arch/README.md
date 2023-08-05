# ice cream app

## Config
Before starting, replace the default values ​​​​in the configDB.json file
on the values ​​of your environment.
Path: /configs/configDB.json

## GO 
- RUN
```
go run cmd/ice-cream-app/main.go
```

- BUILD
```
go build -o /bin/icecream-app cmd/ice-cream-app/main.go 
```

- EXEC
```
./bin/icecream-app
```

## DOCKER

- BUILD
```
docker build -t icecream .
```

- RUN
```
docker run -p 8080:8080 --name icecream-test -idt --mount type=bind,source="$(pwd)\configs\configDB.json",target=/opt/app/configs/configDB.json icecream
```

## MINIKUBE

Before starting, replace the default values ​​​​in the secret.yaml file
on the values ​​of your environment.
Path: /deploy/k8s/secret.yaml

- CREATE
```
kubectl apply -f ./deploy/k8s/service.yaml
kubectl apply -f ./deploy/k8s/secret.yaml
kubectl apply -f ./deploy/k8s/deployment.yaml
# tunel
minikube service icecream-app
```
- DELETE
```
kubectl delete -f ./deploy/k8s/service.yaml
kubectl delete -f ./deploy/k8s/secret.yaml
kubectl delete -f ./deploy/k8s/deployment.yaml
```
