use work_thor;
alter table user add identity tinyint(4) default 0 not null comment '用户身份信息';
