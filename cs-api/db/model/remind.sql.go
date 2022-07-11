// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: remind.sql

package model

import (
	"context"
	"time"

	"cs-api/pkg/types"
)

const createRemind = `-- name: CreateRemind :exec
INSERT INTO remind (title, content, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateRemindParams struct {
	Title     string       `db:"title" json:"title"`
	Content   string       `db:"content" json:"content"`
	Status    types.Status `db:"status" json:"status"`
	CreatedBy int64        `db:"created_by" json:"created_by"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedBy int64        `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateRemind(ctx context.Context, arg CreateRemindParams) error {
	_, err := q.exec(ctx, q.createRemindStmt, createRemind,
		arg.Title,
		arg.Content,
		arg.Status,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	return err
}

const deleteRemind = `-- name: DeleteRemind :exec
DELETE
FROM remind
WHERE id = ?
`

func (q *Queries) DeleteRemind(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteRemindStmt, deleteRemind, id)
	return err
}

const getRemind = `-- name: GetRemind :one
SELECT id, title, content, status, created_by, created_at, updated_by, updated_at
FROM remind
WHERE id = ? LIMIT 1
`

func (q *Queries) GetRemind(ctx context.Context, id int64) (Remind, error) {
	row := q.queryRow(ctx, q.getRemindStmt, getRemind, id)
	var i Remind
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}

const listActiveRemind = `-- name: ListActiveRemind :many
SELECT title, content
FROM remind
WHERE status = 1
`

type ListActiveRemindRow struct {
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}

func (q *Queries) ListActiveRemind(ctx context.Context) ([]ListActiveRemindRow, error) {
	rows, err := q.query(ctx, q.listActiveRemindStmt, listActiveRemind)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListActiveRemindRow{}
	for rows.Next() {
		var i ListActiveRemindRow
		if err := rows.Scan(&i.Title, &i.Content); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRemind = `-- name: UpdateRemind :exec
UPDATE remind
SET title      = ?,
    content    = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateRemindParams struct {
	Title     string       `db:"title" json:"title"`
	Content   string       `db:"content" json:"content"`
	Status    types.Status `db:"status" json:"status"`
	UpdatedBy int64        `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	ID        int64        `db:"id" json:"id"`
}

func (q *Queries) UpdateRemind(ctx context.Context, arg UpdateRemindParams) error {
	_, err := q.exec(ctx, q.updateRemindStmt, updateRemind,
		arg.Title,
		arg.Content,
		arg.Status,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}