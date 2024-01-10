-- Slower with sub-queries
SELECT id
FROM Weather w1
WHERE temperature > (
    SELECT temperature
    FROM Weather w2
    WHERE (w1.recordDate - w2.recordDate) = 1
)
-- Fast with a join?
SELECT w1.id
FROM Weather w1
LEFT JOIN Weather w2 ON w1.recordDate - w2.recordDate = 1
WHERE w1.temperature > w2.temperature