#
# Development
OUTD = out
TARGET = $(OUTD)/ggbook


all: $(TARGET)

clean:
	rm -fr $(OUTD)

test:
	go test ./...

build:
	go build -o $(TARGET)

run: $(TARGET)
	./$(TARGET)

#
# Manual testing section

# Example JSON data for PUT and POST requests
JSON_DATA = '{ \
    "name": "Test Guest", \
    "email": "test@example.com"\
}'

API_URL = http://localhost:8080/guests

# list of handy curl command to call our web server locally
curl_post:
	curl -X POST -d $(JSON_DATA) $(API_URL)

curl_get:
	curl $(API_URL)/123

curl_get_all:
	curl $(API_URL)


curl_delete:
	curl -X DELETE $(API_URL)/123