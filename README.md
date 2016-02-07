[![Build Status](https://travis-ci.org/Pholey/distribuTor.svg?branch=master)](https://travis-ci.org/Pholey/distribuTor)

A simple server for creating & managing multiple Tor instances

## Getting started:
  You will need tor as well as postgres pre-installed with no special configurations

  Install Deps:
  `$ ./script/bootstrap`

  Provision database:
  `$ ./script/recycle`

  Test it:
  `$ ./script/test`

  Build it:
  `$ go build`

  Run it:
  `./distribuTor`

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
