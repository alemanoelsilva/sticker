# sticker

A Go lang backend application to make commitments and stick to it.

The basic configuration is to create a `config.json` file with two properties
- connection_string: a connection mysql string as `user:pas@tcp(127.0.0.1:3306)/db`
- port: application HTTP port

To run locally

```
go build . && go run .
```
