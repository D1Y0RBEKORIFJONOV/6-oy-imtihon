-- name: InsertInfo :one
INSERT INTO incomeexpenses (id,userid,type,category,currency,amount,date) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id;

-- name: Getinfo :many
SELECT * FROM incomeexpenses WHERE userid=$1;

