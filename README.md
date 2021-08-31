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

**Instructions to deploy redisConnect image without cloning repo**

1. Apply the job yaml   
  **oc apply -f https://github.com/SherinV/redisConnect/blob/main/deploy/redisconnect_job.yaml\?raw\=true**
2. You should see a redisconnect job created and a redisconnect-xxx pod starting
3. Collect logs of redisconnect-xxx pod with
  oc logs redisconnect-xxx
