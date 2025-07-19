# MWS CLI tool

## Сборка
- Linux: 
  ```bash
  go build -o mws .
  ```
- Windows:
  ```powershell
  go build -o mws.exe .
  ```
- Docker
  ```bash
  docker build -t mws .
  ```

## Запуск
- Linux: 
  ```bash
  ./mws -h
  ```
- Windows:
  ```powershell
  ./mws.exe -h
  ```
- Docker
  ```bash
  docker run --rm mws -h
  ```

## Тесты
Находятся в файле ./cmd/commands/profile/profile_test.go