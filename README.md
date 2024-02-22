# desafio-hitss

https://aprendagolang.com.br/2022/01/26/benchmark-dos-routers-http-chi-vs-gorilla-mux/

docker run -d --name api-users -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
docker ps
docker exec -it api-users psql -U postgres

CREATE DATABASE api_users;
\c api_users
create table users (id serial primary key, nome varchar, sobrenome varchar, contato varchar, endereço text, data_nasc varchar, cpf varchar);

create user dbuser with password '1122';
grant all privileges on database api_users to dbuser;
grant all privileges on all tables in schema public to dbuser;
grant all privileges on all sequences in schema public to dbuser;