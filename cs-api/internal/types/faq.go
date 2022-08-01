package types

type ListFAQParams struct {
	Question string `form:"question" binding:""`
	Status   Status `form:"status" binding:"min=0,max=2"`
	Pagination
}

type ListFAQRow struct {
	ID       int64  `db:"id" json:"id"`
	Question string `db:"question" json:"question"`
	Answer   string `db:"answer" json:"answer"`
	Status   Status `db:"status" json:"status"`
}
