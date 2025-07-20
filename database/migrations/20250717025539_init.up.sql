create table if not exists users (
  id serial primary key,
  name text not null,
  email text unique not null,
  password text not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);

create table if not exists user_tokens (
  id serial primary key, 
  user_id integer not null references users(id) on delete cascade,
  access_token text not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
)