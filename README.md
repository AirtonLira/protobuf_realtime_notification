# Protobuf Realtime Notification

Este projeto demonstra a implementação de um serviço de notificações em tempo real utilizando Protocol Buffers (Protobuf) e gRPC com Golang, com troca de mensagem em streaming.

## Estrutura do Projeto

- protobuf_realtime_notification/
- cmd/
  - grpc/
    - main.go
  - client/
    - main.go
- internal/
  - proto/
    - notification/
      - notification.proto
      - notification.pb.go
      - notification_grpc.pb.go
- go.mod
- go.sum
- README.md


## Requisitos

- Golang 1.20 ou superior
- Protocol Buffers Compiler (`protoc`)
- Plugins do Protobuf para Go:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`

## Próximas Features

- 🚀 DuckDB: Adicionaremos o suporte ao DuckDB para obter novos eventos a serem notificados. Imagine uma avalanche de notificações super rápidas e precisas! 📈

- ⏰ Apache Airflow: Vamos integrar o Apache Airflow para agendar notificações em batch retroativas. Sim, você poderá enviar aquelas notificações que esqueceu, como se fossem novinhas em folha! 📅

## Instalação

### Passo 1: Instalar as Dependências

Instale os plugins do Protobuf para Go:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
````


#### Passo 2: adicione o GOPATh
export PATH="$PATH:$(go env GOPATH)/bin"

### Passo 3: Clonar o Repositório
```sh
git clone https://github.com/AirtonLira/protobuf_realtime_notification.git
cd protobuf_realtime_notification
````

### Passo 4: Gerar os Arquivos Go a partir do Protobuf
```sh
protoc --go_out=. --go-grpc_out=. internal/proto/notification/notification.proto
````

### Passo 5: Executar o Servidor
### Navegue até o diretório cmd/grpc e execute o servidor gRPC:

```sh
cd cmd/grpc
go run main.go
````

### Passo 6: Executar o Cliente
### Abra uma nova janela de terminal, navegue até o diretório cmd/client e execute o cliente gRPC:
```sh
cd cmd/client
go run main.go
````