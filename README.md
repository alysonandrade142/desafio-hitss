
# API Golang com PostgreSQL e Worker para Filas

API Golang com PostgreSQL e Worker para Filas
Este repositório apresenta uma API em Golang que interage com um banco de dados PostgreSQL para salvar dados. Além disso, inclui um worker para a execução e gerenciamento de duas filas distintas: uma chamada `queue-processing` para processamento de atividades e outra chamada `queue-response` para gerenciamento de respostas. As instruções a seguir ajudarão você a configurar e utilizar esta aplicação.


Pré-requisitos
Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina antes de começar.

Docker Installation Guide
Docker Compose Installation Guide
Instruções
Clone este repositório em sua máquina local:

```
bash
Copy code
git clone https://github.com/seu-usuario/nome-do-repositorio.git
Navegue até o diretório clonado:
```

```
cd nome-do-repositorio
Execute o seguinte comando para iniciar o contêiner PostgreSQL e o worker em segundo plano:
```
```
docker-compose up -d
```
Para inicializar o banco
```
docker exec -it  postgres psql -U dbuser api_users
```
Para criar a tabela de dados
```
create table users (id serial primary key, nome varchar, sobrenome varchar, contato varchar, endereço text, data_nasc varchar, cpf varchar);
```

Execute a API Golang. O código-fonte da API pode ser encontrado nos diretórios apropriados.

Para iniciar a API, o consumo dos serviços podem ser efetuados via postman..
```
cmd -> api -> main.go
```
Worker
```
cmd -> consumer -> main.go
```
Você pode começar a utilizar a API para salvar dados no banco de dados PostgreSQL e aproveitar o sistema de filas para tarefas assíncronas.