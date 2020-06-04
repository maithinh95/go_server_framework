all: clean build_app run_app
build_app:
	go build -ldflags "-s -w" -o bin/go-server
run_app:
	export CONFIG_LOCATION=./config.d/setting.conf && export LOG_DIR_LOCATION=./var/log/go-server && ./bin/go-server
clean:
	rm -rf bin
	rm -rf var