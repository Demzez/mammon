## Errors struct:
1. Main type:
   - 'operation' | 'what' 'how(successfully | failed)': 'error'
     - operation - the package or location where something is happening
     - what - what is working correctly or incorrectly
     - how - what exactly is happening
     - error - stack error

## Way to throw errors:
1. If some system error arises, wrap it in this way:

``fmt.Errorf("%s | schema init failed: %w", op, err)``

2. Then throw this error without wrapping to function when it will log:

``log := logger.With(slog.String("op", op))``

...

``log.Error("database init failed", slog.String("error", err.Error()))``

Обработчик: 
- отвечает за все логирование и обяснение ошибок
- возварщает необходимые http коды

Сервис:
- может выбрасывать свои ошибки
- добавляет trace контекст ошибки в виде собственного `op`

Репозиторий 
- только выбрасывет ошибки и ничего не логирует