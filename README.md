#Go - MUX

### A typical top-level directory layout

    .
    ├── main.go                   # initialize and run in main.go
    ├── app.go                    # Go Mux and config
    ├── model.go                  # product model
    ├── main_test.go              # unit testing for handle in app.go
    └── README.md

### ENV set on shell
```

export APP_DB_USERNAME=postgres
export APP_DB_PASSWORD=
export APP_DB_NAME=postgres

```

### Run
```
docker run --name postgres -e POSTGRES_PASSWORD= -d postgres
go test -v
```