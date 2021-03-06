// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: constant.sql

package model

import (
	"context"
	"time"
)

const checkFastReplyCategory = `-- name: CheckFastReplyCategory :one
SELECT 1
FROM constant
WHERE id = ? LIMIT 1
`

func (q *Queries) CheckFastReplyCategory(ctx context.Context, id int64) (interface{}, error) {
	row := q.queryRow(ctx, q.checkFastReplyCategoryStmt, checkFastReplyCategory, id)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const createFastReplyCategory = `-- name: CreateFastReplyCategory :exec
INSERT INTO constant (` + "`" + `type` + "`" + `, ` + "`" + `key` + "`" + `, value, created_by, created_at, updated_by, updated_at)
VALUES (1, 'FastReply', ?, ?, ?, ?, ?)
`

type CreateFastReplyCategoryParams struct {
	Value     string    `db:"value" json:"value"`
	CreatedBy int64     `db:"created_by" json:"created_by"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedBy int64     `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateFastReplyCategory(ctx context.Context, arg CreateFastReplyCategoryParams) error {
	_, err := q.exec(ctx, q.createFastReplyCategoryStmt, createFastReplyCategory,
		arg.Value,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	return err
}

const getCsConfig = `-- name: GetCsConfig :one
SELECT id, type, ` + "`" + `key` + "`" + `, value, created_by, created_at, updated_by, updated_at
FROM constant
WHERE ` + "`" + `type` + "`" + ` = 2
  and ` + "`" + `key` + "`" + ` = 'CsConfig' LIMIT 1
`

func (q *Queries) GetCsConfig(ctx context.Context) (Constant, error) {
	row := q.queryRow(ctx, q.getCsConfigStmt, getCsConfig)
	var i Constant
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Key,
		&i.Value,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}

const listFastReplyCategory = `-- name: ListFastReplyCategory :many
SELECT id, type, ` + "`" + `key` + "`" + `, value, created_by, created_at, updated_by, updated_at
FROM constant
WHERE ` + "`" + `type` + "`" + ` = 1
  and ` + "`" + `key` + "`" + ` = 'FastReply'
`

func (q *Queries) ListFastReplyCategory(ctx context.Context) ([]Constant, error) {
	rows, err := q.query(ctx, q.listFastReplyCategoryStmt, listFastReplyCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Constant{}
	for rows.Next() {
		var i Constant
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Key,
			&i.Value,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedBy,
			&i.UpdatedAt,
		); err != nil {
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

const updateCsConfig = `-- name: UpdateCsConfig :exec
UPDATE constant
SET value      = ?,
    updated_by = ?,
    updated_at = ?
WHERE ` + "`" + `type` + "`" + ` = 2
  and ` + "`" + `key` + "`" + ` = 'CsConfig'
`

type UpdateCsConfigParams struct {
	Value     string    `db:"value" json:"value"`
	UpdatedBy int64     `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (q *Queries) UpdateCsConfig(ctx context.Context, arg UpdateCsConfigParams) error {
	_, err := q.exec(ctx, q.updateCsConfigStmt, updateCsConfig, arg.Value, arg.UpdatedBy, arg.UpdatedAt)
	return err
}
