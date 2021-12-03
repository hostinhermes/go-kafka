# go-kafka
Projeto de estudo do kafka implementado em golang

Implementado o envio e o consumo de mensagens em uma fila

# Subir ambiente do go e do kafka
-> docker-compose up -d

# Criar fila no kafka 
Acessar o container do kafka

-> docker exec -it go-kafka_kafka_1 bash

Executar comando para criação do topico

-> kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3 --replication-factor=1

# Rodar producer
Acessar o container do go

-> docker exec -it gokafka bash

Rodar programa main.go do diretorio producer

-> go run cmd/producer/main.go

# Rodar consumer
Acessar o container do go

-> docker exec -it gokafka bash

Rodar programa main.go do diretorio consumer

-> go run cmd/producer/main.go

