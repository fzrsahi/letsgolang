package main

import "fmt"

type ValidationError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func insertData(id string, data string) (string, error) {
	dbId := "123"
	dataInserted := false
	if dbId != id {
		return "", &NotFoundError{"Data Not Found"}
	}

	if dataInserted {
		return "", &ValidationError{"Data already inserted!"}
	}

	return "Data Inserted!", nil
}

func main() {

	result, err := insertData("1243", "barang")
	if err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Println("Error Validasi :", validationErr.Message)
		} else if notFoundErr, ok := err.(*ValidationError); ok {
			fmt.Println("Error Not Found :", notFoundErr.Message)
		} else {
			fmt.Println("Unknown Error :", err.Error())
		}
	} else {
		fmt.Println(result)
	}

}
