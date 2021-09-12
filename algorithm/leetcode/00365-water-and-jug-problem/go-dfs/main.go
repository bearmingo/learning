package main

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	stack := make([]data, 0)
	seen := make(map[data]int)
	stack = append(stack, data{remainX: 0, remainY: 0})

	for len(stack) > 0 {
		d := stack[0]
		stack = stack[1:]
		if d.remainX == targetCapacity || d.remainY == targetCapacity || d.remainX+d.remainY == targetCapacity {
			return true
		}
		if _, ok := seen[d]; ok {
			continue
		}

		seen[d] = 1

		// 把 X 壶灌满。
		stack = append(stack, data{remainX: jug1Capacity, remainY: d.remainY})
		// 把 Y 壶灌满。
		stack = append(stack, data{remainX: d.remainX, remainY: jug2Capacity})
		// 把 X 壶倒空。
		stack = append(stack, data{remainX: 0, remainY: d.remainY})
		// 把 Y 壶倒空。
		stack = append(stack, data{remainX: d.remainX, remainY: 0})
		// 把 X 壶的水灌进 Y 壶，直至灌满或倒空。
		stack = append(stack, data{remainX: d.remainX - min(d.remainX, jug2Capacity-d.remainY), remainY: d.remainY + min(d.remainX, jug2Capacity-d.remainY)})
		// 把 Y 壶的水灌进 X 壶，直至灌满或倒空。
		stack = append(stack, data{remainX: d.remainX + min(d.remainY, jug1Capacity-d.remainX), remainY: d.remainY - min(d.remainY, jug1Capacity-d.remainX)})
	}

	return false

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

type data struct {
	remainX, remainY int
}
