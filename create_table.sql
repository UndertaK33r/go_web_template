create database file_s;
use file_s;
create table `tbl_file`(
   `id` int(11) primary key not null auto_increment,
   `file_shal` char(40) not null default '' comment '文件hash',
   `file_name` varchar(255) not null default '' comment '文件名',
   `file_size` bigint(20) default 0 comment '文件大小',
   `file_addr` varchar(1024) not null default '' comment '文件路径',
   `create_at` timestamp not null default current_timestamp comment '创建时间',
   `update_at` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
   `status` tinyint(1) default 1 comment '状态 1正常 0删除',
   `ext_1` varchar(255) not null default '' comment '扩展字段1',
   `ext_2` varchar(255) not null default '' comment '扩展字段2',
   unique key `idx_file_hash` (`file_shal`),
   key `idx_status` (`status`)
)engine = innodb default charset = utf8mb4;
