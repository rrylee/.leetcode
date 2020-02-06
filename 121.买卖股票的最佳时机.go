/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

/*
buy=max(buy, -prices[i])
sell=max(sell, prices[i] + buy)
*/

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buy := -prices[0]
	sell := 0

	for i := 1; i < len(prices); i++ {
		if -prices[i] > buy {
			buy = -prices[i]
		}
		if buy+prices[i] > sell {
			sell = buy + prices[i]
		}
	}

	return sell
}

// @lc code=end

