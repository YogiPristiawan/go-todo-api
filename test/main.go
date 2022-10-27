package main

import "fmt"

type res struct {
	Id int
}

type t struct {
	Id int
}

func main() {
	ts := []t{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
		{
			Id: 3,
		},
	}

	r := []res{}

	fmt.Println("before R", r)

	mapping(r, ts)

	fmt.Println("after R", r)

}

func mapping(r []res, tes []t) {
	for _, val := range tes {
		var ress res
		ress.Id = val.Id

		r = append(r, ress)
	}
}
