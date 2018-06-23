.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test:
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o gnc .

serve: build
	./gnc

clean:
	rm ./gnc

pack:
	GOOS=linux make build
	docker build -t gcr.io/gamenightcrewicu/gnc-site:$(TAG) .

upload:
	docker tag gcr.io/gamenightcrewicu/gnc-site:$(TAG) gcr.io/gamenightcrewicu/gnc-site:$(TAG)
	docker push gcr.io/gamenightcrewicu/gnc-site:$(TAG)

deploy:
	envsubst < ../kube/deployment.yml | kubectl apply -f -

ship: test pack upload deploy clean