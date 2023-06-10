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

## Run project

1. [Install Docker Desktop](https://docs.docker.com/desktop/install/windows-install/).
2. [Enable Kubernetes](https://docs.docker.com/desktop/kubernetes/).
3. [Install kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/)
4. Make sure nothing else is running on port 80
5. Run following (powershell)

    ```powershell
    kubectl config use-context docker-desktop
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.0/deploy/static/provider/cloud/deploy.yaml
    kubectl create namespace devops-release
    ```

6. Build docker image v1

    ```powershell
    $image="devops-release:v1"
    docker build -t $image .
    ```

7. Deploy blue

    ```powershell    
    $yml = get-Content -Path  ./deployment/deployment-template-blue.yml
    $yml = $yml.replace("#{image}#", "$image")
    set-Content -Path ./deployment/deployment-blue.yml -Value $yml

    kubectl apply -f ./deployment/deployment-blue.yml -n devops-release
    ```

8. Deploy ingress, routing to blue

    ```powershell
    kubectl apply -f ./deployment/ingress-blue.yml -n devops-release
    ```
    Wait for ingress controller to show ADDRESS localhost
    
    ```powershell
    kubectl get ingress --watch -n devops-release
    # press control+c to stop watching
    ```

9. Browse to http://kubernetes.docker.internal/vehicles

10. Change something in the code, i.e. add Volkswagen ID.5

11. Build docker image v2

    ```powershell
    $image="devops-release:v2"
    docker build -t $image .
    ```

12. Deploy green

    ```powershell
    $yml = get-Content -Path  ./deployment/deployment-template-green.yml
    $yml = $yml.replace("#{image}#", "$image")
    set-Content -Path ./deployment/deployment-green.yml -Value $yml

    kubectl apply -f ./deployment/deployment-green.yml -n devops-release
    ```

    Green has been deployed, but http://kubernetes.docker.internal/vehicles still returns blue version

13. Deploy ingress, routing to green

    ```powershell
    kubectl apply -f ./deployment/ingress-green.yml -n devops-release
    ```
14. Browse to http://kubernetes.docker.internal/vehicles to see green version

15. Blue version is still present. Routing to be reserved back to blue by applying previous ingress version

    ```powershell
    kubectl apply -f ./deployment/ingress-blue.yml -n devops-release
    ```

    If we add an extra route with another hostname, which always routes to green, to both ingress templates, and the same for blue, then we can test the blue or green version without switching the route for other users.

16. Cleanup
    ```powershell
    kubectl delete namespace devops-release
    ```
