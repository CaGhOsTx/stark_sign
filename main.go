package main

import (
	"os"
	"strconv"

	"github.com/yanue/starkex"
)

func main() {
	networkId, err := strconv.Atoi(os.Args[1])
	market := os.Args[2]
	side := os.Args[3]
	positionId, err := strconv.ParseInt(os.Args[4], 10, 64)
	size := os.Args[5]
	price := os.Args[6]
	fee := os.Args[7]
	clientId := os.Args[8]
	expiration := os.Args[9]
	key := os.Args[10]

    param := starkex.OrderSignParam{
		NetworkId:  networkId,
		Market:     market,
		Side:       side,
		PositionId: positionId,
		HumanSize:  size,
		HumanPrice: price,
		LimitFee:   fee,
		ClientId:  clientId,
		Expiration: expiration,
	}
	sign, err := starkex.OrderSign(key, param)
	fo, err := os.Create("signature.txt")
	if(err == nil) {
		fo.WriteString(sign);
	}
}