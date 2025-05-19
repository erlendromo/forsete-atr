# forsete-atr

## Repository content
```
.
├── LICENSE
├── Makefile
├── README.md # This file
├── assets
│   ├── migrations
│   │   ├── 00001_init.down.sql
│   │   ├── 00001_init.up.sql
│   │   ├── 00002_indexes.down.sql
│   │   ├── 00002_indexes.up.sql
│   │   ├── 00003_add_samples.down.sql
│   │   └── 00003_add_samples.up.sql
│   └── scripts
│       ├── htrflow.sh # Run htrflow script
│       └── linux.sh # Setup linux (ubuntu) environment script
├── backend
│   ├── README.md
│   ├── deployments
│   │   ├── application
│   │   │   ├── README.md
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   └── frontend
│   │       ├── README.md
│   │       ├── main.tf
│   │       ├── outputs.tf
│   │       └── variables.tf
│   ├── locals.tf
│   ├── main.tf
│   ├── modules
│   │   ├── instance
│   │   │   ├── README.md
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   ├── terraform.tf
│   │   │   └── variables.tf
│   │   ├── ip
│   │   │   ├── README.md
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   ├── terraform.tf
│   │   │   └── variables.tf
│   │   ├── keypair
│   │   │   ├── README.md
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   ├── terraform.tf
│   │   │   └── variables.tf
│   │   ├── network
│   │   │   ├── README.md
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   ├── terraform.tf
│   │   │   └── variables.tf
│   │   ├── port
│   │   │   ├── README.md
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   ├── terraform.tf
│   │   │   └── variables.tf
│   │   └── security_group
│   │       ├── README.md
│   │       ├── main.tf
│   │       ├── outputs.tf
│   │       ├── terraform.tf
│   │       └── variables.tf
│   ├── outputs.tf
│   ├── providers.tf
│   ├── terraform.tf
│   ├── .terraform.lock.hcl
│   ├── terraform.tfvars # Needs to be added manually
│   └── variables.tf
├── docker
│   ├── Dockerfile
│   ├── docker-compose-cuda.yaml
│   └── requirements.txt
├── docker-compose.yaml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── example.env # Used to create the actual .env file
├── go.mod
├── go.sum
├── main.go
├── src
│   ├── api
│   │   ├── app_context
│   │   │   └── app_context.go
│   │   ├── handler
│   │   │   └── v2
│   │   │       ├── atr
│   │   │       │   └── atrhandler.go
│   │   │       ├── auth
│   │   │       │   └── authhandler.go
│   │   │       ├── image
│   │   │       │   └── imagehandler.go
│   │   │       ├── model
│   │   │       │   └── modelhandler.go
│   │   │       ├── output
│   │   │       │   └── outputhandler.go
│   │   │       └── status
│   │   │           └── statushandler.go
│   │   ├── middleware
│   │   │   ├── auth.go
│   │   │   ├── context.go
│   │   │   ├── logger.go
│   │   │   └── statusresponsewriter.go
│   │   └── router
│   │       ├── httprouter
│   │       │   ├── httprouter.go
│   │       │   └── httprouter_test.go
│   │       ├── httpsrouter
│   │       │   ├── httpsrouter.go
│   │       │   └── httpsrouter_test.go
│   │       └── router.go
│   ├── business
│   │   ├── domain
│   │   │   ├── image
│   │   │   │   └── image.go
│   │   │   ├── model
│   │   │   │   └── model.go
│   │   │   ├── output
│   │   │   │   └── output.go
│   │   │   ├── pipeline
│   │   │   │   ├── pipeline.go
│   │   │   │   └── step
│   │   │   │       ├── exportstep.go
│   │   │   │       ├── modelstep.go
│   │   │   │       ├── orderstep.go
│   │   │   │       └── step.go
│   │   │   ├── session
│   │   │   │   └── session.go
│   │   │   └── user
│   │   │       └── user.go
│   │   └── usecase
│   │       ├── querier # Uses database or in-memory-mock
│   │       │   ├── image
│   │       │   │   ├── mock.go
│   │       │   │   ├── querier.go
│   │       │   │   └── sqlx.go
│   │       │   ├── model
│   │       │   │   ├── mock.go
│   │       │   │   ├── querier.go
│   │       │   │   └── sqlx.go
│   │       │   ├── output
│   │       │   │   ├── mock.go
│   │       │   │   ├── querier.go
│   │       │   │   └── sqlx.go
│   │       │   ├── pipeline
│   │       │   │   ├── mock.go
│   │       │   │   ├── querier.go
│   │       │   │   └── sqlx.go
│   │       │   ├── session
│   │       │   │   ├── mock.go
│   │       │   │   ├── querier.go
│   │       │   │   └── sqlx.go
│   │       │   └── user
│   │       │       ├── mock.go
│   │       │       ├── querier.go
│   │       │       └── sqlx.go
│   │       ├── repository # Uses queriers
│   │       │   ├── image
│   │       │   │   ├── repository.go
│   │       │   │   └── repository_test.go
│   │       │   ├── model
│   │       │   │   ├── repository.go
│   │       │   │   └── repository_test.go
│   │       │   ├── output
│   │       │   │   ├── repository.go
│   │       │   │   └── repository_test.go
│   │       │   ├── pipeline
│   │       │   │   ├── repository.go
│   │       │   │   └── repository_test.go
│   │       │   ├── session
│   │       │   │   ├── repository.go
│   │       │   │   └── repository_test.go
│   │       │   └── user
│   │       │       ├── repository.go
│   │       │       └── repository_test.go
│   │       └── service # Uses repositories
│   │           ├── atr
│   │           │   └── service.go
│   │           └── auth
│   │               └── service.go
│   ├── cmd
│   │   └── app.go
│   ├── config
│   │   ├── api
│   │   │   └── api_config.go
│   │   ├── config.go
│   │   └── db
│   │       └── db_config.go
│   ├── database
│   │   ├── database.go
│   │   ├── mock
│   │   │   └── mock.go
│   │   └── postgresql
│   │       └── postgresql.go
│   └── util
│       ├── constant.go
│       ├── decoder.go
│       ├── errors.go
│       ├── logdata.go
│       ├── timer.go
│       ├── util.go
│       └── writer.go
```

