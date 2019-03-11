[![CircleCI](https://circleci.com/gh/joonilkim/gstaad/tree/develop.svg?style=svg)](https://circleci.com/gh/joonilkim/gstaad/tree/develop)

# gstaad
This is a PoC project for my MSA project. (ECS + terraform + GRPC)

- ALB for routing
- Envoy - Deploys as sidecar to avoid extra hop and to take service specific things. GRPC for internal communication and REST for frontend serving. Uses envoy's JSON to GRPC transcoder to serve REST. Decodes JWT and send its payload to upstream services for authorization.
- Route53 and CloudMap for Service Discovery and Registry.
- Uses GRPC's load balancing for connection management such as circuit breaking and keepalive.
- Uses `Fargate` to add public ip to `awsvpc` network mode.
- X-Ray - not implemented yet
- Prometheus - not implemented yet