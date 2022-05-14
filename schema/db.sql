create table user(
id integer primary key autoincrement,
first_name varchar(512) not null,
last_name varchar(512),
mobile varchar(30),
email varchar(512),
created_at varchar(512),
updated_at varchar(512)
);

create table darood_stats(
id integer primary key autoincrement,
darood_id integer not null,
user_id integer not null,
date varchar(20) // dd-mm-yyyy
count integer default 0
);


create table darood_stats(
id integer primary key autoincrement,
darood_id integer not null,
user_id integer not null,
date varchar(20) // dd-mm-yyyy
count integer default 0
);