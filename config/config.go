package config

import "time"

// [server]
var AppMode = "debug"
var HttpPort = ":3000" //端口

// [database]

// [limiter]
var LimiterState = 0 //0不开启限流；1固定窗口；2滑动窗口；3滑动日志；4漏桶；5令牌桶

//通用参数
var SizeInt = 10
var LimitInt = 1

// 固定窗口参数  limit和interval，分别表示限制数量和时间窗口长度。
var FWLimit = 10
var FWInterval = 1 * time.Second

// 滑动窗口参数  windowSize和limit，分别表示窗口大小和限制数量
var SWLimit = 10
var SWWindowSize = 1 * time.Second

// 2滑动日志参数  windowSize和limit，分别表示窗口大小和限制数量
var SLLimit = 10
var SLWindowSize = 1 * time.Second

// 漏桶参数  rate和capacity，分别表示漏水速率和桶容量
var LBCapacity = 10
var LBRate = 1 * time.Second

// 令牌桶参数  rate和capacity，分别表示产生令牌的速率和桶容量。
var TBCapacity = 10
var TBRate = 1 * time.Second
