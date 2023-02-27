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
gcloud app deploy --project <GCP_PROJECT_ID>
````