// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: authorization_rules.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteAuthorizationRule = `-- name: DeleteAuthorizationRule :one
DELETE FROM authorization_rules
WHERE rule_id = $1
RETURNING rule_id, role_id, is_administrator, permission_code, is_allowed, created_at
`

func (q *Queries) DeleteAuthorizationRule(ctx context.Context, ruleID int64) (AuthorizationRule, error) {
	row := q.db.QueryRow(ctx, deleteAuthorizationRule, ruleID)
	var i AuthorizationRule
	err := row.Scan(
		&i.RuleID,
		&i.RoleID,
		&i.IsAdministrator,
		&i.PermissionCode,
		&i.IsAllowed,
		&i.CreatedAt,
	)
	return i, err
}

const getAuthorizationRule = `-- name: GetAuthorizationRule :one
SELECT rule_id, role_id, is_administrator, permission_code, is_allowed, created_at
FROM authorization_rules
WHERE rule_id = $1
FOR NO KEY UPDATE
`

func (q *Queries) GetAuthorizationRule(ctx context.Context, ruleID int64) (AuthorizationRule, error) {
	row := q.db.QueryRow(ctx, getAuthorizationRule, ruleID)
	var i AuthorizationRule
	err := row.Scan(
		&i.RuleID,
		&i.RoleID,
		&i.IsAdministrator,
		&i.PermissionCode,
		&i.IsAllowed,
		&i.CreatedAt,
	)
	return i, err
}

const getAuthorizationRuleByRole = `-- name: GetAuthorizationRuleByRole :many
SELECT rule_id, role_id, is_administrator, permission_code, is_allowed, created_at
FROM authorization_rules
WHERE role_id = $1
FOR NO KEY UPDATE
`

func (q *Queries) GetAuthorizationRuleByRole(ctx context.Context, roleID int64) ([]AuthorizationRule, error) {
	rows, err := q.db.Query(ctx, getAuthorizationRuleByRole, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AuthorizationRule{}
	for rows.Next() {
		var i AuthorizationRule
		if err := rows.Scan(
			&i.RuleID,
			&i.RoleID,
			&i.IsAdministrator,
			&i.PermissionCode,
			&i.IsAllowed,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAuthorizationRule = `-- name: InsertAuthorizationRule :one
INSERT
INTO authorization_rules
    (role_id, is_administrator, permission_code, is_allowed)
VALUES
    ($1, $2, $3, $4)
RETURNING rule_id, role_id, is_administrator, permission_code, is_allowed, created_at
`

type InsertAuthorizationRuleParams struct {
	RoleID          int64  `json:"role_id"`
	IsAdministrator bool   `json:"is_administrator"`
	PermissionCode  string `json:"permission_code"`
	IsAllowed       bool   `json:"is_allowed"`
}

func (q *Queries) InsertAuthorizationRule(ctx context.Context, arg *InsertAuthorizationRuleParams) (AuthorizationRule, error) {
	row := q.db.QueryRow(ctx, insertAuthorizationRule,
		arg.RoleID,
		arg.IsAdministrator,
		arg.PermissionCode,
		arg.IsAllowed,
	)
	var i AuthorizationRule
	err := row.Scan(
		&i.RuleID,
		&i.RoleID,
		&i.IsAdministrator,
		&i.PermissionCode,
		&i.IsAllowed,
		&i.CreatedAt,
	)
	return i, err
}

const insertMultipleAuthorizationRules = `-- name: InsertMultipleAuthorizationRules :many
INSERT
INTO authorization_rules
    (role_id, is_administrator, permission_code, is_allowed)
VALUES
    (UNNEST($1::BIGINT[]), UNNEST($2::BOOLEAN[]), UNNEST($3::varchar(255)[]), UNNEST($4::BOOLEAN[]))
RETURNING rule_id, role_id, is_administrator, permission_code, is_allowed, created_at
`

type InsertMultipleAuthorizationRulesParams struct {
	Column1 []int64  `json:"column_1"`
	Column2 []bool   `json:"column_2"`
	Column3 []string `json:"column_3"`
	Column4 []bool   `json:"column_4"`
}

func (q *Queries) InsertMultipleAuthorizationRules(ctx context.Context, arg *InsertMultipleAuthorizationRulesParams) ([]AuthorizationRule, error) {
	rows, err := q.db.Query(ctx, insertMultipleAuthorizationRules,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AuthorizationRule{}
	for rows.Next() {
		var i AuthorizationRule
		if err := rows.Scan(
			&i.RuleID,
			&i.RoleID,
			&i.IsAdministrator,
			&i.PermissionCode,
			&i.IsAllowed,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isAllowed = `-- name: IsAllowed :one
SELECT is_allowed
FROM authorization_rules
WHERE role_id = $1 AND permission_code = $2
FOR NO KEY UPDATE
`

type IsAllowedParams struct {
	RoleID         int64  `json:"role_id"`
	PermissionCode string `json:"permission_code"`
}

func (q *Queries) IsAllowed(ctx context.Context, arg *IsAllowedParams) (bool, error) {
	row := q.db.QueryRow(ctx, isAllowed, arg.RoleID, arg.PermissionCode)
	var is_allowed bool
	err := row.Scan(&is_allowed)
	return is_allowed, err
}

const updateAuthorizationRule = `-- name: UpdateAuthorizationRule :one
UPDATE authorization_rules
SET
    role_id = COALESCE($2, role_id),
    is_administrator = COALESCE($3, is_administrator),
    permission_code = COALESCE($4, permission_code),
    is_allowed = COALESCE($5, is_allowed)
WHERE rule_id = $1
RETURNING rule_id, role_id, is_administrator, permission_code, is_allowed, created_at
`

type UpdateAuthorizationRuleParams struct {
	RuleID          int64       `json:"rule_id"`
	RoleID          pgtype.Int8 `json:"role_id"`
	IsAdministrator pgtype.Bool `json:"is_administrator"`
	PermissionCode  pgtype.Text `json:"permission_code"`
	IsAllowed       pgtype.Bool `json:"is_allowed"`
}

func (q *Queries) UpdateAuthorizationRule(ctx context.Context, arg *UpdateAuthorizationRuleParams) (AuthorizationRule, error) {
	row := q.db.QueryRow(ctx, updateAuthorizationRule,
		arg.RuleID,
		arg.RoleID,
		arg.IsAdministrator,
		arg.PermissionCode,
		arg.IsAllowed,
	)
	var i AuthorizationRule
	err := row.Scan(
		&i.RuleID,
		&i.RoleID,
		&i.IsAdministrator,
		&i.PermissionCode,
		&i.IsAllowed,
		&i.CreatedAt,
	)
	return i, err
}
