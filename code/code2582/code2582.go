package code2582

// passThePillow 2852题 递枕头
/*
	思路：判断time对2(n-1)取模为t，随之判断t是大于n-1还是小于，获取最后值
*/
func passThePillow(n int, time int) int {
	t := time % (2 * (n - 1))
	var result int
	if t <= (n - 1) {
		result = 1 + t
	} else {
		result = n - (t - (n - 1))
	}
	return result
}
