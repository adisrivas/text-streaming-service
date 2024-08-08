# Text Streaming Service

Implementation of text streaming endpoint in Golang.

### Features
- It supports mutiple inference services.
- Inference service can be switched depending on following criteria:
   - Current service is down
   - The response has status code other than 200
   - Response time is more than 5 seconds
- Health check is performed at every 5 minutes and the current state is recorded for each inference service.
- Rate limit is enforced for each user depending on the plan they are subscribed to.
- Stubs of each service are implemented for testing purpose.
- Errors are logged of relevant states for debugging purpose.


### Steps
1. **Clone the Repository**

    ```bash
    git clone https://github.com/adisrivas/text-streaming-service.git
    ```

2. **Install Dependencies**
    ```bash
    go mod tidy
    ```

3. **Run**
    ```bash
    go run start/main.go
    ```

### Endpoint
```bash
http://localhost:8000/query?prompt=<>
```
