// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: others_accounts.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createOthersAccount = `-- name: CreateOthersAccount :exec
INSERT INTO others_accounts (
created_at,
updated_at,
account_name,
account_number,
sort_code,
account_id
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, account_name, account_number, sort_code, account_id
`

type CreateOthersAccountParams struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AccountName   string
	AccountNumber string
	SortCode      string
	AccountID     uuid.UUID
}

func (q *Queries) CreateOthersAccount(ctx context.Context, arg CreateOthersAccountParams) error {
	_, err := q.db.ExecContext(ctx, createOthersAccount,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.AccountName,
		arg.AccountNumber,
		arg.SortCode,
		arg.AccountID,
	)
	return err
}

const getOthersAccountByDetails = `-- name: GetOthersAccountByDetails :one
SELECT id, created_at, updated_at, account_name, account_number, sort_code, account_id FROM others_accounts WHERE account_id = $1 AND account_number = $2 AND sort_code = $3
`

type GetOthersAccountByDetailsParams struct {
	AccountID     uuid.UUID
	AccountNumber string
	SortCode      string
}

func (q *Queries) GetOthersAccountByDetails(ctx context.Context, arg GetOthersAccountByDetailsParams) (OthersAccount, error) {
	row := q.db.QueryRowContext(ctx, getOthersAccountByDetails, arg.AccountID, arg.AccountNumber, arg.SortCode)
	var i OthersAccount
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountName,
		&i.AccountNumber,
		&i.SortCode,
		&i.AccountID,
	)
	return i, err
}

const getOthersAccountByID = `-- name: GetOthersAccountByID :one
SELECT id, created_at, updated_at, account_name, account_number, sort_code, account_id FROM others_accounts WHERE account_id = $1 AND id = $2
`

type GetOthersAccountByIDParams struct {
	AccountID uuid.UUID
	ID        int32
}

func (q *Queries) GetOthersAccountByID(ctx context.Context, arg GetOthersAccountByIDParams) (OthersAccount, error) {
	row := q.db.QueryRowContext(ctx, getOthersAccountByID, arg.AccountID, arg.ID)
	var i OthersAccount
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountName,
		&i.AccountNumber,
		&i.SortCode,
		&i.AccountID,
	)
	return i, err
}

const getOthersAccounts = `-- name: GetOthersAccounts :many
SELECT id, created_at, updated_at, account_name, account_number, sort_code, account_id FROM others_accounts WHERE account_id = $1
`

func (q *Queries) GetOthersAccounts(ctx context.Context, accountID uuid.UUID) ([]OthersAccount, error) {
	rows, err := q.db.QueryContext(ctx, getOthersAccounts, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OthersAccount
	for rows.Next() {
		var i OthersAccount
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccountName,
			&i.AccountNumber,
			&i.SortCode,
			&i.AccountID,
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

const updateOthersAccountName = `-- name: UpdateOthersAccountName :exec
UPDATE others_accounts
SET account_name = $1
WHERE account_number = $2 AND sort_code = $3 AND account_id = $4
`

type UpdateOthersAccountNameParams struct {
	AccountName   string
	AccountNumber string
	SortCode      string
	AccountID     uuid.UUID
}

func (q *Queries) UpdateOthersAccountName(ctx context.Context, arg UpdateOthersAccountNameParams) error {
	_, err := q.db.ExecContext(ctx, updateOthersAccountName,
		arg.AccountName,
		arg.AccountNumber,
		arg.SortCode,
		arg.AccountID,
	)
	return err
}