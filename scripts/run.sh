#!/usr/bin/bash

# Generate injectors
wire gen awesomeProject/config

# Generate swagger
swag init --parseDependency --parseInternal

# Run the server
go run main.go