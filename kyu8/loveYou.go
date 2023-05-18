package kata

func HowMuchILoveYou(i int) string {
	var options [6]string = [6]string{
		"i love you",
		"a little",
		"a lot",
		"passionately",
		"madly",
		"not at all",
	}

	pos := i % 6

	return options[pos-1]
}

func main() {
	HowMuchILoveYou(7)
	HowMuchILoveYou(2)
	HowMuchILoveYou(1)
	HowMuchILoveYou(5)
	HowMuchILoveYou(62)
}
