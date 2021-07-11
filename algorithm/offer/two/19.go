package two

/**
请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。


input: s=aa p=a
output: false
*/

func Hand19(s, p string) bool {
	// 动态规划 dp[i][j] 前j个p能否匹配到前i个s
	rows, cols := len(s), len(p)
	dp := make([][]bool, rows+1)
	for i := range dp {
		dp[i] = make([]bool, cols+1)
	}
	dp[0][0] = true
	for i := 2; i <= cols; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			curS, curP := s[i-1], p[j-1]
			if curS == curP || curP == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if curP == '*' && j >= 2 {
				if p[j-2] == '.' || p[j-2] == curS {
					dp[i][j] = dp[i-1][j] || dp[i][j-1] // 匹配多次或者1次
				}
				dp[i][j] = dp[i][j] || dp[i][j-2] //匹配0次
			}
		}
	}
	return dp[rows][cols]
}
