package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Active bool
	Mass   float64
}

func Users() []User {
	return []User{
		{
			"John Doe",
			30,
			true,
			80.0,
		},
		{
			"Jake Doe",
			20,
			false,
			60.0,
		},
		{
			" Jane Doe ",
			150,
			true,
			.75,
		},
		{
			"\t",
			-10,
			true,
			8000.0,
		},
		{
			"Vm0weE5GVXhUblJWV0dSUFZtMW9WVll3WkRSV1ZteDBaRVYwVmsxWGVGWlZiVEZIWVd4S2MxTnNiR0ZXVm5Cb1ZsVmFWMVpWTVVWaGVqQTk=\nVm0weE5GVXhUblJWV0dSUFZtMW9WVll3WkRSV1ZteDBaRVYwVmsxWGVGWlZiVEZIWVd4S2MxTnNiR0ZXVm5Cb1ZsVmFWMVpWTVVWaGVqQTk=",
			0,
			true,
			0,
		},
		{
			"\x00\x10\x20\x30\x40\x50\x60\x70",
			0,
			true,
			0,
		},
	}
}

func main() {
	users := Users()

	const sourceLenght = 10
	const dash = "-------------------------------------------"

	longestName := len("Name")

	for i, u := range users {
		if len(u.Name) > sourceLenght {
			users[i].Name = users[i].Name[:7] + "..."
		}
		nameLen := len(fmt.Sprintf("%10.10q", u.Name))

		if nameLen > longestName {
			longestName = nameLen
		}
	}

	titleFormat := fmt.Sprintf("%%%ds | %%3s | %%6s | %%-8s\n", longestName)
	rowFormat := fmt.Sprintf("%%%d.%dq | %%3d | %%6s | %%6.1fkg\n", longestName, longestName-2)
	ozRowFormat := fmt.Sprintf("%%%d.%dq | %%3d | %%6s | %%6.1foz\n", longestName, longestName-2)
	uRowFormat := fmt.Sprintf("%%%d.%dq | %%3d | %%6s | %8.1s\n", longestName, longestName-2, "-")

	fmt.Printf(
		titleFormat,
		"Name",
		"Age",
		"Active",
		"Mass",
	)

	fmt.Printf("%s\n", dash)

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

		if u.Mass == 0 {
			fmt.Printf(uRowFormat, u.Name, age, active)
			fmt.Printf("%s\n", dash)
			continue
		}
		fmt.Printf(format, u.Name, age, active, u.Mass)
		fmt.Printf("%s\n", dash)
	}
}
