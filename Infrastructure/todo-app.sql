CREATE TABLE IF NOT EXISTS 'todo'(
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    `data` varchar(180) default "",
    `status` boolean default 0
)