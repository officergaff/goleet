-- Write your MySQL query statement below
SELECT EmployeeUNI.unique_id, Employees.name
FROM Employees LEFT JOIN EmployeeUNI 
ON EmployeeUNI.id = Employees.id
-- ALT
SELECT unique_id, name
FROM Employees e
LEFT JOIN EmployeeUNI u ON e.id = u.id