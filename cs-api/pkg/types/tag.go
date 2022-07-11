package types

type ListTagParams struct {
	Name   string `form:"name" binding:""`
	Status Status `form:"status" binding:"min=0,max=2"`
	Pagination
}

type ListTagRow struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Status Status `db:"status" json:"status"`
}
