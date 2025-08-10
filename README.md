# TMF Service application
- This application will contain TMF641, TMF638 and TMF633 specification.
- Application is develop in golang and use [modular monolith architecture boilerplate](https://github.com/ArkjuniorK/gomomo) by ArkjuniorK
## Modules
- ### Service inventory
    Swagger doc http | https://{host}/tmf-api/serviceInventory/v4/docs


# Setup notes (for develop)

## Run tmf-service application
```aiignore
docker-compose up -d
```
### Development
```aiignore
docker-compose -f docker-compose.yml -f docker-compose-dev.yml up -d
```

### Install G (Go version manager)
```bash
curl -sSL https://git.io/g-install | sh -s -- bash
```
```bash
g set <version of go>
```

### Install Task
- Install task in your local machine, in order to do that you can follow [installation instructions](https://taskfile.dev/#/installation)\
```bash
$ brew install go-task/tap/go-task 
```

```
  $ task --list
  task: Available tasks for this project:
    * build:        Build the app
    * run:          Run the app
    * swagger.doc:      Doc for swagger
    * swagger.gen:      generate Go code
    * swagger.validate:     Validate swagger
```
### Install Go Swagger

- Install ```go-swagger``` tool. [Installation page](https://goswagger.io/go-swagger/install/)
```aiignore
sudo apt update
sudo apt install -y apt-transport-https gnupg curl debian-keyring debian-archive-keyring
```

- Register GPG signing key
```aiignore
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | sudo gpg --dearmor -o /usr/share/keyrings/go-swagger-go-swagger-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' | sudo tee /etc/apt/sources.list.d/go-swagger-go-swagger.list
```
- Install
```aiignore
sudo apt update 
sudo apt install swagger
```


