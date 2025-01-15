#
# Development targets
OUT_DIR = out
TARGET = $(OUT_DIR)/ggbook

GO_TEST_FLAGS = -v -coverprofile=coverage.out


all: $(TARGET)

clean:
	rm -fr $(OUT_DIR)
	rm *.db
	rm coverage.out

test:
	go test $(GO_TEST_FLAGS) ./...

build:
	go build -o $(TARGET)

build_frontend:
	cd html && npm run build

run: build build_frontend
	./$(TARGET)

migrate: build
	./$(TARGET) -migrate

coverage: test
	go tool cover -html=coverage.out -o coverage.html

#
# Manual testing targets

API_URL = http://localhost:8080/guests
JSON_HEADER = -H "Content-Type: application/json"

# list of handy curl command to call our web server locally
curl_post:
	curl -X POST $(JSON_HEADER) --data-binary "@testdata/oneguest.json" $(API_URL)

curl_get:
	curl $(API_URL)/1

curl_get_all:
	curl $(API_URL)
	
curl_get_all_pretty:
	curl $(API_URL) | gojq '.'

curl_delete:
	curl -X DELETE $(API_URL)/1
