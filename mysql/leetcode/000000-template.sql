
SET GLOBAL log_bin_trust_function_creators=1;

DROP FUNCTION IF EXISTS getNthHighestSalary
DELIMITER $$

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
  SELECT e.salary 
	FROM (SELECT salary, DENSE_RANK() OVER(order by salary DESC) rn
		  FROM Employee) e
	WHERE e.rn = N LIMIT 0,1
  );
END
$$

DELIMITER ;

SELECT getNthHighestSalary(2)