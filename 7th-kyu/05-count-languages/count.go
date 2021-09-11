package main

func Count(l []developer) map[string]int {
	m := make(map[string]int)

	for _, dev := range l {
		// Access a non-existed key in map results in 0
		m[dev.language] += 1
	}

	return m
}
