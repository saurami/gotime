package partyrobot

import "strconv"

func welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

func happyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " +
			strconv.FormatInt(int64(age), 10) + // strconv.Itoa(age)
			" years old!"
}

func assignTable(name string, table int, neighbor, direction string, distance float64) string {
	return `Welcome to my party, ` + name + `!
You have been assigned to table 0` + strconv.Itoa(table) + `. Your table is ` + direction + `, exactly ` + strconv.FormatFloat(distance, 'g', 3, 64) + ` meters from here.
You will be sitting next to ` + neighbor + `.
`
}
