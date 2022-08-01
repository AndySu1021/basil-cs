package report

import (
	"context"
	"database/sql"
	"github.com/AndySu1021/go-util/logger"
	"github.com/Masterminds/squirrel"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type DailyStaff struct {
	db *sql.DB
}

type TmpDailyStaffRow struct {
	Date              time.Time
	StaffID           int64
	TotalMember       int32
	TotalScoredMember int32
	TotalScore        int32
}

func (d DailyStaff) Run() {
	// log execution time
	defer func(t time.Time) {
		logger.Logger.Infow("ExecReportDailyStaff", "latency", time.Since(t).Milliseconds())
	}(time.Now())

	ctx := context.Background()
	startTime, endTime, err := GetTimeRange("")
	if err != nil {
		logger.Logger.Errorf("get time range error: %s", err)
		return
	}

	result, err := d.getDailyStaffData(ctx, startTime, endTime)
	if err != nil {
		logger.Logger.Errorf("get report daily staff data error: %s", err)
		return
	}

	if err = d.upsertDailyStaffData(ctx, result); err != nil {
		logger.Logger.Errorf("upsert report daily staff error: %s", err)
		return
	}

	return
}

func (d DailyStaff) getDailyStaffData(ctx context.Context, startTime, endTime time.Time) (result []TmpDailyStaffRow, err error) {
	result = make([]TmpDailyStaffRow, 0)
	rows, err := squirrel.Select("staff_id", "count(*) AS total_member", "SUM(IF(score = 0, 0, 1)) AS total_scored_member", "sum(score) as total_score").
		From("room").
		Where(squirrel.GtOrEq{"accepted_at": startTime}).
		Where(squirrel.Lt{"accepted_at": endTime}).
		GroupBy("staff_id").
		RunWith(d.db).
		QueryContext(ctx)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var i TmpDailyStaffRow
		if err = rows.Scan(&i.StaffID, &i.TotalMember, &i.TotalScoredMember, &i.TotalScore); err != nil {
			return
		}
		i.Date = startTime
		result = append(result, i)
	}
	if err = rows.Close(); err != nil {
		return
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

func (d DailyStaff) upsertDailyStaffData(ctx context.Context, tmpRows []TmpDailyStaffRow) (err error) {
	if len(tmpRows) == 0 {
		return
	}

	builder := squirrel.Insert("report_daily_staff").
		Columns("date", "staff_id", "total_member", "total_scored_member", "average_score")

	for _, row := range tmpRows {
		averageScore := float64(0)
		if row.TotalScoredMember != 0 {
			averageScore = float64(row.TotalScore) / float64(row.TotalScoredMember)
		}
		builder = builder.Values(
			row.Date,
			row.StaffID,
			row.TotalMember,
			row.TotalScoredMember,
			decimal.NewFromFloat(averageScore),
		)
	}

	tmp := []string{
		"total_member = VALUES (total_member)",
		"total_scored_member = VALUES (total_scored_member)",
		"average_score = VALUES (average_score)",
	}

	_, err = builder.SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE " + strings.Join(tmp, ","))).
		RunWith(d.db).
		ExecContext(ctx)
	if err != nil {
		return
	}

	return
}
