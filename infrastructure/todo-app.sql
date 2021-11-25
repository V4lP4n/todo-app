CREATE TABLE IF NOT EXISTS 'task'(
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    `data` varchar(180) default "",
    `status` boolean default 0,
    `list_id` INTEGER not NULL,
    FOREIGN KEY(list_id) references list(id)
);

CREATE TABLE IF not EXISTS 'task_list'(
    `id` INTEGER not NULL PRIMARY KEY AUTOINCREMENT,
    `name` varchar(100) default "Insert name here",
    `user_id` INTEGER not NULL,
    FOREIGN KEY(user_id) references user(id)

);

CREATE TABLE if not EXISTS 'user'(
    `id` INTEGER not NULL PRIMARY KEY AUTOINCREMENT,
    `uname` varchar(100),
    `passwd_hash` varchar(100)
);

Insert into user (id, uname, passwd_hash) values (0, 'test', 'qew'); 
Insert into task_list (id, name, user_id) values (0, "default", 0);
Insert into task (data,list_id) 
values 
("test1",0),
("test2",0);