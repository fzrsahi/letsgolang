package main

import (
	"fmt"
	"time"
)

type CheckClass interface {
	check() bool
}

type ClassA struct {
	ChairmanName string
	SubjectHours string
}

func (c ClassA) check() bool {
	if now := time.Now().Format("1504"); c.SubjectHours != now {
		return false
	}
	return true
}

type ClassB struct {
	ChairmanName string
	SubjectHours string
}

func (c ClassB) check() bool {
	if now := time.Now().Format("1504"); c.SubjectHours != now {
		return false
	}
	return true
}

type ClassC struct {
	ChairmanName string
	SubjectHours string
}

func (c ClassC) check() bool {
	if now := time.Now().Format("1504"); c.SubjectHours != now {
		return false
	}
	return true
}

func CheckClassStatus(class []CheckClass) []CheckClass {
	studyNow := []CheckClass{}
	for _, class := range class {
		classStatus := class.check()
		if classStatus {
			studyNow = append(studyNow, class)
		}
	}
	return studyNow
}

func main() {

	classA := ClassA{"Fazrul", "1355"}

	classB := ClassB{"John Doe", "1355"}

	classC := ClassC{"Agung", "1355"}

	check := CheckClassStatus([]CheckClass{classA, classB, classC})

	fmt.Println("Jadwal Kuliah Sekarang(nama keting,jam kuliah) :", check)

}
