package main

import (
	"fmt"

	"github.com/aphill70/sheet-rotation/persistence"
)

func main() {
	sheet, _ := persistence.NewSheet("1y3ySYxxxsmLRSZBKJz0CFk9KCdvm8H38pjMcH_Uzixk")

	rotations, _ := sheet.GetRotations()

	if len(rotations) > 0 {
		rotations[0].GetNextYearsRotation("2023")

		rotations[1].GetNextYearsRotation("2023")

		fmt.Printf("%+v\n", rotations[0].RecipientToGiver)
		fmt.Printf("%+v\n", rotations[1].RecipientToGiver)

		sheet.WriteNewAssignments("2023", rotations)
	}

}
