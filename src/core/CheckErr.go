package core

// CheckErr func
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
