# Elections Blockchain 🗳️ ⛓️

A simple Blockchain implementation in Golang for an Election vote casting application

## Rationale

This project was started due to the scarcity of authenticity and transparency of the elections held in the Republic of Sudan in the past. It aims to break those barriers by implementing a blockchain solution. By using Golang and Docker, the project is able to handle high traffic load and performance bottlenecks other languages suffer from while working on any platform with a small disk footprint (~15MB is the final size of the image).

E-voting systems provide numerous benefits which include: 
- Prevention of fraud, by reducing human involvement 
- Acceleration of results processing 
- Reduction of spoilt ballots by improved presentation and automatic validation of ballots 
- Reduction of costs due to decreasing voting overhead 
- Increase involvement in democratic process due to easer availability (remote voting) 
- Potential for more direct democracy.

> _Quoted from Towards the intelligent agents for blockchain e-voting systems_

## Features

- [x] Block and Blockchain validation
- [x] Vote duplication check
- [x] SHA256 hashing (with padding)
- [ ] Public-Private key-pair generation (sent via email)
- [x] Optimized Docker Image 👏
- [ ] Decetralization
- [ ] View elections results

## Requirements (development)

- Go >= 1.12.5
- Dep >= 0.5.4
- Docker >= 19.03.1

## Environment variables

```bash
PORT=8080
```

## Bootstrap

Development:

1. Populate vendor directory with dependencies ```dep ensure --vendor-only```
2. Serve on port 8080 (env) ```go run main.go```

Production:

1. Build the optimized image ```docker build -t elections-bc .```
2. Expose image on port 8080 ```docker run -d -p 8080:8080 elections-bc```

Testing:

- Testing with coverage ```go test ./... -v -cover```

## License

MIT &copy; Osama Adil
