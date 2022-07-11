-- name: TagSeeder :exec
INSERT
IGNORE INTO tag
(id, name, status, created_at, created_by, updated_at, updated_by)
VALUES
    (1, '其他', 1, now(), 0, now(), 0);

-- name: RoleSeeder :exec
INSERT IGNORE INTO role
(id, name, permissions, created_at, created_by, updated_at, updated_by)
VALUES
    (1, '超級管理員', '["*"]', now(), 0, now(), 0);

-- name: StaffSeeder :exec
INSERT IGNORE INTO staff
(id, role_id, name, username, password, created_at, created_by, updated_at, updated_by)
VALUES
    (1, 1, '超級管理員', 'admin', ?, now(), 0, now(), 0);

-- name: ConstantSeeder :exec
INSERT IGNORE INTO constant
(id, type, `key`, value, created_at, created_by, updated_at, updated_by)
VALUES
    (1, 2, 'CsConfig', '{"max_member":20,"member_pending_expire":5,"greeting_text":"您好，很高興為您服務"}', now(), 0, now(), 0);