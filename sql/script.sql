create database if not exists devconnect;
use devconnect;

drop table if exists users;

create table users (
    id int auto_increment,
    name varchar(50) not null,
    email varchar(50) not null,
    password varchar(50) not null,
    nick varchar(50) not null,
    created_at datetime not null default current_timestamp(),
    updated_at datetime not null default current_timestamp() on update current_timestamp,
    primary key (id),
    unique key email (email)
) engine=InnoDB default charset=utf8mb4;

--mysql -u root -p devconnect < /c/Coding/Go/DevConnect/sql/script.sql