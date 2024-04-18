package main

func main() {

	x := 1

	for {
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
	}

}
