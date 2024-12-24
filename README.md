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