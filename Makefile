EXECUTABLE_FILE = dns-record-updater

.PHONY: build
build:
	go build -o ${EXECUTABLE_FILE}

.PHONY: clean
clean:
	rm ${EXECUTABLE_FILE}
