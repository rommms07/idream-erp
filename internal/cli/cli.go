package cli

func Start(args []string) {
	for _, val := range args[1:] {
		println(val)
	}
}
