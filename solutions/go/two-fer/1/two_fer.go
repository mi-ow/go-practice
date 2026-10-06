// Package twofer returns a String e.g. One for Do-yun, one for me.
//If name is empty it returns One for you, one for me.
package twofer

// ShareWith returns a String
func ShareWith(name string) string {
	if(name == ""){
		return "One for you, one for me."
    }
    return "One for "+name +", one for me."
}
