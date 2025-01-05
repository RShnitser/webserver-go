# API for Chirpy

## user resource

```json
{
  "id": "50746277-23c6-4d85-a890-564c0044c2fb",
  "created_at": "2021-07-07T00:00:00Z",
  "updated_at": "2021-07-07T00:00:00Z",
  "email": "user@example.com",
  "password": "12345",
  "is_chirpy_red": false
}
```

### POST /api/users

Request Body:

```json
{
  "email": "user@example.com",
  "password": "12345"
}
```

Response Body:

```json
{
  "id": "50746277-23c6-4d85-a890-564c0044c2fb",
  "created_at": "2021-07-07T00:00:00Z",
  "updated_at": "2021-07-07T00:00:00Z",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

### PUT /api/users

Request Body:

```json
{
  "email": "user@example.com",
  "password": "12345"
}
```

Response Body:

```json
{
  "id": "50746277-23c6-4d85-a890-564c0044c2fb",
  "created_at": "2021-07-07T00:00:00Z",
  "updated_at": "2021-07-07T00:00:00Z",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

## chirpy resource

```json
{
  "id": "94b7e44c-3604-42e3-bef7-ebfcc3efff8f",
  "created_at": "2021-01-01T00:00:00Z",
  "updated_at": "2021-01-01T00:00:00Z",
  "body": "Hello, world!",
  "user_id": "123e4567-e89b-12d3-a456-426614174000"
}
```

### POST /api/chirps

Request Body:

```json
{
  "body": "Hello, world!",
}
```

Response Body:

```json
{
  "id": "94b7e44c-3604-42e3-bef7-ebfcc3efff8f",
  "created_at": "2021-01-01T00:00:00Z",
  "updated_at": "2021-01-01T00:00:00Z",
  "body": "Hello, world!",
  "user_id": "123e4567-e89b-12d3-a456-426614174000"
}
```

### GET /api/chirps

Response Body:

```json
[
  {
    "id": "94b7e44c-3604-42e3-bef7-ebfcc3efff8f",
    "created_at": "2021-01-01T00:00:00Z",
    "updated_at": "2021-01-01T00:00:00Z",
    "body": "Yo fam this feast is lit ong",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  },
  {
    "id": "f0f87ec2-a8b5-48cc-b66a-a85ce7c7b862",
    "created_at": "2022-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z",
    "body": "What's good king?",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  }
]
```

### GET /api/chirps/{chirpID}

Response Body:

```json
 {
    "id": "94b7e44c-3604-42e3-bef7-ebfcc3efff8f",
    "created_at": "2021-01-01T00:00:00Z",
    "updated_at": "2021-01-01T00:00:00Z",
    "body": "Yo fam this feast is lit ong",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  }
```