create type consume_type as enum ('push', 'pull');

create table if not exists tb_topic (
    id serial primary key,
    name varchar(128) NOT NULL,
    consume_type consume_type NOT NULL DEFAULT 'push', 
    is_deleted boolean NOT NULL DEFAULT false
);

create table if not exists tb_dispatch (
    id serial primary key,
    topic_id integer NOT NULL,
    "target" varchar(255) NOT NULL,
    is_deleted boolean NOT NULL DEFAULT false,
);
