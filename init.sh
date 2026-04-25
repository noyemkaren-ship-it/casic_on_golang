#!/bin/bash

set -e

PROJECT_NAME="tiger_app"
MAIN_FILE="main.go"
MOD_FILE="go.mod"
GO_MODULE="test_gin"

download() {
    echo "📦 Устанавливаю Gin..."
    go get -u github.com/gin-gonic/gin
    
    echo "📦 Устанавливаю GORM и SQLite..."
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/sqlite
    
    echo "📦 Чищу зависимости..."
    go mod tidy
    
    echo "✅ Зависимости установлены!"
}

if [ ! -f "$MOD_FILE" ]; then
    echo "🔧 go.mod не найден. Инициализирую модуль $GO_MODULE..."
    go mod init $GO_MODULE
else
    echo "✅ go.mod уже существует."
fi

# Проверка main.go
if [ ! -f "$MAIN_FILE" ]; then
    echo "🔧 main.go не найден. Создаю заготовку..."
    cat > $MAIN_FILE << 'EOF'
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8080")
}
EOF
    echo "✅ main.go создан!"
else
    echo "✅ main.go уже существует."
fi

download

echo "🚀 Запускаю проект..."
go run .