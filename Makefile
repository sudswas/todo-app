GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=main
GOPATH ?=/root
export GOPATH

all: build kube-deploy

build:
	$(GOGET) -u github.com/gorilla/mux
	$(GOGET) -u github.com/go-sql-driver/mysql
	$(GOBUILD) -o $(BINARY_NAME) myapp/web
	docker build -t myapp .

kube-deploy:
	cd kubenetes/
	kubectl create -f todo-deployment.yaml
	kubectl create -f todo-svc.yaml
