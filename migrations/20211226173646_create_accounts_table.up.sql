create table if not exists account(
    id uuid primary key default gen_random_uuid(),
    email varchar(255) unique not null,
    username varchar(16) unique not null,
    password varchar(255) not null,
    is_verified boolean default false not null,
    is_archived boolean default false not null,
    created_at timestamp with time zone default current_timestamp not null,
    updated_at timestamp with time zone default current_timestamp not null
);

