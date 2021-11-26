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

Insert into user (uname, passwd_hash) values ('test', 'qew'); 
Insert into user (uname, passwd_hash) values ('valera', 'ewq'); 
Insert into task_list (name, user_id) values ("default", 1);
Insert into task_list (name, user_id) values ("testtaskofusertest", 1);
Insert into task_list (name, user_id) values ("testtaskofuservalera", 2);
Insert into task (data,list_id) 
values 
("test1",1),
("test2",1),
("test3",3),
("test4",2),
("test5",3),
("test6",1),
("test7",3);