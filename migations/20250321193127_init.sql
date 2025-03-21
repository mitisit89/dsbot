-- +goose Up
-- +goose StatementBegin
CREATE table if not exists discord_user (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE  if not exists movies (
    id serial PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    trailer VARCHAR(500),
    watched BOOLEAN DEFAULT FALSE,
    discord_user_id INT,
    CONSTRAINT fk_movies_user FOREIGN KEY (discord_user_id) REFERENCES discord_user (id) ON DELETE SET NULL
);

CREATE table if not exists games (
    id serial PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    discord_user_id INT,
    CONSTRAINT fk_games_user FOREIGN KEY (discord_user_id) REFERENCES discord_user (id) ON DELETE SET NULL
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists games_rating;
drop table if exists movie_rating;
drop table if exists games;
drop table if exists movies;
drop table if exists discord_user;
-- +goose StatementEnd
