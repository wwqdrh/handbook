package five

/**
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

*/

func Hand43(n int) int {
	res, digit := 0, 1
	high, cur, low := n/10, n%10, 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		high, cur = high/10, high%10
		digit *= 10
	}
	return res
}
