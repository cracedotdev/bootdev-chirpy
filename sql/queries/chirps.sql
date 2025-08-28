-- name: CreateChirp :one
INSERT INTO chirps(id, created_at, updated_at, body, user_id)
VALUES ($1,
        NOW(),
        NOW(),
        $2,
        $3)
RETURNING *;

-- name: GetChirps :many
select *
from chirps
order by created_at asc;

-- name: GetChirp :one
select *
from chirps
where id = $1;

-- name: DeleteAllChirps :exec
DELETE
FROM chirps;
