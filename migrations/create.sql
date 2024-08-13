create table if not exists patient (
    id text primary key,
    name text not null,
    surname text not null,
    birthdate text not null,
    pesel text,
    address text not null,
    medicine text,
    illness integer default 0,
    registered text not null,
    child integer default 0,
    phone text not null
)
