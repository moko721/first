package main

import "fmt"

func main() {
	var annualIncome int
	var days int
	var sundays int
	var additionalSalary int
	var workingHours int
	var workingMinute int
	fmt.Print("時給を入れてください:")
	fmt.Scan(&annualIncome)
	fmt.Print("勤務日数を入力してください：")
	fmt.Scan(&days)
	fmt.Print("日曜日の日数を入力してください：")
	fmt.Scan(&sundays)
	fmt.Print("時給から追加される金額を入力してください")
	fmt.Scan(&additionalSalary)
	fmt.Print("何時間働きますか：")
	fmt.Scan(&workingHours)
	fmt.Print("何分働きますか")
	fmt.Scan(&workingMinute)
	minuteSalary := annualIncome / 60 * workingMinute
	sundayHourlyWage := annualIncome + additionalSalary
	sundayMinuteSalary := sundayHourlyWage / 60 * workingMinute
	//hourlyWage := 1013 / 4 * 3
	//c := sundaySalary / 4 * 3
	//workingHoursOfTheDay := workingHours + workingHours
	dailySalary := annualIncome * workingHours + minuteSalary
	sundaySalary := sundayHourlyWage * workingHours + sundayMinuteSalary
	weekdays := days - sundays
	income := dailySalary * weekdays
	salaryForOneMonth := sundaySalary*sundays + income
	fmt.Println(salaryForOneMonth)
}
