## Build docker image and run in OCP cluster
Build the image
```
docker build . -t redisconnect
```

Push image to quay.
```
docker login -u {YOUR-QUAY_USER} quay.io
docker tag redisconnect quay.io/${QUAY_USER}/redisconnect:latest 
docker push quay.io/${QUAY_USER}/redisconnect:latest 
```

Run Job in OCP cluster
```
oc login ...
oc apply -f deploy/redisconnect_job.yaml
```
