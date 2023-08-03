package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {
	m := make(map[int]int)

	for _, cur := range droneIds {
		m[cur] = m[cur] + 1
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -1
}
