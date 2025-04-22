# setup
- here we have a golang app with pgsql as backend
- we are creating a dev and prod docker deployment setup
reference: https://docs.docker.com/reference/dockerfile/

### some of the docker commands ive used :
```bash
docker exec -it pgsql psql -U postgres -l
docker build -f Dockerfile.golang . # without name and tag
docker build -t godocker:v1 -f Dockerfile.golang . # with name and tag

```
