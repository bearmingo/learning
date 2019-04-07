package main

// 直接减divisor的计算次数dividend/divisor
// 所以使用位移，先找出一个n, 使得divisor*(2^n) * 2 大于被除数且divisor*(2^n) 小于被除数
// 这个步骤的计算次数是log_2(dividend/divisor)
// 例如100 / 3: 直接减是需要计算33次
// 使用方法2:
//    2^5 * 3 = 96, n = 5, n计算需要6次循环，剩余4
//    2^0 * 3 = 3, n = 0,  n计算需要1次循环 剩余1
//     总共循环次数是6 + 1 = 7
//
// 还要处理一个异常dividend=-2147483648 divisor=-1， 这是题目要求返回2147483647(
// 2147483648这个值超出了MaxInt32最大值)
//
// go的int 是64位的，所以这里也需要特别处理下
import "fmt"

const MaxUint32 = int(^uint32(0))
const MinUint32 = 0
const MaxInt32 = int(MaxUint32 >> 1)
const MinInt32 = -MaxInt32 - 1

func divide(dividend int, divisor int) int {
	// Special cases
	if dividend == MinInt32 && divisor == -1 {
		return MaxInt32
	}
	if divisor == dividend {
		return 1
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}
	if dividend == 0 || divisor == MinInt32 {
		return 0
	}

	// Sign
	plus := true
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		plus = false
	}

	a, b := dividend, divisor
	if dividend < 0 {
		a = -dividend
	}
	if divisor < 0 {
		b = -divisor
	}
	if a < b {
		return 0
	}

	ret := 0
	remaind := a

	for remaind >= b {
		sub := b
		count := 1
		for remaind > (sub << 1) {
			sub <<= 1
			count <<= 1
		}
		remaind -= sub
		ret += count
	}

	if plus {
		return ret
	}
	return -ret
}

func main() {
	fmt.Println(divide(-2147483648, -1))
}
