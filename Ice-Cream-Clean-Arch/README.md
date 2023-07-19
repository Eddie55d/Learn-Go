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
