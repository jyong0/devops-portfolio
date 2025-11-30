-- 002_add_age_to_users.sql

-- age 컬럼이 없으면 추가
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users'
          AND column_name = 'age'
    ) THEN
        ALTER TABLE users
        ADD COLUMN age INT;
    END IF;
END $$;
