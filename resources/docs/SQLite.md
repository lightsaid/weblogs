# SQLite

参考 [菜鸟教程](https://www.runoob.com/sqlite/sqlite-create-database.html)


``` SQL
drop table if exists "users";
create table users (
    id integer not null primary key autoincrement,
    email text not null,
    username text not null,
    avatar text,
    is_admin int default 0,
    is_delete int default 0,
    created_at text not null,
    updated_at text not null
);

create unique index unq_email on users(email); 

INSERT into users(email, username, created_at, updated_at) 
	values('zhangsan@qq.com','张三', datetime('now'),datetime('now'));
	
SELECT username, email, DATETIME(created_at), DATE(updated_at) from users;  

-- datetime 操作
-- insert into users(craeted_at) values(datetime('now'));
-- select datatime(created_at) from users;
```

docker-compose.yml
``` yml
services: 
  sqlite3:
    container_name: sqlite3
    image: nouchka/sqlite3:latest
    stdin_open: true
    tty: true
    volumes:
      - ../database/sqlite3:/root/db/
    ports:
      - "1024:1024"
```

### 在Golang中使用 sqlite3
如果报错 Sqlite3 stdlib.h: No such file or directory,
则需要安装 sudo apt-get install g++