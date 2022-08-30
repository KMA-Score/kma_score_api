
# KMA Score API

KMA Score API written in Go


## Installation
There are two ways you can use: 
### 1. Compile and run by yourself

1. [Install Golang (1.18.3 or above)](https://go.dev/doc/install)

2. Copy a dump SQLite3 Score Database to folder. Leave it near main.go

    *Tip: You can dump it by using [KMA Score Extractor](https://github.com/Haven-Code/KMA-Score-Extractor)*


3. Create/Edit enviroment file

```env
# .env

PORT = 8000 # Change port here
```

4. Run locally

```shell
go run main.go
```

- Or you can build

```shell
go build -o /kma-score-api
```

### 2. Using our [Docker image](https://hub.docker.com/r/arahiko/kma-score-api)
1. Pull Docker image
```shell
docker push arahiko/kma-score-api:latest
```
2. Run
```shell
docker run -p 8080:8080 --name kma_score -e DB_PATH <username>:<password>@tcp(<host>:<port>)/<dbname>
```

## API Reference

#### Get all scores by student code

```http
GET /student/{StudentId}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `StudentId` | `string` | **Required** |

#### Get all subject

```http
GET /subjects
```

#### Edit score in database

```http
POST /add-score/{StudentId}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `StudentId` | `string` | **Required** |

