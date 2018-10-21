
-- DROP TABLE [ IF EXISTS ] name [, ...] [ CASCADE | RESTRICT ]
-- IF EXISTS 如果指定的表不存在，那么发出一个 notice 而不是抛出一个错误。
-- name 要删除的现存表的名字(可以有模式修饰)。
-- CASCADE 级联删除依赖于表的对象(比如视图)。
-- RESTRICT 如果存在依赖对象，则拒绝删除该表。这个是缺省。
drop table if exists posts cascade;
drop table if exists comments;

CREATE TABLE posts (
  id serial PRIMARY KEY,
  content text,
  author VARCHAR(255)
);

CREATE TABLE comments (
  id serial PRIMARY KEY,
  content text,
  author VARCHAR(255),
  post_id INTEGER REFERENCES posts(id)  -- post_id作为外键(foreign key)，对posts表的主键id进行引用
);
