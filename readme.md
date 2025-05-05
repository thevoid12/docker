# docker setup
## windows/mac
- install docker destop version
- for windows it creates a linux subsystem to run docker
## debian
- there are multiple ways to install including docker destop version
- here i am implementing with starter script
```bash
 sudo apt install curl
 curl -fsSL https://get.docker.com -o get-docker.sh
 sudo usermod -aG docker ${USER}
 newgrp docker
 docker --version
 docker ps
 sudo systemctl enable docker
 sudo systemctl start docker
 sudo apt-get update
sudo apt-get install docker-compose-plugin
```
- here we have a golang app with pgsql as backend
- we are creating a dev and prod docker deployment setup
reference: https://docs.docker.com/reference/dockerfile/

# postgres docker url
DATABASE_URL=postgres://postgres:postgres@pgsql_dev:5432/docker_test1
where pgsql_dev is the host(container name)
docker test1 is the databvse name
5432 is the port
postgres:postgres is username:password

### some of the docker commands ive used :
```bash
docker exec -it pgsql psql -U postgres -l
docker build -f Dockerfile.golang . # without name and tag
docker build -t godocker:v3 -f ./docker_images/Dockerfile.golang .# with name and tag
docker-compose -f docker-compose-dev.yml up -d 
# the latest docker copmpose syntax doesnt have -
# docker compose -f .....
```
#### to setup dev environment
https://github.com/air-verse/air
```bash 
go install github.com/air-verse/air@latest
air init
```
#### to run a debugger
```
make compose-up-debug-build
```
/docker-entrypoint-initdb.d is a special directory in the official PostgreSQL Docker image. When the container starts for the first time, it automatically:

Executes all .sql files in alphabetical order
Executes all .sh shell scripts in alphabetical order
Only runs these files if the database is being initialized for the first time

## TODO:
- [x] set up prod as well as for development
- [x] as part of development set up 
  - [x] hot reloading
  - [x] go debugger
- [ ] manage secrets (pull them out of .env) try using a vault
