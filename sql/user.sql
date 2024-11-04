create table family (
	id serial primary key
);

create table app_user (
	id serial primary key,
	created_at timestamp default now(),
	email varchar(50) unique,
	username varchar(50) unique,
	telegram_id varchar(50) unique,
	avatar_path varchar(50),
	family_id int references family(id) on delete set null
);

