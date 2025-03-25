# forsete-atr

## Setup for docker deployment

### Install dependencies
- Open a terminal in the root directory (e.g. where this README.md file is located)
- On Windows, run `scripts/windows.bat` (cmd)
- On Linux, run `./scripts/linux.sh` (/bin/bash)
- On MacOS, run `./scripts/macos.sh` (/bin/bash)

### Environment configuration
- Create a `.env` file in the root directory
- Look at the `example.env` file for the required environment variables
- NOTE: If you have CUDA resources available, set the DEVICE to `cuda` in the `.env` file. If not, set it to `cpu`.

### Usage
- To run the application on `CPU`, run `make composecpu`
- To run the application on `GPU`, run `make composegpu`
- To attach to the logs, run `make attach`
- To stop the container, run `make composedown`

### How to use
- Default IP is `localhost` and default API_PORT is `8080`
- Open a browser and navigate to `http://<IP>:<API_PORT>/forsete-atr/v1/swaggo/`. This will display the API documentation.
- Use the API documentation to interact with the API.
- Alternatively use a REST-client like Postman or Thunderclient to interact with the API.
