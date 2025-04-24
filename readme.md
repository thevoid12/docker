# setup
- here we have a golang app with pgsql as backend
- we are creating a dev and prod docker deployment setup
reference: https://docs.docker.com/reference/dockerfile/

### some of the docker commands ive used :
```bash
docker exec -it pgsql psql -U postgres -l
docker build -f Dockerfile.golang . # without name and tag
docker build -t godocker:v3 -f ./docker_images/Dockerfile.golang .# with name and tag
docker-compose -f docker-compose-dev.yml up -d
```
#### to setup dev environment
https://github.com/air-verse/air
```bash 
go install github.com/air-verse/air@latest
air init
```
## TODO:
- [ ] set up prod as well as for development
- [ ] as part of development set up 
  - [ ] hot reloading
  - [ ] go debugger
- [ ] manage secrets (pull them out of .env) try using a vault
