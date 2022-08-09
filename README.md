# golang_reverse_integer

Given a signed 32-bit integer `x`, return `x` *with its digits reversed*. If reversing `x` causes the value to go outside the signed 32-bit integer range `[-$2^{31}$, $2^{31} - 1$]`, then return `0`.

**Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**

## Examples

**Example 1:**

```
Input: x = 123
Output: 321

```

**Example 2:**

```
Input: x = -123
Output: -321

```

**Example 3:**

```
Input: x = 120
Output: 21

```

**Constraints:**

- $`-2^{31} <= x <= 2^{31} - 1`$

## 解析

給定一個 32-bit 的有號整數 num

要求寫一個演算法來反轉以10為base 的digit

如果是 32-bit 的有號整數

代表第一個位數是是 sign bit 用來表示症負號

代表最小數值是 -$2^{31}$ = 0b1000,0000,0000,0000,0000,0000,0000,0000

代表最大數值是 $2^{31}-1$=0b0111,1111,1111,1111,1111,1111,1111,1111

需要注意 % 運算有以下特性

假設 a > 0, b > 0

則 -a % b = -(a%b)

所以可以利用這個特性

每次用% 算出當下 digit

然後用 / 10 做 digit shift

當 res > 0, 檢查 $2^{31}-1$ 的前一個 digit最大數值 =( $2^{31}-1$)/10   

當 res  = ( $2^{31}-1$)/10  , 則 digit 必須要 ≤ 7 = 0b0111 才能放入 32-bit 的有號整數

當 res < 0, 檢查 -$2^{31}$ 的前一個 digit最小數值 =( $-2^{31}$)/10   

當 res = ( $-2^{31}$)/10 , 則 digit 必須 ≥ -8 = 0b1000 才能放入32-bit 的有號整數

每次當發現digit 可以放入則 更新 res = res * 10 + digit

當 x 被 shift 到 0 時

回傳 res

因為計算跟 digit 個數有關所以時間複雜度= O(logn)  , n 代表輸入值的大小

## 程式碼
```go
package sol

import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		// positive out of range
		if res > math.MaxInt32/10 || (res == math.MaxInt32 && digit > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32 && digit < -8) {
			return 0
		}
		res = res*10 + digit
	}
	return res
}

```
## 困難點

1. 需要想出怎麼檢查超過範圍方法
2. 需要有2 補數的概念
3. 需要知道 % 的特性

## Solve Point

- [x]  初始化 res := 0, digit := 0
- [x]  當 x ≠ 0, 做以下運算
- [x]  digit = x % 10,  x /= 10
- [x]  當 res > 最大有號正整數/10 或是 (res =  最大有號正整數/10 且 digit > 7) 則回傳 0
- [x]  當 res < 最小有號負整數/10 或是 (res =  最小有號負整數/10 且 digit < -8) 則回傳 0
- [x]  更新 res = res * 10 + digit
- [x]  回傳 res