create table if not exists discord_user (id serial primary key, name varchar(255) not null);

create table if not exists games (id serial primary key, name varchar(255) not null);

create table if not exists movies (
    id serial primary key,
    name varchar(255) not null,
    traller varchar(512) default null,
    watched bool default false not null
);

create table if not exists movie_rating (
    id serial primary key,
    rating smallint check (
        rating > 0
        and rating <= 10
    ),
    discord_user_id int,
    movie_id int,
    foreign key (discord_user_id) references discord_user (id) on delete cascade,
    foreign key (movie_id) references movies (id) on delete cascade
);

create table if not exists games_rating (
    id serial primary key,
    name varchar(255) not null,
    rating smallint check (
        rating > 0
        and rating <= 10
    ),
    discord_user_id int,
    game_id int,
    foreign key (discord_user_id) references discord_user (id) on delete cascade,
    foreign key (game_id) references games (id) on delete cascade
);

create index movie_name_idx on movies (name);

create index game_name_idx on games (name);
