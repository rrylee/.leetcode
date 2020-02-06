/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

/*
buy=max(buy, sell-prices[i])
sell=max(sell, prices[i]+buy)
*/

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy, sell := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		if sell-prices[i] > buy {
			buy = sell - prices[i]
		}
		if buy+prices[i] > sell {
			sell = buy + prices[i]
		}
	}
	return sell
}

// @lc code=end

