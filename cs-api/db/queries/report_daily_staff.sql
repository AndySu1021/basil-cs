-- name: ListReportDailyStaff :many
select r.*, s.name as staff_name
from report_daily_staff r
         inner join staff s on s.id = r.staff_id
where `date` between ? and ?
order by `date` desc;