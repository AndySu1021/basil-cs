package repository

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"reflect"
)

func (r *repository) WithTx(tx *sql.Tx) model.Querier {
	return r.Queries.WithTx(tx)
}

func (r *repository) Transaction(ctx context.Context, cb iface.Callback) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err = cb(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func (r *repository) ListTag(ctx context.Context, params types.ListTagParams) (tags []types.ListTagRow, total int64, err error) {
	var (
		tableName = "tag"
		columns   = []string{"id", "name", "status"}
	)

	dQuery := squirrel.Select(columns...).From(tableName)
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.Name != "" {
		dQuery = dQuery.Where(squirrel.Like{"name": "%" + params.Name + "%"})
		cQuery = cQuery.Where(squirrel.Like{"name": "%" + params.Name + "%"})
	}
	if params.Status != types.StatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"status": params.Status})
		cQuery = cQuery.Where(squirrel.Eq{"status": params.Status})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &tags); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListStaff(ctx context.Context, params types.ListStaffParams) (staffs []types.ListStaffRow, total int64, err error) {
	var (
		tableName = "staff"
		columns   = []string{
			"staff.id",
			"staff.name",
			"staff.username",
			"staff.status",
			"staff.serving_status",
			"role.name AS role_name",
		}
	)

	dQuery := squirrel.Select(columns...).From(tableName).
		InnerJoin("role on role.id = staff.role_id").
		Where(squirrel.Gt{"staff.id": 1})
	cQuery := squirrel.Select("count(*) AS count").From(tableName).Where(squirrel.Gt{"staff.id": 1})
	if params.Name != "" {
		dQuery = dQuery.Where(squirrel.Like{"staff.name": "%" + params.Name + "%"})
		cQuery = cQuery.Where(squirrel.Like{"staff.name": "%" + params.Name + "%"})
	}
	if params.Status != types.StatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"status": params.Status})
		cQuery = cQuery.Where(squirrel.Eq{"status": params.Status})
	}
	if params.ServingStatus != types.StaffServingStatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"serving_status": params.ServingStatus})
		cQuery = cQuery.Where(squirrel.Eq{"serving_status": params.ServingStatus})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &staffs); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListRoom(ctx context.Context, params types.ListRoomParams) (rooms []types.ListRoomRow, total int64, err error) {
	var (
		tableName = "room"
		columns   = []string{
			"room.id",
			"room.status",
			"room.created_at",
			"room.closed_at",
			"COALESCE(staff.name, '') as staff_name",
			"member.name as member_name",
			"COALESCE(tag.name, '') as tag_name",
		}
	)

	dQuery := squirrel.Select(columns...).From(tableName).
		LeftJoin("staff on staff.id = room.staff_id").
		LeftJoin("tag on tag.id = room.tag_id").
		InnerJoin("member on member.id = room.member_id")
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.RoomID != 0 {
		dQuery = dQuery.Where(squirrel.Eq{"room.id": params.RoomID})
		cQuery = cQuery.Where(squirrel.Eq{"room.id": params.RoomID})
	}
	if params.StaffID != 0 {
		dQuery = dQuery.Where(squirrel.Eq{"staff_id": params.StaffID})
		cQuery = cQuery.Where(squirrel.Eq{"staff_id": params.StaffID})
	}
	if params.Status != types.RoomStatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"room.status": params.Status})
		cQuery = cQuery.Where(squirrel.Eq{"room.status": params.StaffID})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &rooms); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListRole(ctx context.Context, params types.ListRoleParams) (roles []types.ListRoleRow, total int64, err error) {
	var (
		tableName = "role"
		columns   = []string{"id", "name"}
	)

	dQuery := squirrel.Select(columns...).From(tableName).Where(squirrel.Gt{"id": 1})
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.Name != "" {
		dQuery = dQuery.Where(squirrel.Like{"name": "%" + params.Name + "%"})
		cQuery = cQuery.Where(squirrel.Like{"name": "%" + params.Name + "%"})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &roles); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListRemind(ctx context.Context, params types.ListRemindParams) (reminds []types.ListRemindRow, total int64, err error) {
	var (
		tableName = "remind"
		columns   = []string{"id", "title", "content", "status"}
	)

	dQuery := squirrel.Select(columns...).From(tableName).OrderBy("updated_at desc")
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.Content != "" {
		dQuery = dQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
		cQuery = cQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
	}
	if params.Status != types.StatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"status": params.Status})
		cQuery = cQuery.Where(squirrel.Eq{"status": params.Status})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &reminds); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListNotice(ctx context.Context, params types.ListNoticeParams) (notices []types.ListNoticeRow, total int64, err error) {
	var (
		tableName = "notice"
		columns   = []string{"id", "title", "content", "start_at", "end_at", "status"}
	)

	dQuery := squirrel.Select(columns...).From(tableName)
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.Content != "" {
		dQuery = dQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
		cQuery = cQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
	}
	if params.Status != types.StatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"status": params.Status})
		cQuery = cQuery.Where(squirrel.Eq{"status": params.Status})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &notices); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListMember(ctx context.Context, params types.ListMemberParams) (members []types.ListMemberRow, total int64, err error) {
	var (
		tableName = "member"
		columns   = []string{
			"member.id",
			"type",
			"member.name",
			"real_name",
			"mobile",
			"email",
			"status",
			"online_status",
			"member.created_at",
			"site.name as site_name",
		}
	)

	dQuery := squirrel.Select(columns...).From(tableName).InnerJoin("site on site.id = member.site_id")
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.Mobile != "" {
		dQuery = dQuery.Where(squirrel.Like{"mobile": "%" + params.Mobile + "%"})
		cQuery = cQuery.Where(squirrel.Like{"mobile": "%" + params.Mobile + "%"})
	}
	if params.Email != "" {
		dQuery = dQuery.Where(squirrel.Like{"email": "%" + params.Email + "%"})
		cQuery = cQuery.Where(squirrel.Like{"email": "%" + params.Email + "%"})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &members); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListFastReply(ctx context.Context, params types.ListFastReplyParams) (replies []types.ListFastReplyRow, total int64, err error) {
	var (
		tableName = "fast_reply"
		columns   = []string{"fast_reply.id", "category_id", "title", "content", "status", "constant.value AS category"}
	)

	dQuery := squirrel.Select(columns...).From(tableName).InnerJoin("constant on constant.id = fast_reply.category_id")
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.Title != "" {
		dQuery = dQuery.Where(squirrel.Like{"title": "%" + params.Title + "%"})
		cQuery = cQuery.Where(squirrel.Like{"title": "%" + params.Title + "%"})
	}
	if params.Content != "" {
		dQuery = dQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
		cQuery = cQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
	}
	if params.Status != types.StatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"status": params.Status})
		cQuery = cQuery.Where(squirrel.Eq{"status": params.Status})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &replies); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListFAQ(ctx context.Context, params types.ListFAQParams) (faqs []types.ListFAQRow, total int64, err error) {
	var (
		tableName = "faq"
		columns   = []string{"id", "question", "answer", "status"}
	)

	dQuery := squirrel.Select(columns...).From(tableName)
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.Question != "" {
		dQuery = dQuery.Where(squirrel.Like{"question": "%" + params.Question + "%"})
		cQuery = cQuery.Where(squirrel.Like{"question": "%" + params.Question + "%"})
	}
	if params.Status != types.StatusAll {
		dQuery = dQuery.Where(squirrel.Eq{"status": params.Status})
		cQuery = cQuery.Where(squirrel.Eq{"status": params.Status})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &faqs); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) ListMessage(ctx context.Context, params types.ListMessageParams) (messages []types.ListMessageRow, total int64, err error) {
	var (
		tableName = "message"
		columns   = []string{"id", "sender_type", "sender_name", "content_type", "content", "ts"}
	)

	dQuery := squirrel.Select(columns...).From(tableName)
	cQuery := squirrel.Select("count(*) AS count").From(tableName)
	if params.RoomID != 0 {
		dQuery = dQuery.Where(squirrel.Eq{"room_id": params.RoomID})
		cQuery = cQuery.Where(squirrel.Eq{"room_id": params.RoomID})
	}
	if params.StaffID != 0 {
		dQuery = dQuery.Where(squirrel.Eq{"sender_type": types.MessageSenderTypeStaff, "sender_id": params.StaffID})
		cQuery = cQuery.Where(squirrel.Eq{"sender_type": types.MessageSenderTypeStaff, "sender_id": params.StaffID})
	}
	if params.Content != "" {
		dQuery = dQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
		cQuery = cQuery.Where(squirrel.Like{"content": "%" + params.Content + "%"})
	}

	if err = r.pageQuery(ctx, dQuery, params.Pagination, &messages); err != nil {
		return
	}
	if err = r.totalQuery(ctx, cQuery, &total); err != nil {
		return
	}

	return
}

