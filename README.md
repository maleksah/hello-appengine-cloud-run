# Hello Go



## Getting started

This a Hello app developed with go and target to deploy on App engine or Cloud Run

## Deploy on App Engine

First, authenticate with gcloud
`````shell
gcloud auth login
`````

deploy to app engine
````shell
gcloud app deploy -v <VERSION> --project <GCP_PROJECT_ID>
````

TODO: app engine traffic split

# Deploy on Cloud Run
Build docker image
````shell
docker build -t europe-west9-docker.pkg.dev/par-sahnoun-sandbox/google-cloud-onboard/hello:v2 .
````

Push docker image to Artifact registry
````shell
docker push europe-west9-docker.pkg.dev/par-sahnoun-sandbox/google-cloud-onboard/hello:v2 .
````

````shell
gcloud run deploy hello-google-cloud-onboard \
--image europe-west9-docker.pkg.dev/par-sahnoun-sandbox/google-cloud-onboard/hello:v2 \
--region europe-west9 \
--allow-unauthenticated 
````