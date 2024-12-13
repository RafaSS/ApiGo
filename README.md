# ApiGo + HTMX + templ

## Description

ApiGo is a Go-based API for managing and displaying manga lists by interacting with the MangaDex API. It allows users to search, view, and list their favorite mangas.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/COMERCIAL/ApiGo.git
   cd ApiGo
   ```

2. **Install dependencies:**

   Ensure you have Go installed (version 1.23.4).

   ```bash
   go mod download
   ```

3. **Set up environment variables:**

   Create a `.env` file in the project root and add the following:

   ```env
   USERNAMEMANGA=your_username
   PASSWORD=your_password
   CLIENT_ID=your_client_id
   CLIENT_SECRET=your_client_secret
   ```

## Usage

1. **Build and run the application:**

   ```bash
   go run cmd/main.go
   ```

2. **Access the API:**

   Open your browser and navigate to `http://localhost:1323` to view the manga list.

## API Endpoints

- `GET /` - Fetch and display the first page of mangas.
- `POST /search` - Search for mangas by title.
- `GET /list?page={number}` - Load more mangas with pagination.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