func (r *repository) pageQuery(ctx context.Context, builder squirrel.SelectBuilder, pagination types.Pagination, data interface{}) (err error) {
	rows, err := builder.Limit(uint64(pagination.PageSize)).Offset(getOffset(pagination)).RunWith(r.db).QueryContext(ctx)
	if err != nil {
		return
	}
	defer rows.Close()

	// data: pointer to nil slice
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)

	for rows.Next() {
		newVal := reflect.New(v.Type().Elem()).Elem()
		fields := make([]interface{}, newVal.NumField())
		for i := 0; i < newVal.NumField(); i++ {
			fields[i] = newVal.Field(i).Addr().Interface()
		}
		if err = rows.Scan(fields...); err != nil {
			return
		}
		v.Set(reflect.Append(v, newVal))
	}
	if err = rows.Close(); err != nil {
		return
	}
	if err = rows.Err(); err != nil {
		return
	}

	if v.IsNil() {
		s := reflect.MakeSlice(v.Type(), 0, 0)
		v.Set(s)
	}

	return
}

func (r *repository) totalQuery(ctx context.Context, builder squirrel.SelectBuilder, total *int64) error {
	row := builder.RunWith(r.db).QueryRowContext(ctx)
	if err := row.Scan(total); err != nil {
		return err
	}
	return nil
}

func getOffset(params types.Pagination) uint64 {
	if params.Page > 0 {
		return uint64((params.Page - 1) * params.PageSize)
	}
	return 0
}