## Setup for docker deployment

### Prerequisites
- The application needs some dependencies to run. Amongst these nvidia drivers for cuda GPU resources, docker (docker compose) and more. Running `assets/scripts/linux.sh` will download the dependencies automatically, but the script only works on linux systems, preferably ubuntu or similar distributions. It is therefore `highly recommended` to setup and run this application on a `vm with ubuntu`.

### Install dependencies
- Open a terminal in the root directory (e.g. where this README.md file is located)
- On Linux, run `./assets/scripts/linux.sh` (/bin/bash)

### Environment configuration
- The `linux.sh` should now have made a .env file based on the example.env provided in the root directory.
- Change the environment-variables to your liking and setup is complete. If the .env is not present, follow these steps:
  - Create a `.env` file in the root directory
  - Look at the `example.env` file for the required environment variables

- NOTE: If the service is to be deployed on Openstack, the `API_PORT` must be the same as the `application_port` in the terraform configuration. The default value is `8080` for both.
- NOTE: If the service is to be run with `cpu resources (NOT cuda)`, the `TIMEOUT` variable should be set above `10m`, as ATR on images take a long time without gpu resources available.
- NOTE: If the application is run on a distributed cloud service (e.g. Openstack, Azure etc.), make sure to allow `Ingress` on the `API_PORT`.

### Add models
- TODO: write about this

### Usage
- With `make` (NOTE: beware that database will be reset when running `make composedown`):
  - To run the application on `CPU`, run `make composecpu`
  - To run the application on `CUDA`, run `make composecuda`
  - To attach the container logs, run `make attachlogs`
  - To stop the container, run `make composedown`

- Manually:
  - To run the application on `CPU`, run `docker compose -f docker-compose.yaml up --build -d`
  - To run the application on `CUDA`, run `docker compose -f docker-compose.yaml -f docker/docker-compose-cuda.yaml up --build -d`
  - To attach the container logs, run `docker compose logs -f forsete-atr`
  - To stop the application and remove all resources, run `docker compose down --volumes --remove-orphans` (NOTE: This will also remove all data from the database, so use with caution)

### How to use
- Default IP is `localhost` and default API_PORT is `8080`
- If the application is hosted on a vm, use the floating IP address as ip.
- Open a browser and navigate to `http://<IP>:<API_PORT>/forsete-atr/v2/swaggo/`. This will display the API documentation.
- Use the API documentation to interact with the API.
- Alternatively use a REST-client like Postman or Thunderclient to interact with the API.
