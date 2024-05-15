# Docker Compose YAML Validator
## Build the Project
```terminal
go build -o docker-compose-validate ./cmd
```
## Run the Validator
```terminal
./docker-compose-validate validate /path/to/docker-compose.dev.yml
```

## How To Use The Makefile 
Build the binary
```terminal
make build
```
Validate a Docker Compose file
```terminal
make run FILE=~/Code/nest-microservices-lt/docker-compose.dev.yml
```

Clean the build
```terminal
make clean
```

Run tests
```terminal
make test
```
