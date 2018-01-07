#!/bin/bash
docker build -f src/app/frontend/Dockerfile -t dl-kops-frontend .
docker tag dl-kops-frontend swiftdiaries/dl-kops-frontend
docker push swiftdiaries/dl-kops-frontend
