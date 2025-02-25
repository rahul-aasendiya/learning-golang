package utility

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}

// Public function (exported)
func PublicFunction() {
	fmt.Println("This is public function.")
}

// Private function (unexported)
func privateFunction() {
	fmt.Println("This is private function.")
}
