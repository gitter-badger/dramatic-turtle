[![Build Status](https://travis-ci.org/HeikoAlexanderWeber/dramatic-turtle.svg?branch=master)](https://travis-ci.org/HeikoAlexanderWeber/dramatic-turtle)
[![Go Report Card](https://goreportcard.com/badge/github.com/HeikoAlexanderWeber/dramatic-turtle)](https://goreportcard.com/report/github.com/HeikoAlexanderWeber/dramatic-turtle)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ce8860f62a4347c296f047b8ceec26a5)](https://www.codacy.com/project/HeikoAlexanderWeber/dramatic-turtle/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=HeikoAlexanderWeber/dramatic-turtle&amp;utm_campaign=Badge_Grade_Dashboard)

# dramatic-turtle

## Getting started

### Prerequisites

* Installed GO SDK
* Installed GNU Make support
* Available cli (bash, powershell etc)

### Starting development

1. Clone this repository to your folder $(folder)
1. Start bash, powershell or the cli of your choice
1. cd into $(folder)
1. Invoke the Makefile routine "init" (make init). </br> This command will download all the dependencies and build the project.

### Deploying in production

#### Build a docker image

To build a docker image you can simply invoke the Makefile routine "docker-build-image" (make docker-build-image).

#### Run the docker image

To run the docker image you can simply invoke the Makefile routine "docker-run-image" (make docker-run-image).
