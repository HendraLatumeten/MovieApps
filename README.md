
# BACKEND MOVIE

Xsis Test Movie RESTFul API


## Tech Stack

**Golang:** programming language


**Gorilla/mux:** handle http request

**PostgreSql:** DBMS



## Installation Steps

1. Clone Repository
```bash
  git clone https://github.com/HendraLatumeten/CakeStore.git
```
2. Install Dependencies
```bash
  go get -u ./...
```
3. Create Database mysql
```bash
 CREATE DATABASE moviedb;
```
4. Create Table Name
```bash
CREATE TABLE movies (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100),
  description text,
  rating float8,
  image VARCHAR(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```
5. Run App
```bash
 go run .
```





## API Reference

#### 1. List of Movie

```
  GET /movies
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `-` | `string` | return the details of a movies in JSON format|

#### 2. Detail Of Movie

```
  GET /movies/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | return the details of a movies in JSON format|

#### 3. Add New Movie

```
  POST /movies
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `-` | `string` | Add a Movie to the Movies list, the data will be sent as a JSON in the request body |

#### 4. Update Movie

```
  PATCH /movie/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | Add a Movie to the Movies list, the data will be sent as a JSON in the request body |

 
 #### 5. Delete Movie

```
  DELETE /movies/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | Delete Movie from Database|

## Connect with me

 - [Facebook](https://web.facebook.com/hendra.latumeten)
 
 - [Instagram](https://www.instagram.com/hendralatumeten)
 
 - [Linkedin](https://www.linkedin.com/in/hendralatumeten/)
 
