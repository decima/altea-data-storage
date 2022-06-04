run: Altea
	# running
	./Altea

Altea:deps
	# building
	go build Altea.go

deps:clean
	# deps
	go mod download

clean:
	# clean
	go mod download
	go mod tidy
	rm -f Altea


dev:
	go run .


