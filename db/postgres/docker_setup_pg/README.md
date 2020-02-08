## Installing postgres with docker

If you need to install Docker:
https://www.docker.com/get-started

```sh
docker pull postgres
mkdir -p $HOME/docker/volumes/postgres
docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
psql -h localhost -U postgres -d postgres

# note: you may have to wait 30 seconds or so for the server to come up
# note: you will be prompted for the password; see above for the password
```

For details of the above, see:
https://hackernoon.com/dont-install-postgres-docker-pull-postgres-bee20e200198

