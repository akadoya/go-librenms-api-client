# Librenms API Client (Go)

Librenms Golang API Client


## Docker Compose

`docker_compose` directory consists of local librenms test instance set up.

```
cd docker_compose
docker-compose up -d
docker-compose logs -f
  wait until the initialization finishes
```

your local librenms instance will be accessible at http://localhost:8000/.
create API token and start testing.
