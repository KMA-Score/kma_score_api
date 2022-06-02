
# KMA Score API

KMA Score API written in Go


## Installation

- [Install Golang (1.18.3 or above)](https://go.dev/doc/install)

- Copy a dump SQLite3 Score Database to folder. Leave it near main.go

    *Tip: You can dump it by using [KMA Score Extractor](https://github.com/Haven-Code/KMA-Score-Extractor)*

- Create/Edit enviroment file

```env
# .env

PORT = 8000 # Change port here
DB_PATH = ./kma_score.db # Path to DB file
```

- Run locally

```bash
  go run main.go
```

- Build

```bash
  go build -o /kma-score-api
```
    
## API Reference

#### Get all scores by student code

```http
  GET /scores/{studentCode}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `studentCode` | `string` | **Required** |

#### Get all subject

```http
  GET /subjects
```

#### Get Avg score

```http
  GET /avg-score/{studentCode}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `studentCode` | `string` | **Required** |

#### Edit score in database

```http
  POST /add-score/{studentCode}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `studentCode` | `string` | **Required** |

