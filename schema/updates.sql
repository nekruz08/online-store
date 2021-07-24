

UPDATE items SET title = 'iphone 11 pro' ,categories_id = 1 ,price =100  where id = 4 RETURNING items;

insert into roles(name) values ('ADMIN'),('USER');

select * from roles;

UPDATE users SET role_id = 1 WHERE id = 1;



