
# How to run
default id range: 1-100
default worker pool: 40
default logger is off

- run with default parameter
```
go run main.go players
```
- run with logger
```
go run main.go players -l=true
```
- run with custom id range (1-200)
```
go run main.go players -l=true -m=200
```
- run with custom id range (1-200) and worker pool (10)
```
go run main.go players -l=true -m=200 -w=10
```

## Modules
- **model**
model holds all models of the application.
- **store**
store handles all kinds of storage. for example, to support mysql, you have add mysql driver
- **resp**
resp responsibles to generate player output
- **task**
task is working like an engine. it fetches data using "store.Team" and put data to given "resp".
