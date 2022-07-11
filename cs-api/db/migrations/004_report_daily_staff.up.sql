create table if not exists report_daily_staff
(
    id                  bigint unsigned auto_increment primary key,
    date                date            not null comment '報表日期',
    staff_id            bigint unsigned not null comment '職員ID',
    total_member        int             not null comment '總接待人數',
    total_scored_member int             not null comment '總評分人數',
    average_score       decimal(5, 2)   not null comment '平均分數',
    constraint idx_date_staff
        unique (date, staff_id)
) COMMENT ='每日客服報表' COLLATE utf8mb4_general_ci;

