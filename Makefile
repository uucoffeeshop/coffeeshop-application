include .env
export

run: run-product run-counter run-barista run-kitchen run-proxy run-web

run-product:
	cd cmd/product && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run github.com/uucoffeeshop/coffeeshop-application/cmd/product
.PHONY: run-product

run-counter:
	cd cmd/counter && go mod tidy && go mod download && \
	CGO_ENABLED=0 IN_DOCKER=false go run -tags migrate github.com/uucoffeeshop/coffeeshop-application/cmd/counter
.PHONY: run-counter

run-barista:
	cd cmd/barista && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags migrate github.com/uucoffeeshop/coffeeshop-application/cmd/barista
.PHONY: run-barista

run-kitchen:
	cd cmd/kitchen && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags migrate github.com/uucoffeeshop/coffeeshop-application/cmd/kitchen
.PHONY: run-kitchen

run-proxy:
	cd cmd/proxy && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags migrate github.com/uucoffeeshop/coffeeshop-application/cmd/proxy
.PHONY: run-proxy

run-web:
	cd cmd/web && go mod tidy && go mod download && \
	CGO_ENABLED=0 REVERSE_PROXY_URL=5000 WEB_PORT=8888 go run github.com/uucoffeeshop/coffeeshop-application/cmd/web
.PHONY: run-web

wire:
	cd internal/barista/app && wire && cd - && \
	cd internal/counter/app && wire && cd - && \
	cd internal/kitchen/app && wire && cd - && \
	cd internal/product/app && wire && cd -
.PHONY: wire

sqlc:
	sqlc generate
.PHONY: sqlc

clean:
	go clean
