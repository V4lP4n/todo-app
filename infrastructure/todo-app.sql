CREATE TABLE IF NOT EXISTS 'task'(
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    `data` varchar(180) default "",
    `status` boolean default 0,
    `list_id` INTEGER not NULL,
    FOREIGN KEY(list_id) references list(id)
);

CREATE TABLE IF not EXISTS 'list'(
    `id` INTEGER not NULL PRIMARY KEY AUTOINCREMENT,
    `name` varchar(100) default "Insert name here"

);
 
Insert into list (id, name) values (0, 'default');
Insert into task (data,list_id) 
values 
("test1",0),
("test2",0);