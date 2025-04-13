# Proto files for ubrato microservices

## Generate for golang

## Integration

1. Добавить сабмодуль
```bash
git submodule add git@gitlab.ubrato.ru:ubrato/proto.git proto
```

2. Проверить наличие настроек в makefile сервиса
```
PROTO_PATH=./proto
PROTO_OUT=./internal/models/gen/proto
PROTO_OUT_MODULE=gitlab.ubrato.ru/ubrato/НАЗВАНИЕ_ПРОЕКТА/internal/models/gen/proto

include ./proto/proto.mk
```

3. Cгенерировать go код из proto
```bash
make generate.go
```

4. Получение обновлений proto для сервиса
```bash
git submodule update --init --recursive --remote
```