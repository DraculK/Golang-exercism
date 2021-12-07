package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	welcome := "Welcome to my party, " + name + "!"
    return welcome
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	aniversary := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return aniversary
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcome := "Welcome to my party, " + name + "!"
    assign := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
    sit := "You will be sitting next to " + neighbor +"."
    final := welcome + "\n" + assign + "\n" + sit
    return final
}