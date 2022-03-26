
-- https://leetcode-cn.com/problems/nth-highest-salary/

-- 
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N := N-1;
  RETURN (
      # Write your MySQL query statement below.
      SELECT Salary FROM Employee e group by e.Salary ORDER BY Salary DESC LIMIT N, 1
  );
END;

-- CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
  SELECT e.salary 
	FROM (SELECT salary, DENSE_RANK() OVER(order by salary DESC) rn
		  FROM Employee) e
	WHERE e.rn = N LIMIT 0,1
  );
END

--
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N:=N-1;
  RETURN (
  SELECT IFNULL((SELECT DISTINCT salary FROM Employee ORDER BY salary DESC LIMIT N, 1), NULL) 
  );
END


