package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name   string
	Age    int
	Active bool
	Mass   float64
	Books  []string
}

func Users() []User {
	return []User{
		{
			"John Doe",
			30,
			true,
			80.0,
			[]string{"Harry Potter", "1984"},
		},
		{
			"Jake Doe",
			20,
			false,
			60.0,
			[]string{},
		},
		{
			" Jane Doe ",
			150,
			true,
			.75,
			[]string{"Harry Potter", "Game of Thrones"},
		},
		{
			"\t",
			-10,
			true,
			8000.0,
			[]string{"Harry Potter"},
		},
		{
			"Vm0weE5GVXhUblJWV0dSUFZtMW9WVll3WkRSV1ZteDBaRVYwVmsxWGVGWlZiVEZIWVd4S2MxTnNiR0ZXVm5Cb1ZsVmFWMVpWTVVWaGVqQTk=\nVm0weE5GVXhUblJWV0dSUFZtMW9WVll3WkRSV1ZteDBaRVYwVmsxWGVGWlZiVEZIWVd4S2MxTnNiR0ZXVm5Cb1ZsVmFWMVpWTVVWaGVqQTk=",
			0,
			true,
			0,
			[]string{"The Hunger Games"},
		},
		{
			"\x00\x10\x20\x30\x40\x50\x60\x70",
			0,
			true,
			0,
			[]string{"Moby Dick", "It", "The Green Mile"},
		},
	}
}

func main() {
	users := Users()

	const sourceLength = 10
	const dash = "--------------------------------------"
	const dash2 = "----------------"

	longestName := len("Name")

	bookAges := make(map[string][]int)

	for i, u := range users {
		if len(u.Name) > sourceLength {
			users[i].Name = u.Name[:7] + "..."
		}
		nameLen := len(fmt.Sprintf("%10.10q", u.Name))

		if nameLen > longestName {
			longestName = nameLen
		}

		for _, book := range u.Books {
			bookAges[book] = append(bookAges[book], u.Age)
		}
	}

	titleFormat := fmt.Sprintf("%%%ds | %%3s | %%6s | %%-8s | %%s\n", longestName)
	rowFormat := fmt.Sprintf("%%%d.%dq | %%3d | %%6s | %%6.1fkg | %%s\n", longestName, longestName-2)
	ozRowFormat := fmt.Sprintf("%%%d.%dq | %%3d | %%6s | %%6.1foz | %%s\n", longestName, longestName-2)
	uRowFormat := fmt.Sprintf("%%%d.%dq | %%3d | %%6s | %8.1s | %%s\n", longestName, longestName-2, "-")
	summaryTitleFormat := fmt.Sprintf("%%%ds | %%-10s\n", longestName)
	summaryRowFormat := fmt.Sprintf("%%%ds | %%10.1f\n", longestName)

	fmt.Printf(
		titleFormat,
		"Name",
		"Age",
		"Active",
		"Mass",
		"Books",
	)
	fmt.Println(strings.Repeat(dash, 2))

	for _, u := range users {
		active := "-"
		if u.Active {
			active = "+"
		}

		age := u.Age
		if age < 0 {
			age = 100 + age
		}

		format := rowFormat
		if u.Mass > 200 {
			format = ozRowFormat
		}

		if u.Mass < 1 {
			u.Mass *= 100
		}

		books := strings.Join(u.Books, ", ")

		if u.Mass == 0 {
			fmt.Printf(uRowFormat, u.Name, age, active, books)
			fmt.Println(strings.Repeat(dash, 2))
			continue
		}

		fmt.Printf(format, u.Name, age, active, u.Mass, books)
		fmt.Println(strings.Repeat(dash, 2))
	}

	fmt.Println("\nSummary:")
	fmt.Println(strings.Repeat(dash2, 2))
	fmt.Printf(
		summaryTitleFormat,
		"Book",
		"Average Age",
	)

	for book, ages := range bookAges {
		averageAge := calculateAverageAge(ages)
		fmt.Println(strings.Repeat(dash2, 2))
		fmt.Printf(summaryRowFormat, book, averageAge)
	}
	fmt.Println(strings.Repeat(dash2, 2))
}

func calculateAverageAge(ages []int) float64 {
	total := 0
	for _, age := range ages {
		total += age
	}
	average := float64(total) / float64(len(ages))
	return average
}
