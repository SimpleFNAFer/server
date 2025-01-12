create schema server;

create table server.request
(
    id        serial    not null,
    source_ip text      not null,
    timestamp timestamp not null default now(),

    primary key (id)
);

create table server.blocked_ip
(
    id        serial    not null,
    ip        text      not null,
    timestamp timestamp not null default now(),

    primary key (id),

    unique (ip)
);

create table server.admin
(
    id            serial not null,
    login         text   not null,
    password_hash text   not null,

    primary key (id),

    unique (login)
);
