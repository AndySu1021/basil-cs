CREATE TABLE IF NOT EXISTS role
(
    id          bigint unsigned auto_increment primary key,
    name        varchar(20)     not null comment '角色名稱',
    permissions JSON            not null comment '角色權限',
    created_by  bigint unsigned not null comment '創建管理員',
    created_at  datetime        not null comment '創建時間',
    updated_by  bigint unsigned not null comment '更新管理員',
    updated_at  datetime        not null comment '更新時間'
) COMMENT ='權限角色資料表' COLLATE utf8mb4_general_ci;
