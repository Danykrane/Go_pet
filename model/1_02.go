package model

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// func test() {
// 	u := User{
// 		Name: "",
// 		Age:  0,
// 	}
// }