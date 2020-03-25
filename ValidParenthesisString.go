package main

func checkValidString(s string) bool {
	starStack, openBracketStack := make([]int, 0), make([]int, 0)

	//traverse the string
	for i := range s {
		if s[i] == '(' {
			openBracketStack = append(openBracketStack, i)
		} else if s[i] == '*' {
			starStack = append(starStack, i)
		} else {
			//Case when s[i] == ')', check if the openBracket contains element/indexes
			if len(openBracketStack) > 0 {
				//pop the element
				openBracketStack = openBracketStack[:len(openBracketStack)-1] // -1 and not -2 because the outer range is exclusive while taking slice and giving range
			} else if len(starStack) > 0 {
				//pop the element from starStack as it can replace for missing '('
				starStack = starStack[:len(starStack)-1]
			} else {
				return false
			}
		}
	}

	//Balance the open Brackets, if the openBracketStack is not empty
	for len(openBracketStack) > 0 {
		//check if the starStack is not empty and the top of the starStack is greater than the top of the openBracketStack
		if len(starStack) != 0 && starStack[len(starStack)-1] > openBracketStack[len(openBracketStack)-1] {
			//pop the elements from both the stacks
			starStack = starStack[:len(starStack)-1]
			openBracketStack = openBracketStack[:len(openBracketStack)-1]
		} else {
			return false
		}
	}

	return true
}
