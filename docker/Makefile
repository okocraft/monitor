build-image:
	docker build -t $(IMAGE_NAME):latest --target $(TARGET) $(DIR)

build-http-server:
	$(MAKE) build-image IMAGE_NAME=monitor_http_server TARGET=http-server DIR=../app/

build-cleanup-batch:
	$(MAKE) build-image IMAGE_NAME=monitor_cleanup_batch TARGET=cleanup-batch DIR=../app/

build-setup-cmd:
	$(MAKE) build-image IMAGE_NAME=monitor_setup_cmd TARGET=setup-cmd DIR=../app/

run-service:
	docker compose up -d --no-deps $(SERVICE_NAME)

run-db:
	$(MAKE) run-service SERVICE_NAME=monitor_db

run-db-migration:
	$(MAKE) run-service SERVICE_NAME=monitor_db_migration

run-http-server: build-http-server
	$(MAKE) run-service SERVICE_NAME=monitor_http_server

run-cleanup-batch: build-cleanup-batch
	$(MAKE) run-service SERVICE_NAME=monitor_cleanup_batch

run-setup-cmd: build-setup-cmd
	$(MAKE) run-service SERVICE_NAME=monitor_setup_cmd
