# Elections Blockchain 🗳️ ⛓️

[![Build Status](https://travis-ci.org/x249/elections-blockchain.svg?branch=master)](https://travis-ci.org/x249/elections-blockchain)
[![Coverage Status](https://coveralls.io/repos/github/x249/elections-blockchain/badge.svg?branch=master)](https://coveralls.io/github/x249/elections-blockchain?branch=master)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/x249/elections-blockchain/issues)
![GitHub](https://img.shields.io/github/license/x249/elections-blockchain)

A simple Blockchain implementation in Golang for an Election vote casting application

## Rationale

This project was started due to the scarcity of authenticity and transparency of the elections held in the Republic of Sudan in the past. It aims to break those barriers by implementing a blockchain solution. By using Golang and Docker, the project is able to handle high traffic load and performance bottlenecks other languages suffer from while working on any platform with a small disk footprint (~15MB is the final size of the image).

## Benefits

E-voting systems provide numerous benefits which include:

- Prevention of fraud, by reducing human involvement

- Acceleration of results processing

- Reduction of spoilt ballots by improved presentation and automatic validation of ballots

- Reduction of costs due to decreasing voting overhead

- Increase involvement in democratic process due to easer availability (remote voting)

- Potential for more direct democracy.

Introducing an e-voting system is inherently connected with many challenges which are not only technical but also
procedural and legislative. One the most substantial of weaknesses of e-voting systems are:

1. lack of transparency and understanding of the system which leads to lack of trust in the solution and undermines its whole sense
2. lack of widely accepted standards for e-voting systems
3. risk of fraud and manipulation by privileged insiders or hackers
4. increased costs of voting infrastructure with regard to power supply, communication technology, etc.

Blockchain technology can provide a solution to some of these problems. One of the properties of blockchain is an ability to create a platform for a public verification of data stored inside. Utilizing this characteristic of blockchain technology allows for a possibility of the creation of an e-voting system which can be audited by common voters instead of dedicated institutions and officials.

> _Quoted from Towards the intelligent agents for blockchain e-voting systems - Procedia Computer Science 141 (2018) 239-146_

## Features

- [x] Block and Blockchain validation
- [x] Vote duplication check
- [x] SHA256 hashing (with padding)
- [ ] Public-Private key-pair generation (sent via email)
- [x] Optimized Docker Image 👏
- [ ] Decetralization
- [ ] Peer discovery
- [ ] View elections results
- [x] CI/CD

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
