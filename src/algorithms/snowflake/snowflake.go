package snowflake

import "time"

/*
组成:
	1bit 空 + 41bit 时间戳 + 10bit 节点 + 12bit序列号
*/

func CreateIdBuilder (sn int64)func()int64{
	if sn >1023 || sn <0{
		panic("节点数值在0到1023之间")
	}
	var auto int64
	var leftTm int64
	return func()int64{
	Start:
		tm := time.Now().UnixNano()/1000000
		if tm > leftTm{
			auto = 1
		}else if tm == leftTm{
			if auto>4096{
				time.Sleep(time.Millisecond)
				goto Start
			}
			auto++
		}else {
			panic("生成时间异常")
		}
		leftTm = tm
		rightBinValue := tm & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		id := rightBinValue | auto | sn
		return id
	}
}
