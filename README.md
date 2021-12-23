# Stock Checker

A simple app which gets a request from the internet, processes and provides a json response.

## To run

### Basic go
You will need to add some variables to your environment for this to work:

```
export SYMBOL=MSFT
export NDAYS=3
export APIKEY=***
```

Then you can run this with:
- `go get .`   Sorts dependencies
- `go run .`   To run the web server

Then you should be able to navigate to localhost:8080 to see the results

### Docker container
Build the docker container with

```
docker build -t local/stockchecker .
```
(add the --no-cache argument if you want to ensure you have a clean build)

Update the .env file to have your key in there for the api rather than the stars

To run:
```
docker run --rm --env-file .env -p 8080:8080 local/stockchecker
```

### Kubernetes deployment
Ensure that you have mini-kube or another kubernetes cluster running 

First we need to create the secret:
```
kubectl create secret generic stockchecker-secret --from-literal=apikey=******** 
```
(add your key in rather than stars)

Now we need to ensure we have any kind of ingress controller - so for either windows or mac we can use 
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.1/deploy/static/provider/cloud/deploy.yaml
```
See: https://github.com/kubernetes/ingress-nginx/issues/7686 for more info 
and 
https://stackoverflow.com/questions/65193758/enable-ingress-controller-on-docker-desktop-with-wls2

Run apply the manifest with

``` 
kubectl apply -f k8s/manifest.yml
```

Windows: `http://kubernetes.docker.internal/stockchecker`
Mac: `http://localhost/stockchecker`
```