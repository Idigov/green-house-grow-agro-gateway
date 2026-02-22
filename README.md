# green-house-grow-agro-gateway

API Gateway (BFF) для информационной системы управления тепличным комплексом: единая точка входа для Web UI, маршрутизация к внутренним сервисам, аутентификация и агрегация данных.

## Требования

- Go 1.21+

## Запуск

Из корня проекта:

```bash
go run ./cmd/gateway
```

Сервер поднимется на порту из `GATEWAY_PORT` (по умолчанию **8080**).

Проверка: [http://localhost:8080/health](http://localhost:8080/health) — liveness; `/ready` — readiness.

### Разработка с автоперезапуском (Air)

```bash
go install github.com/air-verse/air@latest
air
```

Конфигурация Air — в `.air.toml`.

## Git-хуки (bash)

Перед первым коммитом установите хуки (нужны только bash и Go):

```bash
./scripts/install-hooks.sh
```

Скрипт копирует в `.git/hooks/`:
- **pre-commit** — запрещает коммит на ветку `main` (нужна ветка и PR); перед коммитом запускает `go fmt` и `go vet`. Исключение: коммит на main разрешён с флагом (команду можно скопировать целиком):

```
ALLOW_MAIN_COMMIT=1 git commit -m "ваше сообщение"
```

- **pre-push** — запрещает прямой push в `main` (нужна ветка и PR)

Хуки лежат в `scripts/git-hooks/`; после клона репо каждый выполняет `./scripts/install-hooks.sh`.

## Защита ветки main на GitHub

Чтобы в main нельзя было пушить напрямую и вливать только через PR:

1. Репозиторий на GitHub → **Settings** → **Branches** → **Add branch protection rule**.
2. Branch name pattern: `main`.
3. Включить: **Require a pull request before merging**, при необходимости **Require status checks to pass**.
4. Сохранить. После этого push в `main` и merge без PR будут запрещены на стороне GitHub.
