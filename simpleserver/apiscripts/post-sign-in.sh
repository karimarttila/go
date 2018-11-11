#!/bin/bash

curl -v -H "Content-Type: application/json" -X POST -d '{"first-name":"Jamppa", "last-name":"Jamppanen", "email": "jamppa.jamppanen@foo.com", "password":"Jamppa"}' http://localhost:4047/signin
