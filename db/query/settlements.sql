-- name: CreateSettlement :one
INSERT INTO settlements (payer_id, payee_id, amount, date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSettlement :one
SELECT *
FROM settlements
WHERE id = $1;

-- name: UpdateSettlement :one
UPDATE settlements
SET payer_id = $2, payee_id = $3, amount = $4, date = $5
WHERE id = $1
RETURNING *;

-- name: DeleteSettlement :exec
DELETE FROM settlements
WHERE id = $1;
