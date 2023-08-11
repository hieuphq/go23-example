package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/dwarvesf/go23/ex8/model"
	"github.com/dwarvesf/go23/ex8/validation"
)

func main() {
	u := model.User{
		ID:       1,
		Username: "andy",
		Age:      19,
	}

	rs := validation.ValidateStruct(u)

	fmt.Println(rs)

	u1 := model.User{
		ID:       1,
		Username: "dwarvesf",
		Age:      10,
		Classes: []model.Class{
			{ID: 1, Name: "Math", Teacher: "Mr. A"},
			{ID: 2, Name: "English", Teacher: "Mr. B"},
		},
	}
	u2 := model.User{
		ID:       2,
		Username: "flydragon",
		Age:      99,
		Classes: []model.Class{
			{ID: 1, Name: "Math", Teacher: "Mr. A"},
			{ID: 2, Name: "English", Teacher: "Mr. B"},
		},
	}

	ul1, ul2 := []model.User{u1, u2}, []model.User{u2, u1}
	rsCmp := slices.EqualFunc(ul1, ul2, func(a, b model.User) bool {
		if a.ID != b.ID {
			return false
		}
		if a.Username != b.Username {
			return false
		}
		if a.Age != b.Age {
			return false
		}

		return true
	})
	fmt.Println(rsCmp)

	m := map[string]interface{}{}

	m["a"] = 1
	m["b"] = "2"
	m["c"] = 3.0

	m2 := map[string]interface{}{}

	m2["c"] = 3.0
	m2["a"] = 1
	m2["b"] = "2"
	m2["c"] = 3.0

	fmt.Println(maps.Equal(m, m2))

}
