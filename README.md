<div style="display: flex;align-items: center; justify-content: center; margin-top: 20px">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/Shell_Script-121011?style=for-the-badge&logo=gnu-bash&logoColor=white" />
  <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" />
  <img src="https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" />
</div>

## GO API 🦫

This project is to create a rest api using go for local my project

#### Basic Command

```bash
# create a new go module
go mod init local_my_api

# to install new package
go get -u github.com/gorilla/mux

# to list all the dependency
go mod vendor

# to install dependency
go mod tidy

# to build
go build -o local_my_api cmd/main.go
```

### Project Structure

This project uses go module and dependency injection pattern for the project structure. Implcitly define interface and struct is a must for the project structure.

#### Folder 📂

- `cmd` - for initial entry point and for building the project
- `internal` - for all the internal package and logic
- `pkg` - for all the external package and logic
- `testing` - for all the testing purpose

#### Routing 🪧

- define all the routes in the `internal/routes` folder and combine them before accessing the server for separation of concern
- data flow from `handler` -> `service` -> `repository`

##### handler 📨

- to handle the request and response
- processing the request and sanitise the input
- validate the input for request received

##### service 🔩

- to handle the business logic
- handle the data action e.g. transaction, event, CRUD
- db access starts from here

##### repositoriy 💽

- to access db
- to preload the data

##### middleware 🧱

- to handle the middleware for the request
- to handle the authentication and authorisation
- to intercept the request and response

#### model 📦

- to define the struct for the data
- to define the interface for the struct

## Development 🛠️

Recommended to use `air` for live-reloading (rebuild all for each changes) and `go run` for running the server without building the project `go run cmd/main.go`

[Air Repository](https://github.com/air-verse/air)

## Migrations 🪽

Migrations by go lang is shit. Implemented custom bin sh for migrations purpose

1. Run `chmod +x go_migrate.sh`
2. Run `./go_migrate --migration_name_by_snake`
3. File name is auto copied to clipboard (install xclip/xsel) (debian: `sudo apt-get install xclip`). Should output like this

```bash
➜ ./go_migrate.sh --create_product_table
Migration file created: internal/db/migrations/20240825014020_create_product_table.go
File name copied to clipboard using xclip.
```

4. Paste the ID in migrationList array in file `internal/db/migrations.go`
5. Edit the migration file created

## Deployment 🐳

This project uses Docker to deploy. To disable web app, comment out in yml file

#### For building local and deploy local

1. Create .env.local with this format for db = `'user=postgres password=postgres dbname=local_my host=db port=5432 sslmode=disable'`
2. If password is changed, change in `docker-compose.yml` as well in user and password
3. Build vite app image (`./build.sh`) and ensure the `local/vite:tag` is the same in docker compose
4. Bump version in `version.txt` and run `./build.sh`
5. Ensure the docker compose version is match in `local/go:tag`
6. Run `docker-compose -f docker-compose.local.yml up -d`

#### For updating production version

##### build image

1. Bump version in `version.txt` and run `./build.sh`
2. Ensure the docker compose version is match in `ghcr.io/hrithiqball/go:tag`
3. Git push to main branch

##### run image

1. Ensure the `ghcr.io/hrithiqball/go:tag` and `ghcr.io/hrithiqball/vite:tag` is the same in docker compose
2. Run `docker-compose -f docker-compose.prod.yml up -d`
