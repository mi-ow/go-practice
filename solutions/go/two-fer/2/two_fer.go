// Package twofer returns a String e.g. One for Do-yun, one for me.
//If name is empty it returns One for you, one for me.
package twofer
import "fmt"
        

// ShareWith returns a String in the requested format
func ShareWith(name string) string {
	if(name == ""){
		name = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", name)
}
