package main

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}
	if jug1Capacity == 0 || jug2Capacity == 0 {
		return targetCapacity == 0 || jug1Capacity+jug2Capacity == targetCapacity
	}
	return targetCapacity%gcd(jug1Capacity, jug2Capacity) == 0

}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
