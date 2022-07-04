To start the Application run 
docker compose up 

Required Installation: Docker

How to query the endpoints with curl
GET: curl http://localhost:8090/list

POST: curl -d '{"key":"some key", "value":"some value"}' -H 'Content-Type: application/json' http://localhost:8090/add
