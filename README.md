A RESTful API for creating and managing multiple tor instances

## How to run it
 You will need a local postgres installation with the default settings

 install gopm: `go get -u github.com/gpmgo/gopm`

 install deps: `gopm get`

 install godo `go get -u gopkg.in/godo.v1/cmd/godo`

 install goose `go get bitbucket.org/liamstask/goose/cmd/goose`

 create a database in postgres with the name "distributor"

 Run migrations: `goose up`

 compile: `godo the-thing`

 compile & watch: `godo --watch`

## API Overview

Currently, the API is simple and sweet with your basic CRUD operations:

### Spawn a new tor instance
`POST localhost:8080/node/create`

```
201 Created --
{
  control_port: 9091,
  port: 9091,
}
```

### Get information on a tor instance
`GET localhost/node/{control_port}`

```
200 OK --
{
  control_port: 9091,
  port: 9091,
}
```

### Request a tor instance to cycle it's IP
`PATCH localhost:8080/node/{control_port}`

```
202 Accepted --
```

### Remove a tor instance
`DELETE localhost:8080/node/{control_port}`

```
204 No Content --
```
