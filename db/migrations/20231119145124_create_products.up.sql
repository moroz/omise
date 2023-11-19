create table products (
  id uuid primary key,
  name text not null,
  slug text not null unique,
  description text,
  price int not null,
  inserted_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc'),
  check (price >= 0)
);
