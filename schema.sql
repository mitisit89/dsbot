
create table if not exists user(
    id serial primary key,
    name varchar(255) not null,

)
create table if not exists movie (
    id serial primary key ,
    name varchar(255) not null,
    watched bool default false,
);
create table if not exists game(
    id serial primary key,
    name varchar(255) not null
)

create table if not exists user_movie_rating(
    id serial primary key ,
    rating smallint check(rating > 0 and rating <= 10),
    foreign key(user_id) references user(id) on delete cascade,
    foreign key(movie_id) references movie(id) on delete cascade
)

create table if not exists games_rating(
    id serial primary key,
    name varchar(255) not null,
    rating smallint check(rating > 0 and rating <= 10),
    foreign key(user_id) references user(id) on delete cascade
    foreign key(game_id) references game(id) on delete cascade
)
