# Go Country Infos

## How to build and run the application

### Docker
Download this repository and run
````shell
docker build -t countries .
docker run -it -t countries /bin/sh
app/country-info
````

### With local go installation
If you don't have docker installed, you can run with a local go installation.

#### Install golang
Go to [https://golang.org/doc/install](https://golang.org/doc/install) and follow the instructions for your OS

#### Run
````shell
go mod vendor
go mod download
go run entrypoints/capitals/main.go
````

## How to interact with the application
***N*** is the number of countries to be displayed. *Optional argument*

#### Docker
Inside the container, run ``app/country-info capitals *N*``

**Example**
```shell
docker build -t countries .
docker run -it -t countries /bin/sh
app/country-info capitals 5
```

#### Local Go installation
`` go run entrypoints/capitals/main.go capitals N``

**Example**
``go run entrypoints/capitals/main.go capitals 3``