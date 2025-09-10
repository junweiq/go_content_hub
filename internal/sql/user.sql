-- 檢查用戶是否存在
-- name: CheckExist :one
SELECT EXISTS (
    SELECT 1
    FROM t_user_detail
    WHERE username = ?
) AS `exists`; -- 用反引號包裹別名

-- 創建用戶
-- name: CreateUser :exec
INSERT INTO t_user_detail (username, password, nickname, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- 根據用戶名查詢用戶信息
-- name: GetUserByUsername :one
SELECT id, username, password, nickname, created_at, updated_at
FROM t_user_detail
WHERE username = ?;