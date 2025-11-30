-- 003_set_default_age.sql

-- age 가 NULL 인 유저들 기본값 세팅
UPDATE users
SET age = COALESCE(age, 20);
