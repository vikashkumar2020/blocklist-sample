# Server
server:
	@echo "Starting server..."
	@go run ./cmd/main.go

# Srcipt
script:
	@echo "Starting script..."
	@go run ./scripts/main.go