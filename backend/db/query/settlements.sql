-- name: CreateSettlement :one
INSERT INTO settlements (payer_id, payee_id, amount)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetSettlement :one
SELECT *
FROM settlements
WHERE payer_id = $1 AND payee_id = $2;

-- name: ListSettlements :many
SELECT payer_id, payee_id, amount
FROM settlements, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE payer_id = members.id OR payee_id = members.id;

-- name: UpdateSettlement :one
UPDATE settlements
SET amount = $3
WHERE payer_id = $1 AND payee_id = $2
RETURNING *;

-- name: DeleteSettlement :exec
DELETE FROM settlements
WHERE payer_id = $1 AND payee_id = $2;
