#!/bin/sh
curl -X POST -H 'Content-Type: application/json' http://localhost:8080/v1/log/task -d '{"name": "Some-Task"}'
