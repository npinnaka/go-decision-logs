```
docker-compose up --build --remove-orphans

curl -XPOST -d '{ "input": { "user": "john", "password": "secret" } }' http://localhost:8181/v1/data/foo/allow

docker-compose down
```