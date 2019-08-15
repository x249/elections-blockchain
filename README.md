# Elections Blockchain

A simple Blockchain implementation in Golang for an Election vote casting application

## Rationale

This project was started due to the scarcity of authenticity and transparency of the elections held in the Republic of Sudan in the past. It aims to break those barriers by implementing a decentralized solution where anyone can view the content without mutating it once commited. By using Golng, the project is able to handle high traffic load and performance bottlenecks other languages suffer from.

## Requirements (development)
- Go >= 1.12.5
- Spew ```go get -u github.com/davecgh/go-spew/spew```
- Mux ```go get -u github.com/gorilla/mux```
- Godotenv ```go get -u github.com/joho/godotenv```

## Bootstrap

```bash
$ go run main.go
```

## License
MIT &copy; Osama Adil
