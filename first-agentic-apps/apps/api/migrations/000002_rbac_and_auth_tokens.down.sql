DROP INDEX IF EXISTS idx_auth_tokens_expires_at;
DROP INDEX IF EXISTS idx_auth_tokens_token_type;
DROP INDEX IF EXISTS idx_auth_tokens_user_id;
DROP INDEX IF EXISTS idx_permissions_module_id;
DROP INDEX IF EXISTS idx_role_permissions_permission_id;
DROP INDEX IF EXISTS idx_user_roles_role_id;

DROP TABLE IF EXISTS auth_tokens;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS modules;
DROP TABLE IF EXISTS roles;
