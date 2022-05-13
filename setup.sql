create table parent_categories (
    id serial primary key,
    category_name varchar(255),
    slug varchar(255)
);

create table categories (
    id serial primary key,
    category_name varchar(255),
    slug varchar(255),
    parent_category_id integer references parent_categories(id)
);

create table posts (
    id serial primary key,
    category_id integer references categories(id),
    slug varchar(255),
    title varchar(255),
    eye_catching_img varchar(2048),
    content text
);