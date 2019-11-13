package main

import "fmt"

func main() {
	var days int
	var sundays int
	fmt.Print("勤務日数を入力してください：")
	fmt.Scan(&days)
	fmt.Print("日曜日の日数を入力してください：")
	fmt.Scan(&sundays)
	annualIncome := 1030
	b := annualIncome + 30
	a := 1030 / 4 * 3
	c := b / 4 * 3
	workingHours := 5
	dailySalary := annualIncome*workingHours + a
	d := b*workingHours + c
	weekdays := days - sundays
	income := dailySalary * weekdays
	e := d*sundays + income
	fmt.Println(e)
}
