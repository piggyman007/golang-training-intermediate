package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Album struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

func main() {
	resp, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}

	users := new([]User)
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", users)

	///////////////////////////////////////////////////////////////

	// resp, err = http.Get("http://jsonplaceholder.typicode.com/albums")
	// if err != nil {
	// 	panic(err)
	// }

	// b, err = ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// resp.Body.Close()

	// albums := new([]Album)
	// err = json.Unmarshal(b, albums)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v\n", albums)

}
