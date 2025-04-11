# A2SV Hub API

This is a RESTful API for A2SV Hub built with Go following Clean Architecture principles and using the Gin framework.

## Project Structure

The project follows Clean Architecture with the following layers:

- **Domain**: Contains the business entities and repository interfaces
  - `entity`: Business entities
  - `repository`: Repository interfaces
  - `services`: Domain services
  
- **usecases**: Contains the application business rules
  
- **Repository**: Contains the repository implementations
  - `postgres`: PostgreSQL implementations
  
- **Delivery**: Contains the delivery mechanisms
  - `http`: HTTP API setup with router and handlers
  - `http/handlers`: HTTP API handlers
  
- **Infrastructure**: Contains the frameworks and drivers
  - Database connections
  - External services

## Prerequisites

- Go 1.16+ (for local development)
- PostgreSQL
- Docker and Docker Compose (for containerized deployment)

## Environment Variables

Create a `.env` file in the root directory using the provided `.env.example` as a template:

```
PORT=8080
DATABASE_URL=postgresql://username:password@host:port/database?sslmode=require
```

> **IMPORTANT**: The `DATABASE_URL` environment variable is required for the application to connect to your PostgreSQL database. Never hardcode database credentials in your source code. Always use environment variables for sensitive information.

For Neon PostgreSQL, your connection string should follow this format:
```
DATABASE_URL=postgresql://username:password@hostname-pooler.region.aws.neon.tech/database_name?sslmode=require
```

## Running Locally

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

3. Create your `.env` file by copying `.env.example` and filling in your actual values:
```bash
cp .env.example .env
# Then edit .env with your actual credentials
```

4. Run the application:

```bash
go run main.go
```

## Running with Docker

### Using Docker Compose (Recommended)

1. Create your `.env` file as described above.

2. Build and start the container:

```bash
docker-compose up -d
```

3. To stop the container:

```bash
docker-compose down
```

### Using Docker Directly

1. Build the Docker image:

```bash
docker build -t a2sv-hub-api .
```

2. Run the container:

```bash
docker run -p 8000:8080 --env-file .env -d a2sv-hub-api
```

## API Endpoints

### Users

- **POST /api/users**: Create a new user
- **GET /api/users**: List all users
- **GET /api/users/:id**: Get a user by ID
- **PUT /api/users/:id**: Update a user
- **DELETE /api/users/:id**: Delete a user

## Example Requests in Postman

### Create User
- Method: `POST`
- URL: `http://localhost:8000/api/users`
- Headers: `Content-Type: application/json`
- Body (raw JSON):
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```

### Get All Users
- Method: `GET`
- URL: `http://localhost:8000/api/users`

### Get User by ID
- Method: `GET`
- URL: `http://localhost:8000/api/users/1`

### Update User
- Method: `PUT`
- URL: `http://localhost:8000/api/users/1`
- Headers: `Content-Type: application/json`
- Body (raw JSON):
  ```json
  {
    "name": "Updated Name",
    "email": "john@example.com",
    "password": "newpassword"
  }
  ```

### Delete User
- Method: `DELETE`
- URL: `http://localhost:8000/api/users/1`

## License

This project is licensed under the MIT License.
