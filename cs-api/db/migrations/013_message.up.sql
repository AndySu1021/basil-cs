create table if not exists message
(
    id           bigint unsigned auto_increment primary key,
    room_id      bigint unsigned not null comment '房間ID',
    op_type      tinyint         not null default 0 comment '操作類型 0其他 1客服輸入中 2發送訊息',
    sender_type  tinyint         not null default 0 comment '發送人類型 0系統 1會員 2客服',
    sender_id    bigint unsigned not null default 0 comment '發送人ID',
    sender_name  varchar(255)    not null default '' comment '發送人名稱',
    content_type tinyint         not null default 0 comment '消息類型 0其他 1文字 2圖片',
    content      varchar(255)    not null default '' comment '訊息內容',
    extra        JSON            null comment '額外資訊',
    ts           bigint unsigned not null comment '創建時間戳',
    created_at   datetime        not null comment '創建時間'
) COMMENT ='歷史訊息紀錄表' COLLATE utf8mb4_general_ci;

create index idx_room_id
    on message (room_id);