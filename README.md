# Практическая работа 10

JWT и авторизация пользователей.

Сервис выдаёт access JWT на 15 минут и refresh JWT на 7 дней. Проверяются HS256, подпись, issuer, audience и обязательный exp. Различаются типы access и refresh. Одноразовая ротация refresh защищена mutex; использованные jti хранятся до exp. RBAC ограничивает admin/stats, ABAC разрешает user читать только собственный профиль.

## Запуск на macOS или Linux

Go 1.27.1, curl. Каждая практика — отдельный Go-модуль. Команды выполняются из этой папки.

```bash
cp .env.example .env
set -a; source .env; set +a
go test ./...
go vet ./...
go run ./cmd/server
```

Файл `.env` читается оболочкой через `source`; автоматически сервер его не загружает. Порт API по умолчанию 8080. При одновременном запуске меняйте HTTP_ADDR, APP_ADDR (ПР9) или APP_PORT (ПР10).

## Проверки

Фактические результаты находятся в `evidence/`. Полные curl-команды и ответы — в `evidence/requests.txt` (для практик с HTTP-проверками). Тесты 15 и 16 записаны в `evidence/tests.txt` и `evidence/integration.txt`.

Учебные пользователи: admin@example.com и user@example.com, пароль secret123. POST `/api/v1/login`, POST `/api/v1/refresh` с `{"refresh":"..."}`, GET `/api/v1/me`, `/api/v1/admin/stats`, `/api/v1/users/{id}`. Заголовок Authorization: Bearer <access>. Blacklist хранится в памяти и сбрасывается при рестарте. Старый access действует до exp; одноразовость относится к refresh.

## Подготовка к защите

Registered claims включают sub, iss, aud, exp и iat. Пользовательские поля role и kind дополняют их. exp ограничивает срок применения украденного токена.

JWT проверяется без серверной сессии. Для немедленного отзыва требуется дополнительное состояние. Здесь blacklist нужен для одноразового refresh.

AuthN проверяет личность и кладёт Claims в контекст. AuthZ затем проверяет разрешения. При обратном порядке контекст ещё пуст.

RBAC сравнивает роль admin/user; ABAC дополнительно сравнивает запрошенный id с sub. Нет токена — 401, прав недостаточно — 403.

bcrypt использует соль и cost. Pepper при необходимости хранится отдельно от БД. JWT подписан, но не зашифрован: пароль в payload не помещают.
