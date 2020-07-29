test:
	./devops/scripts/tests/run-mongo.sh
	go test -v ./...
	./devops/scripts/tests/stop-mongo.sh
