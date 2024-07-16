// hello world
// Path: practice/main.go
// hello world

package main

import ("fmt"
		"errors"
	)

func main() {
	const name = "Tuna"
	const age = 30


	ages:=make(map[string]int)

	ages["Tuna"]=30
	ages["Tuna2"]=31

	delete(ages, "Tuna2")

	elem, ok := ages["Tuna2"]
	if ok {
		fmt.Println("Tuna2 is in the map and the value is", elem)
	} else {
		fmt.Println("Tuna2 is not in the map")
	}

	fmt.Println(ages["Tuna2"])

	s := fmt.Sprintf("My name is %v and I am %v years old", name, age)
	fmt.Println(s)
	for i:= 0; i < 10; i++ {
		fmt.Println(i)
	}
	cost := bulkSend(10)
	fmt.Println("Cost is", cost, "dollars for", 10, "messages")

	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(b)
}

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i:= 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01*float64(i))
	}
	return totalCost
}

// define user struct
type user struct {
	name string
}

func deleteFromUser(users map[string]user, name string) (deleted bool, err error) {
	_, ok := users[name]
	if !ok {
		return false, errors.New("user does not exist")
	}
	delete(users, name)
	return true, nil
}

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar:= rune(name[0])
		_,ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

/**
I am interested in working at Trilitech because of their innovative approach to blockchain development, particularly on the Tezos platform. Trilitech's focus on core development, adoption through creative partnerships, and support for entrepreneurial projects aligns with my interests in blockchain technology and community building. Their commitment to fostering a collaborative environment and driving technological advancement makes Trilitech an exciting place to contribute and grow professionally.

*/
