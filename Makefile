all: clean build_app run_app
build_app:
	go build -ldflags "-s -w" -o bin/facepayserver
run_app:
	export CONFIG_LOCATION=./config.d/setting.conf && export LOG_DIR_LOCATION=./var/log/facepayserver && ./bin/facepayserver
clean:
	rm -rf bin
	rm -rf var