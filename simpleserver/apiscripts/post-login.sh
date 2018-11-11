#!/bin/bash

curl -v -H "Content-Type: application/json" -X POST -d '{"email": "jamppa.jamppanen@foo.com", "password":"Jamppa"}' http://localhost:4047/login
