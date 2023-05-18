// Package twofer provides a function to generate dialogue for sharing cookies.
// The ShareWith function determines the appropriate dialogue based on a given name.

package twofer

func ShareWith(name string) string {
	//ShareWish returns the appropriate, personalized dialog the user should use for sharing their cookies
	if name == ""{
		name = "you"
	}
	return "One for " + name + ", one for me."
}


//Write the package comment to summarize what the program is about x
//write a doc string to summarize the ShareWith function x
//Create a conditional where if a name is passed in to ShareWith as an argument, 
	//the return statement will say "One for you {insert name}, one for me". 
//Otherwise if no name is passed into ShareWith as an argument, if it is an empty string
	//the return statement will say "One for you, one for me."