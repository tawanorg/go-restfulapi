APP_NAME					= simple-go-restfulapi
APP_BUILD_SRC			= build/

install:
	@yarn install
	$(info Install packages..)

clean:
	@rm -rf ${APP_BUILD_SRC}/*

build: deps
	cd src && go ${APP_BUILD_SRC} -o ../${APP_BUILD_SRC}/${APP_NAME} *.go
	$(info Building packages..)

start:
	./${APP_BUILD_SRC}/${APP_NAME}
	$(info Starting application..)

dev:
	cd src && go run *.go
