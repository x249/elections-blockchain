# Elections Blockchain 🗳️ ⛓️

A simple Blockchain implementation in Golang for an Election vote casting application

## Rationale

This project was started due to the scarcity of authenticity and transparency of the elections held in the Republic of Sudan in the past. It aims to break those barriers by implementing a decentralized solution where anyone can view the content without mutating it once commited. By using Golang and Docker, the project is able to handle high traffic load and performance bottlenecks other languages suffer from while working on any platform with a small disk footprint.

## Features

- [x] Block creation
- [x] Block validation
- [x] Blockchain validity check
- [x] Optimized Docker Image 👏
- [ ] Decetralization
- [ ] List candidates
- [ ] Check if person already voted

## Requirements (development)

- Go >= 1.12.5
- Spew ```go get -u github.com/davecgh/go-spew/spew```
- Mux ```go get -u github.com/gorilla/mux```
- Godotenv ```go get -u github.com/joho/godotenv```

## Bootstrap

```bash
go run main.go # starts server on port 8080

docker build -t elections-bc -f Dockerfile.multistage . # build the optimized image
```

## License

MIT &copy; Osama Adil
