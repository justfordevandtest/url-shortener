# URL Shortener

#### System requirements Development
- [x]  Go version 1.16 with Go module enabled

### Prerequisite

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
[base-url]/swagger/index.html
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

### Local development
development in local start mongodb redis

```
cd [project folder path]/development
docker-compose up -d
```

### Generate Mocks

generate mocks from interfaces for unit testing

```
  go generate ./...
```

### Component Diagram

[![Component Diagram](https://github.com/justfordevandtest/url-shortener/blob/master/ComponentDiagram.png?raw=true "Component Diagram")](https://github.com/justfordevandtest/url-shortener/blob/master/ComponentDiagram.png?raw=true)
