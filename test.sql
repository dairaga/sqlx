set names utf8;

drop database if exists mytest;
create database mytest charset utf8;

grant all on mytest.* to 'test'@'%' identified by 'test';
flush privileges;

use mytest;

create table test1 (
    c01 bit not null default 0,
    c01_1 bit(1) not null default 0,
    c02_s tinyint not null default 0,
    c02_u tinyint unsigned not null default 0,
    c03 bool not null default false,
    c04_s smallint not null default 0,
    c04_u smallint unsigned not null default 0,
    c05_s mediumint not null default 0,
    c05_u mediumint unsigned not null default 0,
    c06_s int not null default 0,
    c06_u int unsigned not null default 0,
    c07_s bigint not null default 0,
    c07_u bigint unsigned not null default 0,
    c08_s decimal not null default 0,
    c08_u decimal unsigned not null default 0,
    c09_s float not null default 0,
    c09_u float unsigned not null default 0,
    c10_s double not null default 0,
    c10_u double unsigned not null default 0,
    c11 date not null,
    c12 time not null,
    c13 year not null,
    c14 datetime not null,
    c15 timestamp not null,
    c16 char(10) not null,
    c17 varchar(10) not null,
    c18 binary(10) not null,
    c19 varbinary(10) not null,
    c20 tinyblob not null,
    c21 tinytext not null,
    c22 blob not null,
    c23 text not null,
    c24 mediumblob not null,
    c25 mediumtext not null,
    c26 longblob not null,
    c27 longtext not null,
    c28 enum('a', 'b', 'c') not null,
    c29 set('a', 'b', 'c') not null
) charset utf8;

insert into test1 values(b'1', b'1', -2, 2, true, -4, 4, -5, 5, -6, 6, -7, 7, -8, 8, -9, 9, -10, 10, '2019-01-01', "123:04:05", '2019', '2014-09-23 10:01:02', '2014-09-23 10:01:02', 'C16', 'C17', 'C18', 'C19', 'C20', 'C21', 'C22', 'C23', 'C24', 'C25', 'C26', 'C27', 'a', 'b');

create table test2 (
    c01 bit,
    c01_1 bit(1),
    c02_s tinyint,
    c02_u tinyint unsigned,
    c03 bool,
    c04_s smallint,
    c04_u smallint unsigned,
    c05_s mediumint,
    c05_u mediumint unsigned,
    c06_s int,
    c06_u int unsigned,
    c07_s bigint,
    c07_u bigint unsigned,
    c08_s decimal,
    c08_u decimal unsigned,
    c09_s float,
    c09_u float unsigned,
    c10_s double,
    c10_u double unsigned,
    c11 date,
    c12 time,
    c13 year,
    c14 datetime,
    c15 timestamp,
    c16 char(10),
    c17 varchar(10),
    c18 binary(10),
    c19 varbinary(10),
    c20 tinyblob,
    c21 tinytext,
    c22 blob,
    c23 text,
    c24 mediumblob,
    c25 mediumtext,
    c26 longblob,
    c27 longtext,
    c28 enum('a', 'b', 'c'),
    c29 set('a', 'b', 'c')
) charset utf8;

insert into test2 values();
update test2 set c15 = '2014-09-23 10:01:02'