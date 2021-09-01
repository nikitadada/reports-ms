#!/bin/sh -e
# ВАЖНО!
# Данный скрипт является стандартным для всех микросервисов данного типа, его нельзя менять, предварительно не обсудив
# с тимлидом.
PROTO_SERVER_PATH="internal/grpc"
PROTO_CLIENT_PATH="internal/grpcclient"
UNIT_COVERAGE_MIN=84

# Запуск prototool
run_prototool(){
  docker run --rm -v "$(pwd):/work" uber/prototool:1.10.0 $@
}

# Обрабатывает прото файлы prototool
process_proto_files(){
  local COMMAND="$1"
  local PROTO_DIR="$2"

  if [ ! -d "$PROTO_DIR" ]; then
    return 0
  fi

  run_prototool prototool "$COMMAND" "$PROTO_DIR"
}

# Генерация прото файлов
gen_proto(){
  # Удаление сгенерированных файлов из прото
  for CURPATH in "$PROTO_SERVER_PATH" "$PROTO_CLIENT_PATH"; do
    rm -Rf $CURPATH/gen/*
  done

  process_proto_files all $PROTO_SERVER_PATH
  process_proto_files all $PROTO_CLIENT_PATH

  for CURPATH in "$PROTO_SERVER_PATH" "$PROTO_CLIENT_PATH"; do
    if [ -d "$CURPATH" ]; then
      run_prototool chown -R "$(id -u)":"$(id -g)" "/work/$CURPATH/gen"
    fi
  done
}

# Запуск unit-тестов
unit(){
  echo "run unit tests"
  go test ./...
}

unit_race() {
  echo "run unit tests with race test"
  go test -race ./...
}

# Запуск go-lint
lint(){
  echo "run linter"
  go mod vendor
  docker run --rm -v $(pwd):/work:ro -w /work golangci/golangci-lint:latest golangci-lint run -v
  rm -Rf vendor
}

# Запуск линтера proto файлов
lint_proto(){
  echo "run proto linter"
  process_proto_files lint $PROTO_SERVER_PATH
  process_proto_files lint $PROTO_CLIENT_PATH
}

fmt() {
  echo "run go fmt"
  go fmt ./...
}

vet() {
  echo "run go vet"
  go vet ./...
}

unit_coverage() {
  echo "run test coverage"
  go test -coverpkg=./... -coverprofile=cover_profile.out.tmp $(go list ./internal/...)
  # remove generated code and mocks from coverage
  < cover_profile.out.tmp grep -v -e "mock" -e "\.pb\.go" > cover_profile.out
  rm cover_profile.out.tmp
  CUR_COVERAGE=$( go tool cover -func=cover_profile.out | tail -n 1 | awk '{ print $3 }' | sed -e 's/^\([0-9]*\).*$/\1/g' )
  rm cover_profile.out
  if [ "$CUR_COVERAGE" -lt $UNIT_COVERAGE_MIN ]
  then
    echo "coverage is not enough $CUR_COVERAGE < $UNIT_COVERAGE_MIN"
    return 1
  else
    echo "coverage is enough $CUR_COVERAGE >= $UNIT_COVERAGE_MIN"
  fi
}

# Запуск всех тестов
test(){
  fmt
  vet
  unit
  unit_race
  unit_coverage
  lint
  lint_proto
}

# Подтянуть зависимости
deps(){
  go get ./...
}

# Собрать исполняемый файл
build(){
  deps
  go build ./cmd/report-action
}

# Запустить сбор метрик нагрузки на cpu из pprof
pprof_cpu(){
  local SECS=${3:-$PPROF_DEFAULT_CPU_DURATION}
  local HOST=$2

  go tool pprof -http :$PPROF_UI_PORT $HOST/debug/pprof/profile?seconds=$SECS
}

# Запустить сбор метрик памяти из pprof
pprof_heap(){
  local HOST=$2

  go tool pprof -http :$PPROF_UI_PORT $HOST/debug/pprof/heap
}

# Добавьте сюда список команд
using(){
  echo "Укажите команду при запуске: ./run.sh [command]"
  echo "Список команд:"
  echo "  unit - запустить unit-тесты"
  echo "  unit_race - запуск unit тестов с проверкой на data-race"
  echo "  unit_coverage - запуск unit тестов и проверка покрытия кода тестами"
  echo "  lint - запустить все линтеры"
  echo "  lint_proto - запустить линтер proto файлов"
  echo "  test - запустить все тесты"
  echo "  deps - подтянуть зависимости"
  echo "  build - собрать приложение"
  echo "  fmt - форматирование кода при помощи 'go fmt'"
  echo "  vet - проверка правильности форматирования кода"
  echo "  gen_proto - генерация прото файлов (для клиентов и сервера)"
  echo "  pprof_cpu HOST [SECONDS] - сбор метрик нагрузки на cpu из pprof"
  echo "  pprof_heap HOST - запустить сбор метрик памяти из pprof"
}

############### НЕ МЕНЯЙТЕ КОД НИЖЕ ЭТОЙ СТРОКИ #################

command="$1"
if [ -z "$command" ]
then
 using
 exit 0;
else
 $command $@
fi
