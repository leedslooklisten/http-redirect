#!bin/bash

# deploy script for http-redirect web server

if [ "$1" = "" ]; then
    echo "Usage: sh deploy.sh [deployment version]"
    exit
fi


# build docker image
GOOS=linux go build # build linux Go executable

# delete old image
# docker stop http-redirect
# docker rm http-redirect
# docker rmi us.gcr.io/leedslooklisten-com/http-redirect:$1-1

# build container
docker build -t us.gcr.io/leedslooklisten-com/http-redirect:$1 .

# local testing
# docker run -d --name http-redirect -p 8080:80 us.gcr.io/leedslooklisten-com/http-redirect:$1

# push container to GCR
docker push us.gcr.io/leedslooklisten-com/http-redirect:$1

# update container on GCloud VM
gcloud compute instances update-container redirect-instance --container-image=us.gcr.io/leedslooklisten-com/http-redirect:$1



# copy TLS keys to Gcloud instance
# gcloud compute scp /private/etc/letsencrypt/live/leedssound.com/cert.pem benjaminleeds@redirect-instance:~/
# gcloud compute scp /private/etc/letsencrypt/live/leedssound.com/privkey.pem benjaminleeds@redirect-instance:~/