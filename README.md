# Academy Genius

Платформа для подготовки к олимпиадам по физике.

**Стек:** Nuxt 3 (Vue 3 + TypeScript + Tailwind CSS 4) + Go (Gin + pgx/v5) + PostgreSQL

---

## Установка Go

### Ubuntu / Debian

```bash
# Удалить старую версию (если есть)
sudo rm -rf /usr/local/go

# Скачать Go 1.22+ (проверь актуальную версию на https://go.dev/dl/)
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz

# Распаковать в /usr/local
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz

# Добавить в PATH (добавь в ~/.bashrc или ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Применить
source ~/.bashrc

# Проверить
go version
```

### macOS (Homebrew)

```bash
brew install go
go version
```

### Windows

1. Скачай установщик с https://go.dev/dl/
2. Запусти `.msi` установщик
3. Перезапусти терминал
4. `go version`

---

## Установка PostgreSQL

### Ubuntu 22.04 / 24.04

```bash
# Установить PostgreSQL и contrib
sudo apt update
sudo apt install -y postgresql postgresql-contrib

# Запустить и добавить в автозагрузку
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Создать пользователя и БД
sudo -u postgres psql <<EOF
CREATE USER academy_genius WITH PASSWORD 'academia_genius_123';
CREATE DATABASE academy_genius_db OWNER academy_genius;
GRANT ALL PRIVILEGES ON DATABASE academy_genius_db TO academy_genius;
EOF

# Проверить подключение
psql -h localhost -U academy_genius -d academy_genius_db
```

### macOS (Homebrew)

```bash
brew install postgresql@16
brew services start postgresql@16

# Создать пользователя и БД
createdb academy_genius_db
psql -d academy_genius_db <<EOF
CREATE USER academy_genius WITH PASSWORD 'academia_genius_123';
GRANT ALL PRIVILEGES ON DATABASE academy_genius_db TO academy_genius;
ALTER DATABASE academy_genius_db OWNER TO academy_genius;
EOF
```

### Windows

1. Скачай установщик с https://www.postgresql.org/download/windows/
2. Установи с дефолтными настройками (порт 5432)
3. Запусти pgAdmin или psql:
```sql
CREATE USER academy_genius WITH PASSWORD 'academia_genius_123';
CREATE DATABASE academy_genius_db OWNER academy_genius;
```

---

## Установка golang-migrate

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Проверить: `migrate -version`

---

## Быстрый старт

```bash
# 1. Клонировать репозиторий
git clone https://github.com/4codegit/academia-genius.git
cd academia-genius

# 2. Бэкенд

cd backend
go mod download

# Применить миграции
migrate -path migrations -database "postgres://academy_genius:academia_genius_123@localhost:5432/academy_genius_db?sslmode=disable" up

# Запустить
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=academy_genius
export DB_PASSWORD=academia_genius_123
export DB_NAME=academy_genius_db
export JWT_SECRET=your-super-secret-key-change-in-production
export PORT=8080

go run cmd/server/main.go

# 3. Фронтенд (в другом терминале)

cd ../frontend
npm install
npm run dev
```

**Frontend:** http://localhost:3000
**Backend API:** http://localhost:8080/api/v1

---

## Структура проекта

```
academia-genius/
├── backend/                  # Go API сервер
│   ├── cmd/server/main.go    # Точка входа
│   ├── internal/
│   │   ├── config/           # Конфигурация
│   │   ├── models/           # Модели данных
│   │   ├── repository/       # Слой доступа к данным
│   │   ├── service/          # Бизнес-логика
│   │   └── delivery/http/    # HTTP обработчики и middleware
│   └── migrations/           # SQL миграции
├── frontend/                 # Nuxt 3 приложение
│   ├── pages/                # 7 страниц
│   ├── components/           # Vue компоненты
│   ├── composables/          # Vue composables
│   ├── stores/               # Pinia stores
│   ├── middleware/           # Nuxt middleware
│   └── types/                # TypeScript интерфейсы
└── .gitignore
```

---

## Страницы

| Страница | Описание |
|---|---|
| Home | Новости платформы |
| Problems | Задачи по 8 темам физики с фильтрацией |
| Courses | Доступные курсы |
| Books | Книги по категориям (5 табов) |
| Alumni | Выпускники (Пиров Далер — featured) |
| Schedule | Расписание вебинаров |
| Cabinet | Личный кабинет (защищён, метрики + карта знаний) |
