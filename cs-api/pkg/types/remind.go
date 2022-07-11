package types

type ListRemindParams struct {
	Content string `form:"content" binding:""`
	Status  Status `form:"status" binding:"min=0,max=2"`
	Pagination
}

type ListRemindRow struct {
	ID      int64  `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	Status  Status `db:"status" json:"status"`
}
