create database prepare default charset=utf8;

create table user (
	id int(11) not null auto_increment, 
	phone_number varchar(15) not null, 
	password varchar(255) not null, 
	primary key(id)
)engine=innodb default charset=utf8;
