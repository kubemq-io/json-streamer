# Player-Demo

This application demostrate sending json structs to postgres target with kubemq

## Run 
``` 
kubectl apply -f https://raw.githubusercontent.com/kubemq-io/json-streamer/master/deployment.yaml
```

## Configuration

Command arguments:

-a: kubemq address, default: kubemq-cluster-grpc.kubemq:50000

-q: kubemq queue name, default: songs

-t: target postgres table name, default: songs

