# Assignment

A simple CRUD (Create, Read, Update, Delete) web application for table movies

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Endpoints](#endpoints)
- [Usage](#usage)

## Features

- **Create**: Add new item to the database.
- **Read**: View existing item(s) from the database.
- **Update**: Modify existing item in the database.
- **Delete**: Remove item from the database.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/flyingfrisbee/assignment.git

2. **Run the app with docker compose if you have docker installed:**

   ```bash
   docker compose up -d

3. **Or if you have golang installed:**

   ```bash
   go run .

## Endpoints

| Method | Endpoint         | Body                | Response          | Description                    |
|--------|------------------|---------------------|-------------------|--------------------------------|
| GET    | {host}:8080/movies          | null                | `[{"id":1,"title":"Pengabdi 2 Comunion","description":"Adalah sebuah film horor Indonesia tahun 2022","rating":7.0,"image":"https://myimage.com","created_at":"2022-08-0110:56:31","updated_at": "2022-08-13 09:30:23"},{"id":2,"title":"Pengabdi 1 Comunion","description":"Adalah sebuah film horor Indonesia tahun 2022","rating":7.0,"image":"https://myimage.com","created_at":"2022-08-0110:56:31","updated_at": "2022-08-13 09:30:23"}]`   | Get all movies                 |
| GET    | {host}:8080/movies/:id      | null                | `{"id":1,"title":"Pengabdi 2 Comunion","description":"Adalah sebuah film horor Indonesia tahun 2022","rating":7.0,"image":"https://myimage.com","created_at":"2022-08-0110:56:31","updated_at": "2022-08-13 09:30:23"}`        | Get a movie by ID              |
| POST   | {host}:8080/movies          | `{"title":"Pengabdi","description":"Sebuah deskripsi","rating":4.7,"image":"https://someimage.com"}` | `{"id":1,"title":"Pengabdi","description":"Sebuah deskripsi","rating":4.7,"image":"https://someimage.com","created_at":"2022-08-0110:56:31","updated_at": "2022-08-13 09:30:23"}`        | Create a new movie             |
| PATCH  | {host}:8080/movies/:id      | `{"title":"Pengabdi","description":"Sebuah deskripsi","rating":4.7,"image":"https://someimage.com"}`           | `{"id":1,"title":"Pengabdi","description":"Sebuah deskripsi","rating":4.7,"image":"https://someimage.com","created_at":"2022-08-0110:56:31","updated_at": "2022-08-13 09:30:23"}`        | Update a movie by ID           |
| DELETE | {host}:8080/movies/:id      | null                | null        | Delete a movie by ID           |

## Usage

1. **Get all movies**: Send a GET request to `/movies` endpoint.
2. **Get a movie by ID**: Send a GET request to `/movies/:id` endpoint, where `:id` is the ID of the movie.
3. **Create a new movie**: Send a POST request to `/movies` endpoint with the required JSON structure in the request body.
4. **Update a movie by ID**: Send a PATCH request to `/movies/:id` endpoint with the updated JSON in the request body, where `:id` is the ID of the movie.
5. **Delete a movie by ID**: Send a DELETE request to `/movies/:id` endpoint, where `:id` is the ID of the movie.

