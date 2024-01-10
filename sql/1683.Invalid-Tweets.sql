-- Write your PostgreSQL query statement below
-- We should use CHAR_LENGTH instead of LENGTH in case of non-ASCII characters in a string
-- LENGTH counts the length of a string in bytes
-- CHAR_LENGTH measures the length in of a string in characters
SELECT tweet_id
FROM Tweets
WHERE CHAR_LENGTH(content) > 15