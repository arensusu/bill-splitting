-- name: CreateSettlement :one
INSERT INTO settlements (group_id, payer_id, payee_id, amount)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSettlement :one
SELECT *
FROM settlements
WHERE group_id = $1 AND payer_id = $2 AND payee_id = $3;

-- name: ListSettlements :many
SELECT *
FROM settlements
WHERE group_id = $1;

-- name: UpdateSettlement :one
UPDATE settlements
SET amount = $4, is_confirmed = $5
WHERE group_id = $1 AND payer_id = $2 AND payee_id = $3
RETURNING *;

-- name: DeleteSettlement :exec
DELETE FROM settlements
WHERE group_id = $1 AND payer_id = $2 AND payee_id = $3;
