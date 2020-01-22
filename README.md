# Shortest flight route

This application provides an API for searching the shortest routes between airports.

### User Stories (Requirements)

- As a user I can make a GET request to an endpoint with an `origin` and `destination` query parameter, and receive back the shortest route between the two, as an array of connecting flights. A shortest route is defined as the route with the fewest connections. If there are mulitple routes with the same number of connections, you may choose any of them. 
- As a user I am provided meaningful feedback should no route exist between the airports.
- As a user I am provided meaningful feedback if an error occurred with my request.

> NOTE: THE SHORTEST PATH FUNCTIONALITY MAY NOT RELY ON AN EXTERNAL LIBRARY. YOU MUST DEVELOP THIS ON YOUR OWN.

### Testing

Within the `data` folder you will find two subdirectories, `test` and `full`. `test` offers a small subset of the data found in `full` which you may use when developing your solution. Your final solution however must be performant with the `full` set.

Some test cases to consider on the `test` data set.

| Origin | Destination | Expected Result          |
|--------|-------------|--------------------------|
| YYZ    | JFK         | YYZ -> JFK               |
| YYZ    | YVR         | YYZ -> JFK -> LAX -> YVR |
| YYZ    | ORD         | No Route                 |
| XXX    | ORD         | Invalid Origin           |
| ORD    | XXX         | Invalid Destination      |

## Running and testing the app

### Requirements

1. Docker
2. Docker compose
3. Go 1.12+

### Building and executing

Open a terminal window, navigate to the source root and run the following command:

```sh
docker-compose up
```

This will spin up a MySQL container, populate the database, build the application and execute it.

To stop the application run the following command:

```sh
docker-compose down
```

It is also possible to execute the application alone, or targeting another db host or container, updating the config file and executing the application as follows:

```sh
docker build --tag sfp-api .
docker run -d sfp-api
```

### Testing the app

The server listens on port `5000` by default. You can change this in `config.toml` file.

The app contains unit tests for the core logic. Those tests can be run from the source root path with:

```sh
go test ./...
```

#### Routes

All routes are prefixed using the prefix option in the server section of the config file.

##### `GET /v1/routes`

This route returns the shortest route between two airports if one exists.

##### Params

**origin:** (query) The route origin airport IATA 3 code
**destination:** (query) The route destination airport IATA 3 code

###### Results

**200 OK:** Success

```json
{"flights": [{"origin":"JFK","destination":"YYZ","airline":"AA"}]}
```

**404 Not Found:** No route exists

```json
{"error":true,"message":"no route exists between origin and destination"}
```

**422 Unprocessable Entity:** Validation error

```json
{"error":true,"message":"invalid origin"}
```

**503 Service Unavailable:** Other errors, like database connections, etc.

```json
{"error":true,"message":"Service Unavailable"}
```

### Configuration

> ***Note:** This application uses a single configuration file `config.toml` to run. Although there are secrets in the config file, kept there for simplicity in this project, this is not suitable for production. A better way of handling that is to use environment variables or a configuration service, storing only configuration keys in the config file.*

The config file contains options for setting DB connection and server options.

The server also supports loading configuration values from environment variables, overriding the config file values. For example, it is possible to override the database password with an environment variable named `API_DB_PASSWORD`. It is also possible to change the server prefix with `API_SERVER_PREFIX`. This feature is available for all configuration options.

### TODO: Improvements

- [ ] Add integration tests
- [ ] Add caching on repositories
- [ ] Add caching on API results
