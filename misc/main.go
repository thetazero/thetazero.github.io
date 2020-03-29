package main

import (
	"fmt"
	"strings"
)

const (
	INSERT  = 1
	DELETE  = 1
	REPLACE = 6969
)

func main() {
	// str1 := "hello good sir how do you do"
	// str2 := "boy do I love an ounce of rocks"

	str1 :=
		`Financial assistance is now available for low-income residents. Destination: Home has raised more than $11 million from private and public sources.
Where to donate N95 masks, medical supplies, and other hospital needs.
County mosquito control scheduled for Thursday near the Baylands, a closure is anticipated.
New online tools to be well and have fun with your family or with friends virtually.
Read the full Coronavirus Daily Report for more in-depth information.`
	str2 :=
		`Welcome Fire Station 3 fire crews back to the neighborhood (virtually of course)! 
Share the creative ways you are adapting to our new shelter in place environment.  
Parking lots soon to be closed at open space preserves. 
Looking for ways to help or need help? Find resources in this newsletter.    
Read the full Coronavirus Daily Report for more in-depth information.`

	memo := make(map[[2]int]theseData)
	res := editDistance(memo, str1, str2, len(str1), len(str2))
	fmt.Println(str1)
	constructEdit(memo, str1, str2, len(str1), len(str2), true)
	fmt.Println(res)
}

func format(str string) string {
	return "\"" + strings.ReplaceAll(str, "\n", "\\n") + "\","
}

func constructEdit(memo map[[2]int]theseData, str1, str2 string, i, j int, consise bool) {
	if i == 0 && j == 0 {
		return
	}
	this := memo[[2]int{i, j}]
	// fmt.Println(this)
	new := str1
	if consise {
		if this.Opperation == 1 { //insert
			new = str1[:i] + string(str2[j-1]) + str1[i:]
		} else if this.Opperation == 2 { //delete
			new = str1[:i-1] + str1[i:]
		} else if this.Opperation == 3 { //replace
			new = str1[:i-1] + string(str2[j-1]) + str1[i:]
		}
		if this.Opperation != 0 {
			fmt.Print("\n", new, "\n")
		}
	} else {
		if this.Opperation != 0 {
			fmt.Print("\033[1;2m", [2]int{i, j}, " ~> ", this.BestIndex, "\033[0m  ")
		}
		if this.Opperation == 1 { //insert
			new = str1[:i] + string(str2[j-1]) + str1[i:]
			fmt.Println("[insert]:", str1, "\033[1m~>\033[0m", new)
		} else if this.Opperation == 2 { //delete
			new = str1[:i-1] + str1[i:]
			fmt.Println("[delete]:", str1, "\033[1m~>\033[0m", new)
		} else if this.Opperation == 3 { //replace
			new = str1[:i-1] + string(str2[j-1]) + str1[i:]
			fmt.Println("[replace]:", str1, "\033[1m~>\033[0m", new)
		}
	}
	constructEdit(memo, new, str2, this.BestIndex[0], this.BestIndex[1], consise)
}

func editDistance(memo map[[2]int]theseData, str1 string, str2 string, i int, j int) theseData {

	key := [2]int{i, j}
	value, ok := memo[key]
	if ok {
		return value
	}

	// fmt.Println(key)
	ans := theseData{}
	if i == 0 && j == 0 {
		return theseData{}
	} else if i == 0 { //first string empty so inserts
		// fmt.Println(key, "i~>", [2]int{i, j - 1})
		ans = editDistance(memo, str1, str2, i, j-1)
		ans.Opperation = 1
		ans.BestIndex = [2]int{i, j - 1}
		ans.Cost += INSERT
	} else if j == 0 { //second string empty so delete
		// fmt.Println(key, "d~>", [2]int{i - 1, j})
		ans = editDistance(memo, str1, str2, i-1, j)
		ans.BestIndex = [2]int{i - 1, j}
		ans.Opperation = 2
		ans.Cost += DELETE
	} else if str1[i-1] == str2[j-1] {
		// fmt.Println(key, "n~>", [2]int{i - 1, j - 1})
		ans = editDistance(memo, str1, str2, i-1, j-1)
		ans.Opperation = 0
		ans.BestIndex = [2]int{i - 1, j - 1}
	} else {
		// fmt.Println(key, "i~>", [2]int{i, j - 1})
		insert := editDistance(memo, str1, str2, i, j-1)
		insert.BestIndex = [2]int{i, j - 1}
		insert.Opperation = 1
		insert.Cost += INSERT

		// fmt.Println(key, "d~>", [2]int{i - 1, j})
		delete := editDistance(memo, str1, str2, i-1, j)
		delete.BestIndex = [2]int{i - 1, j}
		delete.Opperation = 2
		delete.Cost += DELETE

		// fmt.Println(key, "r~>", [2]int{i - 1, j - 1})
		replace := editDistance(memo, str1, str2, i-1, j-1)
		replace.BestIndex = [2]int{i - 1, j - 1}
		replace.Opperation = 3
		replace.Cost += REPLACE
		if insert.Cost < delete.Cost && insert.Cost < replace.Cost {
			// fmt.Println("insert:", string(str2[j]))
			ans = insert
		} else if delete.Cost < replace.Cost {
			// fmt.Println("delete:", string(str1[i]))
			ans = delete
		} else {
			// fmt.Println("replace", string(str1[i]), "with", string(str2[j])))
			ans = replace
		}
	}
	memo[key] = ans
	return ans
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type theseData struct {
	Cost       int
	BestIndex  [2]int
	Opperation byte //0 nothing, 1 insert, 2 delete, 3 replace
}
