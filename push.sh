#!/bin/bash

echo "🐅 Тигр настраивает GitHub и пушит проект..."

# Проверяем, инициализирован ли Git
if [ ! -d ".git" ]; then
    echo "🔧 Инициализирую Git..."
    git init
    git branch -M main
else
    echo "✅ Git уже инициализирован."
fi

# Спрашиваем ссылку на репозиторий
echo "Вставь ссылку на ТВОЙ GitHub репозиторий (например: https://github.com/Tiger/tiger-casino.git):"
read repo_url

# Проверяем, есть ли уже remote origin
if git remote get-url origin &> /dev/null; then
    echo "🔧 Remote origin уже существует. Обновляю..."
    git remote set-url origin "$repo_url"
else
    echo "🔧 Добавляю remote origin..."
    git remote add origin "$repo_url"
fi

# Создаём .gitignore если его нет
if [ ! -f ".gitignore" ]; then
    echo "🔧 Создаю .gitignore..."
    cat > .gitignore << 'EOF'
# База данных SQLite
*.db

# Временные файлы
*.tmp
*.log

# Папка с зависимостями (если не хочешь пушить вендор)
# vendor/

# Файлы IDE
.vscode/
.idea/
EOF
    echo "✅ .gitignore создан!"
fi

# Добавляем все изменения
git add .

# Спрашиваем сообщение для коммита
echo "Введи сообщение для коммита:"
read commit_message

# Если ничего не ввёл — коммит по умолчанию
if [ -z "$commit_message" ]; then
    commit_message="Первый коммит 🐅"
fi

# Коммитим
git commit -m "$commit_message"

# Определяем текущую ветку
BRANCH=$(git branch --show-current)
git push -u origin "$BRANCH"

echo "✅ Проект запушен на GitHub!"
echo "🔗 Проверяй: $repo_url"