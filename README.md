# Run Locally
## Requirement
Install docker and go first, then start mongoDB
```
docker pull mongo:latest
docker run -d -p 27017:27017 mongo
```
## Backend
Under backend folder run this to start server on port 9876
```
go mod tidy
go run main.go
```
## Frontend
Starts the frontend on port 8765
```
python3 -m http.server 8765
```

# REPLIT
frontend:
```
https://shopifyfrontend.dracocococo.repl.co/
```
backend:
```
https://shopify.dracocococo.repl.co/api/inventory
```

