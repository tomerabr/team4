package main

import "time"

type User struct {
	ID			int	   	  `json:"id"`
	Username 	string 	  `json:"username"`
	Password 	string 	  `json:"password"`
	Base_ID  	string 	  `json:"base_ID"`
	Branch_ID 	string 	  `json:"branch_ID"`
	Capsule_ID 	string 	  `json:"capsule_ID"`
	Permissions []bool 	  `json:"permissions"`
	Admin 		[]bool 	  `json:"admin"`
	TimeCreated time.Time `json:"time"`
}

type Space struct {
	Name	string `json:"name"`
	Base	string `json:"base_name"`
}

type Capsule struct {
	Users	[]int `json:"users"`
}

type Branch struct {
	Name	string `json:"name"`
	Users	[]int `json:"users"`
}

type Base struct {
	Name	string `json:"name"`
	Spaces	[]int `json:"spaces"`
}

type Schedule struct {
	Space	   int 		 `json:"id"`
	Start	   string    `json:"start"`
	End		   string 	 `json:"end"`
	Date	   time.Time `json:"date"`
	Capsule_ID int		 `json:"capsule"`
	User_ID	   int 		 `json:"user"`
}