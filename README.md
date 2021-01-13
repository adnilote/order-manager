# Order Manager

## Functionality
- orders aggregation from kafka
- stream orders status by grpc
- positions aggregation by account and client tag

## Project Structure
The project stucture was inspired by [DDD](https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5) and [clean architecture](https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f).
```
.
└── app
   │  
   ├── business
   │  ├── entities
   │  │  - business entities
   │  │  └── proto
   │  │     - protobuf files
   │  └── rules
   │     - business rules
   ├── store
   │   - module for intraction with dbs (sync db with cache)
   ├── usecases
   │   - business logic
   └── api
       - grpc api
```

## Run service
Dependency: kafka:9092

```bash
make run
```

### Run linter 
```
make lint
```

### Run tests
```
make test
```
