-- name: CreateUserExpense :one
INSERT INTO user_expenses (expense_id, user_id, share)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserExpense :one
SELECT *
FROM user_expenses
WHERE expense_id = $1 and user_id = $2;

-- name: UpdateUserExpense :one
UPDATE user_expenses
SET share = $3
WHERE expense_id = $1 and user_id = $2
RETURNING *;

-- name: DeleteUserExpense :exec
DELETE FROM user_expenses
WHERE expense_id = $1 and user_id = $2;
