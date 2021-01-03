// 斐波那契数列,求第N项的值
// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// 0 1 1 2 3 5

func fib(n int) int {
    f0, f1 := 0, 1
    for i := 0 ; i < n; i++{
        f0,f1 = f1, (f0 + f1) % 1000000007;
    }
    return f0
}

// fibV2 递归实现，比较耗时编译超时
func fibV2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibV2(n) + fibV2(n-1)
}