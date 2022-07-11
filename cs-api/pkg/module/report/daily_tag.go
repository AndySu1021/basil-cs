package report

import (
	"context"
	"cs-api/pkg/types"
	"database/sql"
	"github.com/AndySu1021/go-util/logger"
	"github.com/Masterminds/squirrel"
	"strings"
	"time"
)

type DailyTag struct {
	db *sql.DB
}

type TmpDailyTagRow struct {
	Date  time.Time
	TagID int64
	Count int32
}

func (d DailyTag) Run() {
	// log execution time
	defer func(t time.Time) {
		logger.Logger.Infow("ExecReportDailyTag", "latency", time.Since(t).Milliseconds())
	}(time.Now())

	ctx := context.Background()
	startTime, endTime, err := GetTimeRange("")
	if err != nil {
		logger.Logger.Errorf("get time range error: %s", err)
		return
	}

	result, err := d.getDailyTagData(ctx, startTime, endTime)
	if err != nil {
		logger.Logger.Errorf("get report daily tag data error: %s", err)
		return
	}

	if err = d.upsertDailyTagData(ctx, result); err != nil {
		logger.Logger.Errorf("upsert report daily tag error: %s", err)
		return
	}

	return
}

func (d DailyTag) getDailyTagData(ctx context.Context, startTime, endTime time.Time) (result []TmpDailyTagRow, err error) {
	result = make([]TmpDailyTagRow, 0)
	rows, err := squirrel.Select("tag_id", "count(*) AS count").
		From("room").
		Where(squirrel.GtOrEq{"closed_at": startTime}).
		Where(squirrel.Lt{"closed_at": endTime}).
		Where(squirrel.Eq{"status": types.RoomStatusClosed}).
		GroupBy("tag_id").
		RunWith(d.db).
		QueryContext(ctx)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var i TmpDailyTagRow
		if err = rows.Scan(&i.TagID, &i.Count); err != nil {
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

func (d DailyTag) upsertDailyTagData(ctx context.Context, tmpRows []TmpDailyTagRow) (err error) {
	if len(tmpRows) == 0 {
		return
	}

	builder := squirrel.Insert("report_daily_tag").
		Columns("date", "tag_id", "count")

	for _, row := range tmpRows {
		builder = builder.Values(
			row.Date,
			row.TagID,
			row.Count,
		)
	}

	tmp := []string{
		"count = VALUES (count)",
	}

	_, err = builder.SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE " + strings.Join(tmp, ","))).
		RunWith(d.db).
		ExecContext(ctx)
	if err != nil {
		return
	}

	return
}
