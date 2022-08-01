package types

type StaffServingStatus int8

const (
	StaffServingStatusAll StaffServingStatus = iota
	StaffServingStatusClosed
	StaffServingStatusServing
	StaffServingStatusPending
)

type StaffOnlineStatus int8

const (
	StaffOnlineStatusOnline StaffOnlineStatus = iota + 1
	StaffOnlineStatusOffline
)

type ListStaffParams struct {
	Name          string             `form:"name" binding:""`
	Status        Status             `form:"status" binding:"min=0,max=2"`
	ServingStatus StaffServingStatus `form:"serving_status" binding:"min=0,max=3"`
	Pagination
}

type ListStaffRow struct {
	ID            int64              `db:"id" json:"id"`
	Name          string             `db:"name" json:"name"`
	Username      string             `db:"username" json:"username"`
	Status        Status             `db:"status" json:"status"`
	ServingStatus StaffServingStatus `db:"serving_status" json:"serving_status"`
	RoleName      string             `db:"role_name" json:"role_name"`
}
