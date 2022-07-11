CREATE TABLE IF NOT EXISTS notice
(
    id         bigint unsigned auto_increment primary key,
    title      varchar(20)       not null comment '公告標題',
    content    varchar(255)      not null comment '公告內容',
    start_at   datetime          not null comment '開始時間',
    end_at     datetime          not null comment '結束時間',
    status     tinyint default 1 not null comment '狀態 1開啟 2關閉',
    created_by bigint unsigned   not null comment '創建管理員',
    created_at datetime          not null comment '創建時間',
    updated_by bigint unsigned   not null comment '更新管理員',
    updated_at datetime          not null comment '更新時間'
) COMMENT ='系統公告資料表' COLLATE utf8mb4_general_ci;
