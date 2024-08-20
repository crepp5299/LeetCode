func maxProfit(prices []int) int {
    lastPrice := prices[0]
    profit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] < lastPrice {
            lastPrice = prices[i]
            continue
        }
        if prices[i] - lastPrice > profit {
            profit = prices[i] - lastPrice
        }
    }
    return profit
}