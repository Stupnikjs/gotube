package main

func main() {
	for {
		cmd, err := Prompt()
		if err != nil {
			panic(err)
		}
		if cmd == "exit" {
			break
		}
		Wrapper(cmd)

	}

}
