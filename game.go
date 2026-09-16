package main

func checkWinnings(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		check := row[0]
		for _, symbol := range row[1:] {
			if check != symbol {
				win = false
				break
			}
		}

		if win {
			lines = append(lines, multipliers[check])
		} else {
			lines = append(lines, 0)
		}
	}

	return lines
}
