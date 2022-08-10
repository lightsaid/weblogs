-- 用户表
drop table if exists "users";
create table users (
    id integer not null primary key autoincrement,
    email text not null,
    username text not null,
    password varchar(16) not null,
    avatar text,
    if_admin int default 0,  -- 是否管理员(0:是|1:否)
    -- check 约束
    -- active int check(active > -2) default 0, -- 状态 (0:正常|-1:删除)
    active int check(active in (-1, 0, 1)), -- 状态 (-1:删除0:正常|1:活跃)
    -- 此处语法 default (xxx(yyy))
    -- 时间需要使用 localtime, 不然可能存在8个小时时差
    created_at text not null default (datetime('now', 'localtime')),
    updated_at text not null default (datetime('now', 'localtime'))
);
create unique index unq_email on users(email); 


-- 文章/帖子表
drop table if exists "posts";
create table posts(
    id integer not null primary key autoincrement,
    user_id integer not null,
    author text not null,
    title varchar(255) not null,
    content text not null,
    thumb text, --缩略图
    readings int default 0, -- 查看人数
    comments int default 0, -- 评论数
    likes int default 0, -- 喜欢数
    active int check(active in (-1, 0)) default 0, -- 状态 (-1:删除 | 0:正常)
    created_at text not null default (datetime('now', 'localtime')),
    updated_at text not null default (datetime('now', 'localtime')),
    foreign key(user_id) references users(id),
    foreign key(author) references users(username)
);

-- 文章分类表
drop table if exists "categories";
create table categories(
    id integer not null primary key autoincrement,
    user_id integer not null,
    parent_id integer,
    if_parent integer check(if_parent in (0, 1)) default 0, -- (自身是否是父类)是否含有子级 0 没有, 1有
    name varchar(20) not null,
    thumb text, --icon/缩略图
    foreign key (user_id) references users(id)
);
create unique index unq_name on categories(name);

-- 文章和分类映射表
drop table if exists "pc_mapping";
create table pc_mapping(
    id integer not null primary key autoincrement,
    post_id integer not null,
    cate_id integer not null,
    foreign key (post_id) references posts(id),
    foreign key (cate_id) references categories(id)
);


-- 属性表 (包含tag和mark)
drop table if exists "attributes";
create table attributes(
    id integer not null primary key autoincrement,
    user_id integer not null,
    kind varchar(1) not null check(kind in ('T','M')), -- T: 标签tag, M:标记mark
    name varchar(20) not null,
    foreign key (user_id) references users(id)
);

-- 文章属性映射表
drop table if exists "pa_mapping";
create table pa_mapping(
    id integer not null primary key autoincrement,
    post_id integer not null,
    attr_id integer not null,
    foreign key (post_id) references posts(id),
    foreign key (attr_id) references attributes(id)
);

-- 评论表
drop table if exists "comments";
create table comments(
    id integer not null primary key autoincrement,
    post_id integer not null,
    email varchar(100) not null,
    nickename varchar(20) not null,
    content text not null,
    parent_id integer,
    likes int default 0, -- 点赞数
    replynum int default 0, --回复数
    active int check(active in (-1, 0)) default 0, -- 状态 (-1:删除 | 0:正常)
    created_at text not null default (datetime('now', 'localtime')),
    updated_at text not null default (datetime('now', 'localtime'))
);

-- 点赞表 (存在user_id则是注册用户点赞, 没有则是根据ip地址添加点赞)
drop table if exists "likes";
create table likes (
    id integer not null primary key autoincrement,
    ip_address varchar(100) not null,
    kind varchar(1) check(kind in ('P','C')), -- P:文章, C:评论
    user_id integer, -- 用户id
    pc_id integer not null -- 文章ID或者评论ID 
);

-- 插入也需要指定本地时间,不然会有时差问题
INSERT into users(email, username, password, active, created_at, updated_at) 
	values('zhangsan@qq.com','张三', 'abc123', -1, datetime('now', 'localtime'), datetime('now', 'localtime'));
	
SELECT username, email, active, datetime(created_at), date(updated_at) from users;  

select * from users;