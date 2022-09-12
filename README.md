
# KMA Score API

KMA Score API written in Go


## Installation
There are two ways you can use: 
### 1. Compile and run by yourself

1. [Install Golang (1.18.3 or above)](https://go.dev/doc/install)

2. [Install MariaDB (10.6.4 or above)](https://mariadb.org/download/)

    * You can [install MySQL](https://dev.mysql.com/downloads/mysql/) instead of MariaDB
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
docker pull arahiko/kma-score-api:latest
```
2. Run
```shell
docker run -p 8080:8080 --name kma_score \
-e PORT=8080 DB_USERNAME=username DB_PASSWORD=password \
DB_NAME=database_name DB_HOST=localhost DB_PORT=3306 \
MEILISEARCH_HOST=your-host MEILISEARCH_PORT=7700 \
MEILISEARCH_API_KEY=meilisearch_api_key \
arahiko/kma-score-api:lastest
```
* Or you can use this env file when run docker container
```shell
docker run -p 8080:8080 --name kma_score --env-file path/to/.env arahiko/kma-score-api:latest 
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

#### Search students

```http
GET /student/?query={query}
```

| Parameter | Type     | Description                |
|:----------| :------- | :------------------------- |
| `query`   | `string` | **Required** |

