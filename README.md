lifecycle-tester
---

Lifecycle Tester is a lightweight web application designed to test the feature in Kubernetes deployments. It can be used in a container to simulate different application states and observe how Kubernetes behaves when the application is not considered "ready".

The application provides the following endpoints:

| endpoint | description |
| - | - |
| GET / | Returns the current state and request details. |
| POST / | Updates the application state based on the provided JSON payload. |
| GET /healthz | Returns an HTTP 200 status code when the application is considered "ready" and an HTTP 503 status code when it's not.

## State

The application maintains a simple state:

```json
{
  "is_failed_healthz": false
}
```

- `is_failed_healthz`: Determines whether the /healthz endpoint should indicate a failure.

## Usage

To update the state, send a POST request to the / endpoint with the desired state in JSON format:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"is_failed_healthz": true}' http://localhost:8080/
```

This command sets the `is_failed_healthz` state to true, causing the `/healthz` endpoint to return an HTTP 503 status code.
