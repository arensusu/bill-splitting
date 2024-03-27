-- name: CreateExpense :one
INSERT INTO expenses (member_id, amount, description, date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetExpense :one
SELECT member_id, amount, description, date
FROM expenses
WHERE id = $1;

-- name: ListExpenses :many
SELECT *
FROM expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id;

-- name: ListNonSettledExpenses :many
SELECT *
FROM expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id AND is_settled = false;

-- name: UpdateExpense :one
UPDATE expenses
SET member_id = $2, amount = $3, description = $4, date = $5, is_settled = $6
WHERE id = $1
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;
