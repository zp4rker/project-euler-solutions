package main

func main() {
	sum := 0

	prev, current := 0, 1

	for current < 4_000_000 {
		if current%2 == 0 {
			// println(current)
			sum += current
		}
		prev, current = current, prev+current
	}

	println(sum)
}
