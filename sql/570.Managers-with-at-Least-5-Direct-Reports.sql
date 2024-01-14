-- Write your PostgreSQL query statement below
SELECT name
FROM Employee
WHERE id IN (
    SELECT managerId
    FROM Employee
    GROUP BY managerId
    HAVING COUNT(managerId) >= 5
)
-- Faster Join
SELECT name
FROM Employee t1
JOIN (
    SELECT managerId
    FROM Employee
    GROUP BY managerId
    HAVING COUNT(managerId) >= 5) AS t2
ON t1.id = t2.managerId