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

---

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

---

### Kubernetes deployment
Ensure that you have mini-kube or another kubernetes cluster running 

First we need to create the secret:
```
kubectl create secret generic stockchecker-secret --from-literal=APIKEY=*** 
```
(add your key in rather than stars, the caps for APIKEY are needed)

If you want to run this with minikube then we need to start up the tunnel for the load balancer in another browser:

```
minikube tunnel
```

Run apply the manifest with

``` 
kubectl apply -f k8s/manifest.yml
```

Then go to your url
`http://localhost:8080`


---

**Running an ingress controller**

Alternatively you could always run an ingress controller via Traefik or nginx.

Setup the nginx ingress controller:
```
minikube addons enable ingress
```

for non-minikube options use 
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.1/deploy/static/provider/cloud/deploy.yaml
```
I have not tested this one extensively though as windows, wsl2, docker and k8s all together was too much for Christmas!

This should hang off the stockchecker-service-np and the stockchecker-ingress. 

Edit the hosts file to add `test.kat` to point at the ip address given in the ingress (if running locally we need to use the minikube ip instead)
```
kubectl get ingress
NAME                                             CLASS   HOSTS      ADDRESS     PORTS   AGE
ingress.networking.k8s.io/stockchecker-ingress   nginx   test.kat   172.0.0.15   80      23m

--or--
minikube ip
```

If everything is configured correctly you should be able to go to 

http://test.kat

(As a note, I never actually got this to run with windows... it would either return nothing for localhost, or just hang for the minikube ip)