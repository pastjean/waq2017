# Waqquoter

Build and push:
```sh
cd server/
make v1
or
make v2
```

Deploy in  k8s cluster:
```sh
kubectl apply -f 1-*
```

Reports:
```sh
echo "GET http://thequoter.labs.stjean.me/" | vegeta attack -duration=5s | tee results.bin | vegeta report
vegeta report -inputs=results.bin -reporter=json > metrics.json
```