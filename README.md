1. **Original source code**

Originally developed by mikebz@ to test out LLM UX (and learn how to create golang micro services with minimal
dependencies)

```
git clone https://github.com/mikebz/ggbook.git
cd ggbook
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Run database migration:**

```bash
go run main.go -migrate
```

4. **Start the web server:**

```bash
go run main.go
```

or build and run:

```bash
make build
make run
```

The server will listen on `localhost:8080` by default.  You can change this by setting the `SERVER` and `PORT` environment variables.

## API Endpoints

* **`/`**:  Returns a welcome message.
* **`/guests`**:
    * `GET`: Retrieves a list of all guests.
    * `POST`: Creates a new guest entry.  Requires a JSON payload with `name` and `email` fields.
* **`/guests/{id}`**:
    * `GET`: Retrieves a specific guest by ID.
    * `PUT`: Updates a specific guest by ID.  Requires a JSON payload with `name` and `email` fields.
    * `DELETE`: Deletes a specific guest by ID.

## Testing

The project includes a `Makefile` with targets for running tests and making API calls using `curl`.

```bash
make test  # Run unit tests
make curl_post # Example POST request
make curl_get_all # Example GET request for all guests
make curl_get # Example GET request for a specific guest
make curl_delete # Example DELETE request
