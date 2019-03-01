package main

import "fmt"

//begin show OMIT
func main() {
	greeting, err := filterGreet("BeagleBoy")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Here is my greeting:", greeting)
}

func filterGreet(name string) (greeting string, err error) {
	if name == "" {
		greeting, err = "", fmt.Errorf("Not supposed to greet the void...")
	} else {
		msg, err := fmt.Sprint("Hello ", name), error(nil)
		if name == "Huey" || name == "Dewey" || name == "Louie" {
			greeting = msg
		} else {
			err = fmt.Errorf("not a legit nephew : %q", name)
			greeting = fmt.Sprintf("No greeting because of error [%v]", err.Error())
		}
	}
	return
}

//end show OMIT
func greet(name string) (greeting string, err error) {
	return fmt.Sprint("Hello ", name), nil
}
