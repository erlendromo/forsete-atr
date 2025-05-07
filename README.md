# forsete-atr

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
