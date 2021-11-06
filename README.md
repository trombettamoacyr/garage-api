# garage-api

This golang REST API was created as a study about Golang concepts and clean architecture.

The garage-api has a method post that allows user to store a car and its features in two different databases. 

Every car saved has its insurance value checked in other external api [myfakeapi/cars](https://myfakeapi.com/api/cars/).

There are two GET methods that show information about the cars already registered. The details are fetch concurrently (Goroutines) in externals APIs. Owner information in [myfakeapi/users](https://myfakeapi.com/api/users/) and car image in [jsonplaceholder](https://jsonplaceholder.typicode.com/photos/).

### View collection

- TODO

### Dependencies:

- [Gorilla Mux](github.com/gorilla/mux) - Router http
- [Chi](https://github.com/go-chi/chi) - Router http
- [Firestore](cloud.google.com/go/firestore) - Firestore driver
- [Postgres](github.com/lib/pq) - Postgres driver
- [Testify](github.com/stretchr/testify) - Packages for Tests

### Install dependencies
``` 
go build
``` 

### Export environment variables
``` 
source .env
```

``` 
export GOOGLE_APPLICATION_CREDENTIALS='/path/to/project-private-key.json'
``` 

### Test (specific test)
``` 
go test -run NameOfTest
``` 

### Test (service folder)
``` 
go test service/*.go
``` 

### Create postgres container
``` 
make docker-start
``` 
Stop container
``` 
make docker-stop
``` 

### Run application
``` 
go run .
```