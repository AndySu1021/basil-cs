// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: report_daily_staff.sql

package model

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

const listReportDailyStaff = `-- name: ListReportDailyStaff :many
select r.id, r.date, r.staff_id, r.total_member, r.total_scored_member, r.average_score, s.name as staff_name
from report_daily_staff r
         inner join staff s on s.id = r.staff_id
where ` + "`" + `date` + "`" + ` between ? and ?
order by ` + "`" + `date` + "`" + ` desc
`

type ListReportDailyStaffParams struct {
	Date   time.Time `db:"date" json:"date"`
	Date_2 time.Time `db:"date_2" json:"date_2"`
}

type ListReportDailyStaffRow struct {
	ID                int64           `db:"id" json:"id"`
	Date              time.Time       `db:"date" json:"date"`
	StaffID           int64           `db:"staff_id" json:"staff_id"`
	TotalMember       int32           `db:"total_member" json:"total_member"`
	TotalScoredMember int32           `db:"total_scored_member" json:"total_scored_member"`
	AverageScore      decimal.Decimal `db:"average_score" json:"average_score"`
	StaffName         string          `db:"staff_name" json:"staff_name"`
}

func (q *Queries) ListReportDailyStaff(ctx context.Context, arg ListReportDailyStaffParams) ([]ListReportDailyStaffRow, error) {
	rows, err := q.query(ctx, q.listReportDailyStaffStmt, listReportDailyStaff, arg.Date, arg.Date_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListReportDailyStaffRow{}
	for rows.Next() {
		var i ListReportDailyStaffRow
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.StaffID,
			&i.TotalMember,
			&i.TotalScoredMember,
			&i.AverageScore,
			&i.StaffName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
