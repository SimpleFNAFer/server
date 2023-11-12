create schema server;

create table server.request (
    id serial not null,
    source_ip text not null,
    timestamp timestamp not null default now(),

    primary key (id)
);