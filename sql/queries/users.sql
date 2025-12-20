-- name: CreateUser :one

INSERT INTO users (id, created_at, updated_at, name)
values(
  $1,
  $2,
  $3,
  $4
  )
  returning *;


-- name: GetUser :one

SELECT * FROM users
WHERE name = $1;

-- name: GetUserByID :one

SELECT * FROM users
WHERE ID = $1;

-- name: ResetUsers :exec

DELETE FROM users;


-- name: GetAllUsers :many

SELECT * FROM users;
