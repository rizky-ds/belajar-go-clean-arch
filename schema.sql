use development;

drop table if exists todo;

create table todo(
  id          bigint unsigned auto_increment,
  title       varchar(255) not null,
  description varchar(255),

  primary key(id)
)engine=innodb;