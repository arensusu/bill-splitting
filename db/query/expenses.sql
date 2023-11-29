-- name: CreateExpense :one
INSERT INTO expenses (group_id, payer_id, amount, description, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetExpense :one
SELECT *
FROM expenses
WHERE id = $1;

-- name: ListGroupExpenses :many
SELECT *
FROM expenses
WHERE group_id = $1;

-- name: ListNonSettledGroupExpenses :many
SELECT *
FROM expenses
WHERE group_id = $1 AND is_settled = false;

-- name: UpdateExpense :one
UPDATE expenses
SET group_id = $2, payer_id = $3, amount = $4, description = $5, date = $6, is_settled = $7
WHERE id = $1
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;
