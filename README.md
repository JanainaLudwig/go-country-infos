# Go Country Infos

## How to run
Create a **/config/.env** file. There is an example available inside the config folder.

### Docker
Download this repository and run
````shell
docker-compose -f docker\docker-compose.yaml up --build
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
