# go-urlshortener
Learning golang with urlshortener

# Working

This is simple rest api using `mux` and `golang`. Give support few api endpoints

* `/shortener/urls`
  - List all url in database
* `/shortener/addUrl`
  - Add new url to database
  - if shortenCode is not provided automatically generated.
* `/shortener/deleteUl/{shortenCode}`
  - delete existing url from database
* `/{shortenCode}`
  - Redirect based on shortenCode
  
# Installation
to install this use golang's unoffical package manager `dep`. All necessary depends files will come in `vendor` directory

```shell
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

# Running
```shell
go run main.go
```

# Test
In progress.
Test are written in `BDD` using `ginkgo`
```shell
  cd /test
  ginkgo
```
