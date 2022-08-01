package types

type ListRoleParams struct {
	Name string `form:"name" binding:""`
	Pagination
}

type ListRoleRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
