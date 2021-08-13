# URL Shortener

#### System requirements Development
- [x]  Go version 1.16 with Go module enabled

### Pre-Require

Mockery
```
go get github.com/vektra/mockery/v2/.../
```
Swagger
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### Swagger URL
```
/swagger/index.html
```

### Installation

```
cd [project folder path]
go mod tidy
```

### Testing 
unit testing command

```
  go test ./... -cover
```

integrating testing command

```
  go test ./... -tags integration
```


### Generate Mocks

generate mocks from interfaces for unit testing

```
  go generate ./...
```
