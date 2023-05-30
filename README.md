# DevOps Release

You are part of a development team maintaining HTTP services. Some of your colleagues have developed a new Vehicle API in Go, and you should help them to deploy it to production.

## Project Setup

1. [Install Go](https://go.dev/).

2. From the project's root folder, install its dependencies:

    ```shell
    $ go get -t gitlab.com/enervalis-public/devops-release
    ```

3. Run the test suite as follows:

    ```shell
    $ go test -v ./...
    ```

4. Start the service:

    ```shell
    $ go run .
    ```

5. Now, you should be able to query the service in the port 8080:

    ```shell
    $ curl localhost:8080/vehicles
    [
        {
            "id": "1",
            "model": "Tesla Model Y",
            "maker": "Tesla"
        },
        {
            "id": "2",
            "model": "Tesla Model 3",
            "maker": "Tesla"
        },
        {
            "id": "3",
            "model": "Fiat 500e",
            "maker": "Fiat"
        },
        {
            "id": "4",
            "model": "Peugeot e-208",
            "maker": "Peugeot"
        },
        {
            "id": "5",
            "model": "Volkswagen ID.4",
            "maker": "Volkswagen"
        }
    ]
    ```

## Your Task

Your goal is to deploy it to a Kubernetes cluster following the blue-green deployment strategy.

We would like to run the outcome of your work, so we ask you to document how to set it up using either Minikube or Docker for Desktop.

- If you make any changes to the application code, ensure the test suite continues to pass.
- Document how to configure and deploy the application to a local cluster.
- Keep a clean Git history.
