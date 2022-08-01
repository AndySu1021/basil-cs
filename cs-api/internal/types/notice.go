package types

import "time"

type ListNoticeParams struct {
	Content string `form:"content" binding:""`
	Status  Status `form:"status" binding:"min=0,max=2"`
	Pagination
}

type ListNoticeRow struct {
	ID      int64     `db:"id" json:"id"`
	Title   string    `db:"title" json:"title"`
	Content string    `db:"content" json:"content"`
	StartAt time.Time `db:"start_at" json:"start_at"`
	EndAt   time.Time `db:"end_at" json:"end_at"`
	Status  Status    `db:"status" json:"status"`
}
