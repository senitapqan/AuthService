CREATE TABLE IF NOT EXISTS t_users (
    id serial Primary key,
    username varchar(50) unique not null,
    password varchar(50) not null, 
    email varchar(255) not null, 
    name varchar(50) not null, 
    surname varchar(50) not null,
    phone varchar(50),
    archive bool default false
);

CREATE TABLE IF NOT EXISTS t_roles (
   id serial Primary key, 
   role_name varchar(50) not null
);

CREATE TABLE IF NOT EXISTS t_users_roles (
    user_id int not null,
    role_id int not null, 
    Primary key(user_id, role_id), 
    FOREIGN key (user_id) REFERENCES t_users(id) on delete cascade,
    FOREIGN key (role_id) REFERENCES t_roles(id)
);

CREATE TABLE IF NOT EXISTS t_client (
   id serial Primary key, 
   user_id int not null,
   FOREIGN KEY (user_id) REFERENCES t_users(id) on delete cascade,
   created_time timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS t_teacher (
    id serial Primary key, 
    user_id int not null,
    FOREIGN KEY (user_id) REFERENCES t_users(id) on delete cascade,
    created_time timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS t_admin (
    id serial Primary key, 
    user_id int not null,
    FOREIGN KEY (user_id) REFERENCES t_users(id) on delete cascade,
   	created_time timestamp default current_timestamp
);

insert into t_roles (id, role_name) values(1, 'ADMIN');
insert into t_roles (id, role_name) values(2, 'TEACHER');
insert into t_roles (id, role_name) values(3, 'CLIENT');
