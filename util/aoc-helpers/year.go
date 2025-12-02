package aoc_helpers

import "fmt"

type AOCYear struct {
	Id   int
	days []*AOCDay
}

func NewYear(id int, days []*AOCDay) *AOCYear {
	return &AOCYear{
		Id:   id,
		days: days,
	}
}

func (y *AOCYear) Run() {
	for _, d := range y.days {
		fmt.Println("==========")
		fmt.Printf("| Running Year %d, Day %d\n", y.Id, d.Id)
		d.Run()
	}
}

func (y *AOCYear) RunDay(id int) {
	fmt.Println("==========")
	for _, d := range y.days {
		if d.Id == id {
			fmt.Printf("| Running Year %d, Day %d\n", y.Id, d.Id)
			d.Run()
			break
		}
	}

}
