# PrimAPI - Primate Facts API

A RESTful API service that provides interesting facts about various primate species. Built with Go and PostgreSQL, with Python integration for data processing.

## Features

- RESTful API endpoints for retrieving primate facts
- Random fact retrieval across all species
- Species-specific fact retrieval
- Docker containerization for easy deployment
- Automatic database population on startup
- PostgreSQL database for data storage

## API Endpoints

### Get a Random Fact
```
GET /fact
```
Returns a random fact about any primate species.

Example response:
```json
{
    "id": 92,
    "fact": "Gorilla infants stay with their mothers for several years, learning essential survival skills and developing strong bonds.",
    "species": "gorilla"
}
```

### Get a Species-Specific Fact
```
GET /fact/{species}
```
Returns a random fact about the specified primate species.

Example response:
```json
{
    "id": 45,
    "fact": "The golden monkey's fur is a vibrant golden-orange, which provides camouflage in the dappled sunlight of the forest canopy.",
    "species": "golden_monkey"
}
```

### Supported Species

The API currently supports facts about the following primate species:

- Aye-aye
- Baboon
- Capuchin
- Chimpanzee
- Gibbon
- Golden Monkey
- Gorilla
- Howler Monkey
- Proboscis Monkey
- Tarsier

## Prerequisites

- Docker
- Docker Compose
- Go 1.24 (for local development)
- Python 3.x (for local development)

## Setup and Installation

1. Clone the repository:
```bash
git clone github.com/ATLIOD/PrimAPI
cd primAPI
```

2. Build and run the containers:
```bash
docker compose up --build
```

The API will be available at `http://localhost:8080`

## Project Structure

```
primAPI/
├── Dockerfile              # Go application container configuration
├── docker-compose.yml      # Multi-container Docker configuration
├── main.go                # Main Go application entry point
├── read_facts.py          # Python script for database population
├── facts/                 # Directory containing primate facts data
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
└── requirements.txt       # Python dependencies
```