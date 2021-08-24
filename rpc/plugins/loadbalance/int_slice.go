/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/int64_slice.go                   |
|                                                          |
| LastModified: Aug 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

type int64Slice []int64

func (nums int64Slice) Aggregate(f func(int64, int64) int64) int64 {
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

func (nums int64Slice) Sum() int64 {
	return nums.Aggregate(func(x, y int64) int64 {
		return x + y
	})
}

func (nums int64Slice) Min() int64 {
	return nums.Aggregate(func(x, y int64) int64 {
		if x > y {
			return y
		}
		return x
	})
}

func (nums int64Slice) Max() int64 {
	return nums.Aggregate(func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	})
}

func (nums int64Slice) GCD() int64 {
	return nums.Aggregate(gcd)
}

func gcd(x, y int64) int64 {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
