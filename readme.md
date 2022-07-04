### Required installation: Docker 

### To start the Application ###
docker compose up 

### Querying endpoints with curl #####

> 1. GET: curl http://localhost:8090/list

> 2. POST: curl -d '{"key":"some key", "value":"some value"}' -H 'Content-Type: application/json' http://localhost:8090/add
