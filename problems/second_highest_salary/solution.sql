# Write your MySQL query statement below
select max(salary) as SecondHighestSalary from Employee , (select max(salary) as max_sal from Employee) as max_salary
where Employee.salary < max_salary.max_sal;