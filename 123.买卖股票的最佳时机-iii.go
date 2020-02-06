/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

/*
firstBuy = max(firstBuy, -prices[i])
firstSell = max(firstSell, prices[i] + firstBuy)
secondBuy = max(secondBuy, firstSell-prices[i])
secondSell = max(secondSell, prices[i] + secondBuy)
*/

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	firstBuy := -prices[0]
	firstSell := 0
	secondBuy := firstBuy
	secondSell := 0
	for i := 1; i < len(prices); i++ {
		firstBuy = maxf(firstBuy, -prices[i])
		firstSell = maxf(firstSell, prices[i]+firstBuy)
		secondBuy = maxf(secondBuy, firstSell-prices[i])
		secondSell = maxf(secondSell, prices[i]+secondBuy)
	}

	return secondSell
}

func maxf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
