create table if not exists watchlist (
id integer primary key autoincrement,
name string not null unique,
wattched integer default 0
);
