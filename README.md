# Тестовое задание в MEDODS
_____________

В этом репозитории представлена моя реализация части сервиса аутентификации для работы с Access и Refresh токенами.
Для улучшения масштабируемости и иерархии проекта мною были использованы принципы **чистой архитектуры** (или onion-like) с использованием четырех слоев: handler, service, infrastructure и models.
Access токен представляет собою JWT, Refresh же случайно генерируется.
**Mongo DB** разворачивал в Docker, настройки подключения выставляются в configs/config.
В проекте можно улучшить в перспективе несколько вещей:
- Написать Unit-тесты для слоев;
- Перевести логгирование на logrus, чтобы оно было многоуровневым;
- Улучшить инкапсулированность слоев;
- Оптимизировать скорость работы приложения.
