CREATE TABLE IF NOT EXISTS remind
(
    id         bigint unsigned auto_increment primary key,
    title      varchar(30)       not null comment '標題',
    content    varchar(255)      not null comment '內容',
    status     tinyint default 1 not null comment '狀態 1開啟 2關閉',
    created_by bigint unsigned   not null comment '創建管理員',
    created_at datetime          not null comment '創建時間',
    updated_by bigint unsigned   not null comment '更新管理員',
    updated_at datetime          not null comment '更新時間'
) COMMENT ='後台提醒事項資料表' COLLATE utf8mb4_general_ci;
