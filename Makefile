# Variables
APP_NAME = jurisdictio_app
DOCKER_MYSQL_NAME = mysql8
MYSQL_IMAGE = mysql:5.7
MYSQL_PORT = 3306
MYSQL_USER = root
MYSQL_PASSWORD = root
MYSQL_DATABASE = database_name

# Docker commands
.PHONY: build
build:
	docker build -t $(APP_NAME) .

.PHONY: run
run: build
	docker run --name $(APP_NAME) --rm -p 8080:8080 --link $(DOCKER_MYSQL_NAME) \
	-e DOCKER_DB_DSN='$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(DOCKER_MYSQL_NAME):$(MYSQL_PORT))/$(MYSQL_DATABASE)?charset=utf8&parseTime=True&loc=Local' \
	-e LOCAL_DB_DSN='$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)?charset=utf8&parseTime=True&loc=Local' \
	-e JWT_SECRET='GOLANGISAWESOME' \
	$(APP_NAME)

.PHONY: mysql
mysql:
	docker run --name $(DOCKER_MYSQL_NAME) -p $(MYSQL_PORT):$(MYSQL_PORT) \
	-e MYSQL_ROOT_PASSWORD=$(MYSQL_PASSWORD) \
	-e MYSQL_DATABASE=$(MYSQL_DATABASE) \
	-d $(MYSQL_IMAGE)

.PHONY: stop-mysql
stop-mysql:
	docker stop $(DOCKER_MYSQL_NAME) && docker rm $(DOCKER_MYSQL_NAME)

.PHONY: clean
clean:
	docker system prune -f

# Helper commands
.PHONY: migrate
migrate:
	docker exec -it $(DOCKER_MYSQL_NAME) mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) -e 'CREATE DATABASE IF NOT EXISTS $(MYSQL_DATABASE);'

.PHONY: logs
logs:
	docker logs $(DOCKER_MYSQL_NAME)

.PHONY: mysql-cli
mysql-cli:
	docker exec -it $(DOCKER_MYSQL_NAME) mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD)

# Default target
.PHONY: all
all: mysql run
