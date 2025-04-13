.PHONY: build run stop

# Сборка для main-service
build-main-service:
	(cd backend/main-service && go build -o main-service ./cmd)

# Сборка для auth-service
build-auth-service:
	(cd backend/auth-service && go build -o auth-service ./cmd)

# Общая сборка всех сервисов
build: build-main-service build-auth-service

# Запуск main-service
run-main-service:
	(cd backend/main-service && go run ./cmd/main.go &)

# Запуск auth-service
run-auth-service:
	(cd backend/auth-service && go run ./cmd/main.go &)

# Запуск всех сервисов
run: build
	$(MAKE) run-main-service
	$(MAKE) run-auth-service

# Остановка сервисов
stop:
	@echo "Останавливаем сервисы..."
	@kill -9 $(shell lsof -t -i :8080)
	@kill -9 $(shell lsof -t -i :8081)
	@echo "Сервисы остановлены"
