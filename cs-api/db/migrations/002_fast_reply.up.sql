create table if not exists fast_reply
(
    id          bigint unsigned auto_increment primary key,
    category_id bigint unsigned   not null comment '分類ID(constantID)',
    title       varchar(255)      not null comment '訊息標題',
    content     varchar(255)      not null comment '訊息內容',
    status      tinyint default 1 not null comment '消息狀態 1開啟 2關閉',
    created_by  bigint unsigned   not null comment '創建管理員',
    created_at  datetime          not null comment '創建時間',
    updated_by  bigint unsigned   not null comment '更新管理員',
    updated_at  datetime          not null comment '更新時間'
) comment '快捷回覆資料表' COLLATE utf8mb4_general_ci;