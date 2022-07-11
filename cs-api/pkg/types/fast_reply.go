package types

type ListFastReplyParams struct {
	Title   string `form:"title" binding:""`
	Content string `form:"content" binding:""`
	Status  Status `form:"status" binding:"min=0,max=2"`
	Pagination
}

type ListFastReplyRow struct {
	ID         int64  `db:"id" json:"id"`
	CategoryID int64  `db:"category_id" json:"category_id"`
	Title      string `db:"title" json:"title"`
	Content    string `db:"content" json:"content"`
	Status     Status `db:"status" json:"status"`
	Category   string `db:"category" json:"category"`
}
