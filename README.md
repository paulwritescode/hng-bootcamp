# JSON API Server

A simple API server that returns a JSON response with a dynamically generated current datetime in ISO 8601 format (UTC), along with some predefined information about a user.

## Project Description

This project is a Go-based HTTP server that handles GET requests and returns JSON responses. The JSON response includes the user's email, the current datetime in ISO 8601 format (UTC), and the user's GitHub URL.

## Setup Instructions

### Prerequisites

- Go installed on your machine (version 1.16 or higher).
- An internet connection to download dependencies.

### Running the Project Locally

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/json-api-server.git
   cd json-api-server
   ```

2. **Build the project:**

   ```sh
   go build
   ```

3. **Run the server:**
   ```sh
   ./json-api-server
   ```

The server will start on port 3000.

## API Documentation

### Endpoint URL

`GET /`

### Request

No request parameters are required.

### Response

The response is in JSON format and includes the following fields:

- `email`: The user's email.
- `current_datetime`: The current datetime in ISO 8601 format (UTC).
- `github_url`: The user's GitHub URL.

### Example Response

```json
{
  "email": "kinyattipaul@gmail.com",
  "current_datetime": "2025-01-31T20:30:00Z",
  "github_url": "https://github.com/paulwritescode/hng-bootcamp"
}
```

