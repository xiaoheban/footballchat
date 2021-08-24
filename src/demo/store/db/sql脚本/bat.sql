//单表
create table posts{
id serial primary key,
content text,
author varchar(255)
};

//有关联关系表
drop table posts cascade if exists;   //级联删除
drop table comments if exists;

create table posts(
  id serial primary key,
  content text,
  author varchar(255)
);

create table comments(
  id serial primary key ,
  content text,
  author varchar(255),
  post_id integer references posts(id)
);
