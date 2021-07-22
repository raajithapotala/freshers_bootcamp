package main

import "fmt"

type salaryCalculator interface {
	getSalary() int
}
type Employee struct {
	employe  string
	basicpay int
	duration int
}

func (M Employee) getSalary() int {
	return M.basicpay * M.duration
}
func main() {
	FullTime := Employee{"Fulltime_Employee", 500, 28}
	Contract := Employee{"Contract_Employee", 100, 28}
	Freelancer := Employee{"Freelancer", 10, 240}
	fmt.Println("Salary of Fulltime Employee per month is ", FullTime.getSalary())
	fmt.Println("Salary of Contract Employee per month is ", Contract.getSalary())
	fmt.Println("Salary of Freelancer is ", Freelancer.getSalary())

}
