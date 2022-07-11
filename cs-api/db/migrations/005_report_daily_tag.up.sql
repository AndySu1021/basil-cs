create table if not exists report_daily_tag
(
    id         bigint unsigned auto_increment primary key,
    date       date            not null comment '報表日期',
    tag_id     bigint unsigned not null comment '標籤ID',
    count      int             not null comment '人數',
    constraint idx_date_tag
        unique (date, tag_id)
) COMMENT ='每日標籤報表' COLLATE utf8mb4_general_ci;

