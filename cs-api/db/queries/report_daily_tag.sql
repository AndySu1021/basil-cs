-- name: ListReportDailyTag :many
select * from report_daily_tag where date between ? and ?;