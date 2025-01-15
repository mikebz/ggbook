#
# Development targets
OUTD = out
TARGET = $(OUTD)/ggbook


all: $(TARGET)

clean:
	rm -fr $(OUTD)
	rm *.db

test:
	go test ./...

build:
	go build -o $(TARGET)

build_frontend:
	cd html && npm run build

run: build build_frontend
	./$(TARGET)

migrate: build
	./$(TARGET) -migrate

#
# Manual testing targets

API_URL = http://localhost:8080/guests

# list of handy curl command to call our web server locally
curl_post:
	curl -X POST -H "Content-Type: application/json" --data-binary "@testdata/oneguest.json" $(API_URL)

curl_get:
	curl $(API_URL)/1

curl_get_all:
	curl $(API_URL)

curl_delete:
	curl -X DELETE $(API_URL)/1
