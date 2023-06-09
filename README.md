Lifecycle Tester
---

Lifecycle Tester is a lightweight web application designed to test the feature in Kubernetes deployments. It can be used in a container to simulate different application states and observe how Kubernetes behaves when the application is not considered "ready".

The application provides the following endpoints:

| endpoint | description |
| - | - |
| GET / | Returns the current state and request details. |
| POST / | Updates the application state based on the provided JSON payload. |
| GET /readiness | Returns an HTTP 200 status code when the application is considered "ready" and an HTTP 503 status code when it's not.
| GET /liveness | Returns an HTTP 200 status code when the application is considered "live" and an HTTP 503 status code when it's not.

## State

The server's state is controlled by a JSON payload sent to the `POST /` endpoint. The state object has the following fields:

- `is_failed_readiness` (bool): If set to `true`, the server will fail health checks at the `/readiness` endpoint.
- `is_failed_liveness` (bool): If set to `true`, the server will fail health checks at the `/liveness` endpoint.
- `shutdown_delay_seconds` (int64): The number of seconds to delay the server shutdown process after receiving a shutdown signal.

## Environment Variables

The Lifecycle Tester application can be configured using the following environment variables:

| Environment Variable      | Description                                                          | Default |
|---------------------------|----------------------------------------------------------------------|---------|
| DISABLE_HEALTH_LOG        | Set to "true" to disable logging of health check requests            | "false" |
| SHUTDOWN_DELAY_SECONDS    | Number of seconds to delay the server shutdown process after receiving a shutdown signal | "0" |
| READINESS_DELAY_SECONDS   | Number of seconds to delay the server readiness state                | "0" |
| LIVENESS_DELAY_SECONDS    | Number of seconds to delay the server liveness state                 | "0" |

## Usage

### Web API

This request will return the server information, including the hostname, request details, and current state.

```bash
curl http://localhost:8080
```

This request will set is_failed_liveness to true, causing the server to fail health checks, and shutdown_delay_seconds to 30, causing the server to delay the shutdown process by 30 seconds after receiving a shutdown signal.

```
curl -X POST -H "Content-Type: application/json" -d '{"is_failed_liveness": true, "shutdown_delay_seconds": 30}' http://localhost:8080
```

### CLI Usage

The CLI tool provides various subcommands for different functionalities:

- `server`: Start the HTTP server to provide an interface for updating the application's state and handling container lifecycle events.
- `state`: Update the application's state, such as changing the status of readiness and liveness probes.
- `sleep`: Make the application sleep for a specified duration in seconds or indefinitely using the 'infinity' keyword.

#### Server

This command will start the HTTP server.

```
lifecycle-tester server
```

#### State

This command will set `is_failed_readiness` to `true`, causing the server to fail health checks at the `/readiness` endpoint, and `shutdown_delay_seconds` to 30, causing the server to delay the shutdown process by 30 seconds after receiving a shutdown signal.

```
lifecycle-tester state --is-failed-readiness=true --shutdown-delay-seconds=30
```

#### Sleep

This command will make the application sleep for 10 seconds.

```
lifecycle-tester sleep 10
```

This command will make the application sleep indefinitely.

```
lifecycle-tester sleep infinity
```
