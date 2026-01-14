package main

// MillionToThousand 将元/百万tokens转换为元/千tokens
func MillionToThousand(pricePerMillion float64) float64 {
	return pricePerMillion / 1000
}

// ThousandToMillion 将元/千tokens转换为元/百万tokens
func ThousandToMillion(pricePerThousand float64) float64 {
	return pricePerThousand * 1000
}

// CalculateCost 根据tokens数量和单价计算消费金额
// tokens: tokens数量
// pricePerMillion: 元/百万tokens
// 返回: 消费金额(元)
func CalculateCost(tokens float64, pricePerMillion float64) float64 {
	return tokens * (pricePerMillion / 1000000)
}

// CalculateTokens 根据消费金额和单价计算tokens数量
// cost: 消费金额(元)
// pricePerMillion: 元/百万tokens
// 返回: tokens数量
func CalculateTokens(cost float64, pricePerMillion float64) float64 {
	if pricePerMillion == 0 {
		return 0
	}
	return cost / (pricePerMillion / 1000000)
}

// CalculatePricePerMillion 根据消费金额和tokens数量计算单价
// cost: 消费金额(元)
// tokens: tokens数量
// 返回: 元/百万tokens
func CalculatePricePerMillion(cost float64, tokens float64) float64 {
	if tokens == 0 {
		return 0
	}
	return cost / tokens * 1000000
}
