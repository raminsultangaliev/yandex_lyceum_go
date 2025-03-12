package main

import (
	"fmt"
	"time"
)

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Address:", p.address)
}

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Println("Employee:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Total Salary: %.2f", e.salary+e.bonus)
}

// type Student struct {
// 	name            string
// 	solvedProblems  int
// 	scoreForOneTask float64
// 	passingScore    float64
// }

// func (s Student) IsExcellentStudent() bool {
// 	score := s.scoreForOneTask * float64(s.solvedProblems)
// 	return score >= s.passingScore
// }

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() bool {
	return time.Now().After(t.deadline)
}

func (t Task) IsTopPriority() bool {
	return t.priority >= 4
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (t ToDoList) TasksCount() int {
	return len(t.tasks)
}

func (t ToDoList) NotesCount() int {
	return len(t.notes)
}

func (t ToDoList) CountTopPrioritiesTasks() int {
	res := 0
	for _, task := range t.tasks {
		if task.IsTopPriority() {
			res++
		}
	}
	return res
}

func (t ToDoList) CountOverdueTasks() int {
	res := 0
	for _, task := range t.tasks {
		if task.IsOverdue() {
			res++
		}
	}
	return res
}