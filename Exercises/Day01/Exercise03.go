package main

import "fmt"

/*
*

There are 3 types of employees in an organization.,
Fulltime
Contractor
Freelancer

Full-time employees are paid by the day and the Contractors are paid by the day.
Freelancers are paid by the hour.

Example:

Full-time employee gets paid 15000 per month
The contractor gets paid 3000 per month
Freelancer gets paid 2000 (if freelancer works 20 hours)

Define an interface that calculates the salary of an employee.
*/
type Employee interface {
	getSalary() int
}

type FulltimeEmployee struct {
	NumOfDaysWorked int
	DailyWage       int
}

type ContractorEmployee struct {
	NumOfDaysWorked int
	DailyWage       int
}

type FreelancerEmployee struct {
	NumOfHoursWorked int
	HourlyWage       int
}

func (f *FulltimeEmployee) getSalary() int {
	return f.NumOfDaysWorked * f.DailyWage
}

func (c *ContractorEmployee) getSalary() int {
	return c.NumOfDaysWorked * c.DailyWage
}

func (f *FreelancerEmployee) getSalary() int {
	if f.NumOfHoursWorked < 10 {
		return 0
	}
	return f.NumOfHoursWorked * f.HourlyWage
}

func main() {
	fulltime := FulltimeEmployee{3, 15000 / 30}
	contractor := ContractorEmployee{30, 3000 / 30}
	freelancer := FreelancerEmployee{30, 2000 / 20}

	fmt.Println("FulltimeEmployee salary: ", fulltime.getSalary())
	fmt.Println("Contractor salary: ", contractor.getSalary())
	fmt.Println("FreelancerEmployee salary: ", freelancer.getSalary())
}
