A RESTful API for creating and managing multiple tor instances

### Install Deps
[gopm](https://github.com/gpmgo/gopm):
`$ go get -u github.com/gpmgo/gopm`

[godo](https://github.com/go-godo/godo):
`$ go get -u gopkg.in/godo.v1/cmd/godo`

[goose](https://bitbucket.org/liamstask/goose):
`$ go get bitbucket.org/liamstask/goose/cmd/goose`

### Spawn a tor instance
`POST localhost:8080/node/create`

```
{
  control_port: 9091,
  port: 9091,
}
```

### Get information on a tor instance
`GET localhost/node/{control_port}`

```
{
  control_port: 9091,
  port: 9091,
}
```
