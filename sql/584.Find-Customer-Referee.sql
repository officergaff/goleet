-- Write your PostgreSQL query statement below
SELECT name
FROM Customer
WHERE COALESCE(referee_id, 0) != 2
-- ALT
SELECT name
FROM Customer
Where referee_id != 2 OR referee_id IS NOT NULL 