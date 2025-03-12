package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func NewUser(id int, name string, email string, age int) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func CreateReport(user User, reportDate string) *Report {
	return &Report{
		User:     	user,
		ReportID: int(time.Now().Unix()),
		Date:     reportDate,
	}
}

func PrintReport(report Report) {
	fmt.Println(report.User.Name)
	fmt.Println(report.Date)
}

func GenerateUserReports(users []User, reportDate string) []Report {
	result := make([]Report, len(users))
	for i, user := range users {
		result[i] = *CreateReport(user, reportDate)
	}
	return result
}