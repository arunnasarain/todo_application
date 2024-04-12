CREATE TABLE IF NOT EXISTS app_user (
    ID int NOT NULL PRIMARY KEY ,
    Title varchar(50) NOT NULL DEFAULT 'NEW TODO',
    Content varchar(500) NOT NULL,
    CreatedAt timestamp default CURRENT_TIMESTAMP not null,
    UpdatedAt timestamp default CURRENT_TIMESTAMP not null
)