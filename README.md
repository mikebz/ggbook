# Develoment

## Getting the source code

Originally developed by mikebz@ to test out LLM UX (and learn how to create golang micro services with minimal
dependencies)

```
git clone https://github.com/mikebz/ggbook.git
cd ggbook
```

## Backend Development

### Go dependencies

```bash
go mod tidy
```

### Run database migration:

This needs to happen before the sample can run.

```bash
make migrate
```

### Start the web server:

```bash
make run
```

The server will listen on `localhost:8080` by default. You can change this by setting the `SERVER` and `PORT` environment variables.

## Front end development

Thank you to: https://medium.com/@alen.ajam/building-a-simple-chat-app-with-vue-js-462c4a53c6ad for inspiration of how to build a simple chat in VueJS.

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

### Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

### Running dev frontend

All the front end GUI is in the /html folder and uses VueJS. In order to
compile and run the front end please follow the standard `npm` instructions:

```
  npm install
  npm run format
  npm run dev
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

## API Endpoints

- **`/`**: Returns a welcome message.
- **`/guests`**:
  - `GET`: Retrieves a list of all guests.
  - `POST`: Creates a new guest entry. Requires a JSON payload with `name` and `email` fields.
- **`/guests/{id}`**:
  - `GET`: Retrieves a specific guest by ID.
  - `PUT`: Updates a specific guest by ID. Requires a JSON payload with `name` and `email` fields.
  - `DELETE`: Deletes a specific guest by ID.

## Testing APIs

The project includes a `Makefile` with targets for running tests and making API calls using `curl`.

```bash
make test  # Run unit tests
make curl_post # Example POST request
make curl_get_all # Example GET request for all guests
make curl_get # Example GET request for a specific guest
make curl_delete # Example DELETE request
```
