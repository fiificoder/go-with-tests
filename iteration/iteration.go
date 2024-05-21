package iteration

// func iteration() {
//  Repeat("a")
// //	fmt.Print(lc)
// }

func Repeat(s string,num int) string{
	var save string
	for i := 0; i < num; i++ {
	//	fmt.Print(s)
		save += s
	}
	return save
}