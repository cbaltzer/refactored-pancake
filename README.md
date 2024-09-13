# refactored-pancake

# Build and run 

The quickest way to get up and running is simply:

```
go mod download
go test risks-api/storage
go run main.go
```

Or, you can run in a container:

```
docker build -t risks-api . 
docker run -dp 8080:8080 risks-api 
```

