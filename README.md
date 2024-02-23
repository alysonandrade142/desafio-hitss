# desafio-hitss

docker-compose up -d
docker exec -it  postgres psql -U dbuser api_users
create table users (id serial primary key, nome varchar, sobrenome varchar, contato varchar, endereço text, data_nasc varchar, cpf varchar);