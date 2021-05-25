# cross parameters
SHELL:=/bin/bash -O extglob

install:
	@echo "Installing dependecies"
	@go get -u github.com/cucumber/godog/cmd/godog
	@go mod tidy
	@npm install cucumber-html-reporter --save-dev

selenium:
	@java -jar resources/selenium-server-standalone-3.141.59.jar

godog:
	@godog feature/*.feature

report:
	@node reporter.js
