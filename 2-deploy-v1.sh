#!/bin/env bash

kubectl apply -f 2-thequoter-v1.yaml
set -g -x IP (kubectl get service thequoter-v1 -o json | jq -r ".status.loadBalancer.ingress[0].ip")

curl  -H "Authorization: Bearer $DNSIMPLE_TOKEN" \
      -H "Accept: application/json" \
      -H "Content-Type: application/json" \
      -X POST \
      -d "{ \
  \"name\": \"thequoter-dns.labs\", \
  \"type\": \"A\", \
  \"content\": \"$IP\", \
  \"ttl\": 60 }" \
  https://api.dnsimple.com/v2/5059/zones/stjean.me/records