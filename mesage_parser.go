package main

import (
	"regexp"
	"strconv"
)

type PoloniexMsg struct {
	pairID 					int64 	// <currency pair id>,
	lastTradePrice 			float64	//"<last trade price>",
	lowestAsk 				float64	//"<lowest ask>",
	highestBid 				float64	//"<highest bid>",
	percentChange24 		float64	//"<percent change in last 24 hours>",
	baseCurrencyVolume24 	float64	//"<base currency volume in last 24 hours>",
	quoteCurrencyVolume24 	float64	//"<quote currency volume in last 24 hours>",
	isFrozen				int64	//<is frozen>,
	highestTradePrice 		float64	//"<highest trade price in last 24 hours>",
	lowestTradePrice 		float64	//"<lowest trade price in last 24 hours>"
}

func parseMsg(m []byte) PoloniexMsg {
	var res PoloniexMsg
	// пример строки и паттерн для парсинга
	// [1002,null,[149,"1552.68651502","1551.54650557","1550.39458431","-0.09651590","144058105.67523484","90407.29938610",0,"1802.83144175","1356.56652364"]]
	// \[\d*?,.*?,\[(\d*?),"([\d.-]*?)","([\d.-]*?)","([\d.-]*?)","([\d.-]*?)","([\d.-]*?)","([\d.-]*?)",(\d),"([\d.-]*?)","([\d.-]*?)"\]\]/gm

	re := regexp.MustCompile(`\[\d*?,.*?,\[(\d*?),"([\d.-]*?)","([\d.-]*?)","([\d.-]*?)","([\d.-]*?)","([\d.-]*?)","([\d.-]*?)",(\d),"([\d.-]*?)","([\d.-]*?)"\]`)
	ss := re.FindStringSubmatch(string(m))
	if len(ss)>1 {
		res.pairID, _ = strconv.ParseInt(ss[1],0,64)
		res.lastTradePrice, _ = strconv.ParseFloat(ss[2],64)
		res.lowestAsk, _ = strconv.ParseFloat(ss[3],64)
		res.highestBid, _ = strconv.ParseFloat(ss[4],64)
		res.percentChange24, _ = strconv.ParseFloat(ss[5],64)
		res.baseCurrencyVolume24, _ = strconv.ParseFloat(ss[6],64)
		res.quoteCurrencyVolume24, _ = strconv.ParseFloat(ss[7],64)
		res.isFrozen, _ = strconv.ParseInt(ss[8],0,64)
		res.highestTradePrice, _ = strconv.ParseFloat(ss[9],64)
		res.lowestTradePrice, _ = strconv.ParseFloat(ss[10],64)
	}
	return res
}