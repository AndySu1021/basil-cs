CREATE TABLE IF NOT EXISTS staff
(
    id              bigint unsigned auto_increment primary key,
    role_id         bigint unsigned         not null comment '角色ID',
    name            varchar(20)             not null comment '職員姓名',
    username        varchar(20)             not null comment '用戶名',
    password        char(32)                not null comment '密碼',
    avatar          varchar(255) default '' not null comment '大頭貼',
    status          tinyint      default 1  not null comment '職員狀態 1開啟 2關閉',
    serving_status  tinyint      default 1  not null comment '職員服務狀態 1關閉 2服務中 3閒置',
    online_status   tinyint      default 2  not null comment '職員線上狀態 1在線 2離線',
    last_login_time datetime                null comment '上次登入時間',
    created_by      bigint unsigned         not null comment '創建管理員',
    created_at      datetime                not null comment '創建時間',
    updated_by      bigint unsigned         not null comment '更新管理員',
    updated_at      datetime                not null comment '更新時間',
    constraint idx_username
        unique (username)
) COMMENT ='職員資料表' COLLATE utf8mb4_general_ci;
