-- 用户表
drop table if exists "users";
create table users (
    id integer not null primary key autoincrement,
    email text not null,
    username text not null,
    password varchar(16) not null,
    avatar text,
    active int not null check(active in (-1, 0, 1)) default 0, -- 状态 (-1:删除|0:正常|1:活跃)
    role int not null default 0,  -- 角色(0:普通用户|1:管理员｜99:超级管理员)
    created_at text not null default (datetime('now', 'localtime')),
    updated_at text not null default (datetime('now', 'localtime'))
);
create unique index unq_email on users(email); 

-- 文章表
drop table if exists "posts";
create table posts(
    id integer not null primary key autoincrement,
    user_id integer not null,
    author text not null,
    title varchar(255) not null,
    content text not null,
    thumb text, --缩略图
    readings int not null default 0, -- 查看人数
    comments int not null default 0, -- 评论数
    likes int not null default 0, -- 喜欢数
    active int not null check(active in (-1, 0, 1)) default 1, -- 状态 (-1:删除 | 0:暂存 | 1:正常)
    created_at text not null default (datetime('now', 'localtime')),
    updated_at text not null default (datetime('now', 'localtime')),
    foreign key(user_id) references users(id),
    foreign key(author) references users(username)
);

-- tag 表 (包含tag和mark)
drop table if exists "tags";
create table tags(
    id integer not null primary key autoincrement,
    user_id integer not null,
    name varchar(16) not null unique,
    foreign key (user_id) references users(id)
);

-- posts和tags映射表
drop table if exists "post_tags";
create table post_tags(
    id integer not null primary key autoincrement,
    post_id integer not null,
    tag_id integer not null,
    foreign key (post_id) references posts(id),
    foreign key (tag_id) references tags(id)
);

----------- 初始化一些数据 -------------

-- password: abc123
insert into users (email, username, password)
    values("lightsaid@163.com", "lightsaid", "$2a$10$bUnC/vaN/LeSc49YBf2iWu9G5CJmIO7Nybz4YHNo00pUW8j9oxTZK");

insert into tags(user_id, name) values(1, "Golang");
insert into tags(user_id, name) values(1, "Docker");
insert into tags(user_id, name) values(1, "TypeScript");
insert into tags(user_id, name) values(1, "React");

insert into posts(user_id, author, title, content) values(1, "lightsaid", "Golang Blog", "Golang编写Blog心路历程...");
insert into posts(user_id, author, title, content) values(1, "lightsaid", "Golang Redis", "Golang 操作 Redis 知识点～");
insert into posts(user_id, author, title, content) values(1, "lightsaid", "Docker 探索", "Docker 探索之旅～");
insert into posts(user_id, author, title, content) values(1, "lightsaid", "TypeScript 之泛型", "TypeScript的泛型好强大的样子...");
insert into posts(user_id, author, title, content) values(1, "lightsaid", "React toolkit", "React toolkit 技巧总结...");
insert into posts(user_id, author, title, content) values(1, "lightsaid", "React Router", "React Router 版本太坎坷～");

insert into post_tags(post_id, tag_id) values(1, 1);
insert into post_tags(post_id, tag_id) values(2, 1);
insert into post_tags(post_id, tag_id) values(3, 2);
insert into post_tags(post_id, tag_id) values(4, 3);
insert into post_tags(post_id, tag_id) values(5, 4);
insert into post_tags(post_id, tag_id) values(6, 4);


select * from users;
select * from tags;
select * from posts;


select * from posts p join post_tags pt on p.id = pt.post_id where pt.tag_id = 1 limit 1 offset 0;

select id, email, password, username, avatar, role, active from users where email="lightsaid@163.com" and active != -1 limit 1;


