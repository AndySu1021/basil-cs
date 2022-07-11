CREATE TABLE IF NOT EXISTS tag
(
    id         bigint unsigned auto_increment primary key,
    name       varchar(20)       not null comment '標籤名稱',
    status     tinyint default 1 not null comment '標籤狀態 1開啟 2關閉',
    created_by bigint unsigned   not null comment '創建管理員',
    created_at datetime          not null comment '創建時間',
    updated_by bigint unsigned   not null comment '更新管理員',
    updated_at datetime          not null comment '更新時間'
) COMMENT ='標籤資料表' COLLATE utf8mb4_general_ci;
