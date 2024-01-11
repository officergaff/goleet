-- Need to learn cross joins
SELECT s.student_id, s.student_name, c.subject_name, COUNT(e.subject_name) attended_exams
FROM Students s
CROSS JOIN Subjects c
LEFT JOIN Examinations e 
ON s.student_id = e.student_id AND c.subject_name = e.subject_name
GROUP BY s.student_id, c.subject_name, s.student_name
ORDER BY s.student_id, c.subject_name