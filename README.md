# URL Shortener

A simple URL shortening service built in Go.

## Overview

This project provides a basic URL-shortening service implemented in Go. It allows users to shorten long URLs into more manageable and shareable links. The service also includes a redirect feature to redirect users from the shortened URL to the original long URL.

## Usage

### Running the Server

Navigate to the project directory and run the following command to start the server:

```sh
go run main.go
```

The server will start listening on port `3000` by default. You can change the port in the `main.go` file if needed.

### Shortening a URL

To shorten a URL, send a POST request to the `/shorten` endpoint with a JSON payload containing the original URL:

```sh
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
```

![image](https://github.com/user-attachments/assets/ce74a0cc-9e8b-4a7b-9f7f-bff04e93bebb)

The response will contain a JSON object with the shortened URL:

```json
{
    "short_url": "https://www.linkedin.com/in/iamprince/"
}
```

![image](https://github.com/user-attachments/assets/82f5a602-442a-4a41-8c7b-db614560cec7)


### Redirecting to the Original URL

To redirect to the original URL, visit the shortened URL in your browser or send a GET request to the `/redirect/{id}` endpoint, where `{id}` is the shortened URL ID:

```sh
curl http://localhost:3000/redirect/abcdef
```

This will redirect you to the original URL associated with the shortened URL.
