package test_interface

import "strconv"

type Activity interface {
	Run() (int, int)
}

type Person struct {
	Name, Age string
	Calory    int
}

func (person *Person) Run() (int, int) {
	caloryBurn := 30
	return person.Calory - caloryBurn, caloryBurn
}

func Diet(activity Activity) string {
	_, burnedCalory := activity.Run()
	return "Berhasil Bakar Kalori sebanyak : " + strconv.Itoa(burnedCalory)
}
