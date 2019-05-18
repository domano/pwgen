# pwgen
Challenge: Create a password generating REST service

![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/domano/pwgen.svg)
![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/domano/pwgen.svg)

# API

## Parameters
There is one endpoint `/passwords` with the following query parameters.

| Parameter | Description | Default | 
| --- | --- | --- | 
| minLength | The minimum length of a password.| 0| 
| specialChars| Minimum amount of special characters. | 0 | 
| numbers | Minimum amount of numbers. | 0 |

### Example:
Request `/passwords?minLength=10&specialChars=3&numbers=3`

Response `["?!o\10wE9q"]`

 
## run
Following environment variables can be set

| ENV           | Description | Default | Required |
|---            |---                                |---                    |---                |
| CERT_FILE     | Path to TLS cert file.            | cert.pem              | Only for docker   |
| KEY_FILE      | Path to TLS unencrypted key file. | key.unencrypted.pem   | Only for docker   |
| PORT          | Port to listen on.                | 8443                  | No                |
| GRACE_PERIOD  | Timeout for graceful shutdown.    | 5s                    | No                |

###  docker
You can easily run pwgen with the publicly available docker image. 

Because this is a password generator only HTTPS is supported and a TLS certificate and key must be provided.

Example with latest public image:

`docker run -v $(pwd):/certs -e CERT_FILE=/certs/cert.pem -e KEY_FILE=/certs/key.unencrypted.pem domano/pwgen`

### locally
`go run cmd/pwgen/main.go`
