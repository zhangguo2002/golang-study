-- name: CreateUser :one
INSERT INTO users(username,email,password)
VALUES ($1,$2,$3)
    RETURNING id,username,email,created,updated;

-- name: GetUser :one
SELECT id,username,email,created,updated
FROM users
WHERE id=$1;

-- name: ListUsers :many
SELECT id,username,email,created,updated
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CreateBlog :one
INSERT INTO blogs(title,content,user_id,created,updated)
VALUES ($1,$2,$3,$4,$5)
    RETURNING id,title,content,user_id,created,updated;

-- name: ListBlogs :many
SELECT id,title,content,user_id,created,updated
FROM blogs
ORDER BY id;

-- name: GetUserByUsernameOrEmail :one
SELECT id,username,email,created,updated,password
FROM users
WHERE username = $1 OR email = $1;

-- name: CreateUserProfile :one
INSERT INTO user_profiles(user_id,profile_image)
VALUES ($1,$2)
RETURNING id,user_id,profile_image,created,updated;

-- name: GetUserProfileByUserId :one
SELECT id,user_id,profile_image,created,updated
FROM user_profiles
WHERE user_id = $1;
