package main

import (
	"github.com/aphill70/sheet-rotation/persistence"
)

func main() {
	sheet, _ := persistence.NewSheet("1y3ySYxxxsmLRSZBKJz0CFk9KCdvm8H38pjMcH_Uzixk")

	rotations, _ := sheet.GetRotations()

	if len(rotations) > 0 {
		rotations[0].GetNextYearsRotation("2020")

		rotations[1].GetNextYearsRotation("2020")

		sheet.WriteNewAssignments("2020", rotations)
	}

}
