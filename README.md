# pwgen
Challenge: Create a password generating REST service

[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/domano/pwgen.svg)](https://cloud.docker.com/repository/docker/domano/pwgen)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/domano/pwgen.svg)](https://cloud.docker.com/repository/docker/domano/pwgen)
![Docker Stars](https://img.shields.io/docker/stars/domano/pwgen.svg)
[![Build Status](https://travis-ci.org/domano/pwgen.svg?branch=master)](https://travis-ci.org/domano/pwgen)
[![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)

[![GoDoc](https://godoc.org/github.com/domano/pwgen/internal?status.svg)](http://godoc.org/github.com/domano/pwgen/internal)
[![Go Report Card](https://goreportcard.com/badge/github.com/domano/pwgen)](https://goreportcard.com/report/github.com/domano/pwgen)
[![codecov](https://codecov.io/gh/domano/pwgen/branch/master/graph/badge.svg)](https://codecov.io/gh/domano/pwgen)

![GitHub](https://img.shields.io/github/license/domano/pwgen.svg)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/domano/pwgen.svg)
![GitHub last commit](https://img.shields.io/github/last-commit/domano/pwgen.svg)


# API

## Parameters
There is one endpoint `/passwords` with the following query parameters.

| Parameter | Description | Default | 
| --- | --- | --- | 
| minLength | The minimum length of a password.| 0| 
| specialChars| Minimum amount of special characters. | 0 | 
| numbers | Minimum amount of numbers. | 0 |
| amount | Number of passwords that will be returned | 1 |
| swap | Boolean value indicating if random vowels should be swapped for numbers | false |

### Example:
Request `/passwords?minLength=10&specialChars=3&numbers=3&amount=2`

Response `["?!o\10wE9q", "h3{{v9BB3%"]`

 
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
