# simple-get-server
Simple server that proxies a pre-defined URL

# Configuration

Configure with the following env vars:

* `URL` (Required): URL of the page to return
* `AUTH_TOKEN` (Optional): HTTP auth-token, defaults to none.

# Build and run locally

```
$ docker build -t simple-get-server .
$ docker run -p8080:8080 -e URL=<YOUR URL> simple-get-server
```

# Deploy to Cloud Run

Click this button:

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://deploy.cloud.run)

Or type the following:

```
$ gcloud beta run deploy --source . --set-env URL=<YOUR URL>
```