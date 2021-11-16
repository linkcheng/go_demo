package main

import (
	"fmt"
	"proto/message"

	"google.golang.org/protobuf/proto"
)

// protoc --go_out=paths=source_relative:./message orders.proto

func main() {
    orders := &message.Orders{
        OrderId: 1,
        Title:   "第一个订单",
    }
    //序列化成二进制数据
    ordersBytes, _ := proto.Marshal(orders)
	
	//反序列化二进制数据
    twoOrders := &message.Orders{}
    proto.Unmarshal(ordersBytes, twoOrders)
    
	fmt.Println(twoOrders.GetTitle())
    fmt.Println(twoOrders.GetOrderId())
}