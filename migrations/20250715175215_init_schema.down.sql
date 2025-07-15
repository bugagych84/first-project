-- Откат миграции (в обратном порядке)
DROP INDEX IF EXISTS idx_tasks_user_id;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;