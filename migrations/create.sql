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
    phone text not null
);

create table if not exists tooth (
    id text primary key,
    patient text not null,
    number int not null,
    t int default 0,
    r int default 0,
    b int default 0,
    l int default 0,
    center int default 0,
    foreign key(patient) references patient(id) on delete cascade
);
