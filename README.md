# Blog REST API

## Run Develeopment

For runing development please make sure that nodejs file watcher `nodemon` is installed

`make run-dev`

## Run Production

`make run`

## Build

`make build`

## Article Endpoints

### List

`curl -X GET http://localhost:5000/artilce/`

### Create

`curl -X POST http://localhost:5000/article/ -d '{"title":"article 1", "description":"article 1 description"}'`

### Update

`curl -X PUT http://localhost:5000/article/1 -d '{"title":"article 1 change", "description":"article 1 description change"}'`

### Delete

`curl -X PUT http://localhost:5000/article/1`

## Todos

1. Refactor Storage code to merge in repository layer
1. Add structured logging
1. update server for gracefull shutdown
1. add code for enviornment variable reading
1. use context in all layers to cancel request
1. add unit tests for all functions in all layers
