#!/bin/bash

# Project name
PROJECT_NAME="social-media-service"

# Create directory structure
echo "Creating project structure..."

mkdir -p $PROJECT_NAME/{cmd,config,internal/{handler,middleware,model,repository,router,service},pkg/{auth,db},proto,scripts}
touch $PROJECT_NAME/.env
touch $PROJECT_NAME/go.mod
touch $PROJECT_NAME/Dockerfile
touch $PROJECT_NAME/docker-compose.yml

# Create main.go
cat <<EOF > $PROJECT_NAME/cmd/main.go
package main

func main() {
    // Entry point
}
EOF

# Create empty files in each directory
touch $PROJECT_NAME/config/config.go
touch $PROJECT_NAME/internal/handler/user_handler.go
touch $PROJECT_NAME/internal/handler/post_handler.go
touch $PROJECT_NAME/internal/middleware/auth.go
touch $PROJECT_NAME/internal/model/user.go
touch $PROJECT_NAME/internal/model/post.go
touch $PROJECT_NAME/internal/repository/user_repository.go
touch $PROJECT_NAME/internal/service/user_service.go
touch $PROJECT_NAME/internal/router/router.go
touch $PROJECT_NAME/pkg/db/postgres.go
touch $PROJECT_NAME/pkg/auth/jwt.go
touch $PROJECT_NAME/scripts/seed.go
touch $PROJECT_NAME/proto/user.proto

echo "Project structure created successfully under ./$PROJECT_NAME"
