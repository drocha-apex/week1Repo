package main

import (
	"fmt"
	"math/rand"
)

const MAX_ELFS = 20
const MAX_FOOD_CAP = 6

type MaxElf struct {
	elfNum, elfCal int
}

func ElfTotalCalorie(elf []int) int {
	total_cal := 0
	for _, food := range elf {
		total_cal += food
	}
	return total_cal
}

func initalizeElfs() [][]int {
	elfs := make([][]int, MAX_ELFS)
	for i := 0; i < MAX_ELFS; i++ {
		elfs[i] = make([]int, MAX_FOOD_CAP)
	}

	for i := 0; i < MAX_ELFS; i++ {
		for j := 0; j < MAX_FOOD_CAP; j++ {
			elfs[i][j] = rand.Intn(800)
		}
	}
	return elfs
}

func MostFood(HungryElfs [][] int)  {
	maxCalElf := MaxElf{
		 elfNum: 0, 
		 elfCal: 0,
	}
	for i := 0; i < MAX_ELFS; i++ {


		if 
		maxCalElf.elfCal = ElfTotalCalorie(elfs[i])
	}
	return 
}

func main() {
	elfs := initalizeElfs()

	for i := 0; i < MAX_ELFS; i++ {
		fmt.Print("Elf #", i + 1, " has a total of ", ElfTotalCalorie(elfs[i]), " Calories\n")
	}

}
