#### Init

```bash
go mod init local_my_api
```

### Gorilla Mux

```bash
go get -u github.com/gorilla/mux
```

### File Structure

#### routing

define all the routes in the routing folder and combine them before accessing the server for separation of concern

data flow from handler -> service -> model

##### handler

to handle the request and response
processing the request and sanitise the input
validate the input for request received

##### service

to handle the business logic
handle the data action e.g. transaction, event, CRUD
db access starts from here

##### model

to define the data structure
to access db
to preload the data

##### to run the server

```bash
# watch for changes
air

# run the server
go run main.go
```

##### to install dependencies

```bash
go mod tidy
```

##### to build

```bash
go build
./local_my_api
```
