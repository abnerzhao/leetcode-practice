// 青蛙跳台阶问题，本质还是递归算法，斐波那契
// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/submissions/

func numWays(n int) int {
    if n == 0 {
        return 1
    }
    if n <= 2 {
        return n
    }
    a := 1
    b := 2
    c := 0
    for i := 3; i <= n; i++ {
        c = (a+b)% 1000000007
        a = b
        b = c
    }
    return c
}