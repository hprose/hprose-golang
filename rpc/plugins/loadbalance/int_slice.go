/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/int_slice.go                     |
|                                                          |
| LastModified: Mar 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

type intSlice []int

func (nums intSlice) Aggregate(f func(int, int) int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	current := nums[0]
	for i := 1; i < n; i++ {
		current = f(current, nums[i])
	}
	return current
}

func (nums intSlice) Sum() int {
	return nums.Aggregate(func(x, y int) int {
		return x + y
	})
}

func (nums intSlice) Min() int {
	return nums.Aggregate(func(x, y int) int {
		if x > y {
			return y
		}
		return x
	})
}

func (nums intSlice) Max() int {
	return nums.Aggregate(func(x, y int) int {
		if x > y {
			return x
		}
		return y
	})
}

func (nums intSlice) GCD() int {
	return nums.Aggregate(gcd)
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
