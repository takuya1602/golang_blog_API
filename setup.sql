create function update_timestamp()
returns trigger as $$
begin
    new.updated_at := now();
    return new;
end;
$$ language 'plpgsql';

create table categories (
    id serial primary key,
    name varchar(255) unique,
    slug varchar(255) unique
);

create table sub_categories (
    id serial primary key,
    name varchar(255) unique,
    slug varchar(255) unique,
    parent_category_id integer references categories(id)
);

create table posts (
    id serial primary key,
    title varchar(255) unique,
    slug varchar(255) unique,
    eye_catching_img varchar(2048),
    content text,
    meta_description text,
    is_public boolean,
    created_at timestamp with time zone default current_timestamp not null,
    updated_at timestamp with time zone default current_timestamp not null,
    sub_category_id integer references sub_categories(id)
);

create table users (
    id serial primary key,
    username varchar(255),
    password varchar(255),
    is_admin boolean default TRUE
);

create trigger update_posts_timestamp before update on posts for each row execute procedure update_timestamp();

insert into categories 
    (name, slug)
values 
    ('プログラミング', 'programming'),
    ('データベース', 'database'),
    ('機械学習', 'machine-learning');

insert into sub_categories
    (name, slug, parent_category_id)
values
    ('Go言語', 'golang', 1),
    ('Python', 'python', 1),
    ('PostgreSQL', 'postgresql', 2),
    ('MySQL', 'mysql', 2),
    ('Kaggle', 'kaggle', 3),
    ('アルゴリズム', 'algorithm', 3);

insert into posts
    (category_id, sub_category_id, title, slug, eye_catching_img, content, meta_description, is_public)
values
    (1, 1, 'Go入門', 'introduction-of-go', 'test.jpeg', 'Go言語は近年注目されているWebアプリケーションの構築のための言語です。', 'Go言語入門', 'false'),
    (1, 2, 'Python入門', 'introduction-of-python', 'test.jpeg', 'Pythonは機械学習分野でよく用いられているインタプリタ言語です。', 'Python入門', 'false');

