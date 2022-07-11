create table if not exists room
(
    id          bigint unsigned auto_increment primary key,
    staff_id    bigint unsigned            not null comment '職員ID',
    member_id   bigint unsigned            not null comment '會員ID',
    ip          varchar(20)                not null comment '訪問IP',
    source      tinyint                    not null comment '來源 1-web 2-app',
    user_agent  varchar(255)    default '' not null comment '請求使用者代理',
    tag_id      bigint unsigned default 0  not null comment '標籤ID',
    score       tinyint         default 0  not null comment '評分 1-5',
    status      tinyint         default 1  not null comment '客服房狀態 1等待中 2排隊中 3服務中 4已關閉',
    accepted_at datetime                   null comment '開始諮詢時間',
    closed_at   datetime                   null comment '關閉時間',
    created_at  datetime                   not null comment '創建時間',
    updated_at  datetime                   not null comment '更新時間'
) comment '諮詢房紀錄表' COLLATE utf8mb4_general_ci;

create index idx_member_id
    on room (member_id);

create index idx_accepted_at
    on room (accepted_at);

create index idx_closed_at
    on room (closed_at);