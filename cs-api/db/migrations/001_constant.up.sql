create table if not exists constant
(
    id         bigint unsigned auto_increment primary key,
    type       tinyint         not null comment '常數類型 1快捷回覆 2客服配置',
    `key`      varchar(255)    not null comment '鍵',
    value      varchar(255)    not null comment '值',
    created_by bigint unsigned not null comment '創建管理員',
    created_at datetime        not null comment '創建時間',
    updated_by bigint unsigned not null comment '更新管理員',
    updated_at datetime        not null comment '更新時間'
) comment '常數資料表' COLLATE utf8mb4_general_ci;