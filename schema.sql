create table if not exists watchlist (
id integer primary key autoincrement,
name string not null unique,
watched integer default 0
);
