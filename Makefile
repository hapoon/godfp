deps:
	glide install
	
deps-update:
	glide update

test:
	go test --cover
	go vet
	