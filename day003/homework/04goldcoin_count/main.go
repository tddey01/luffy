package main

import "fmt"

/*
50枚金币，分配给一下几个人：Matthew，Sarah,Augustus ,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth.
分配规则如下：
a.名字中包含e或者E：1枚金币
b.名字中包含i或者I：2枚金币
c.名字中包含o或者O：3枚金币
d.名字中包含u或者U：4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
*/

var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func distriConut() int {
	sum := 0
	for _, name := range users {
		for _, char := range name {
			switch char {
			case 'e', 'E':
				distribution[name] = distribution[name] + 1
				sum = sum + 1
			case 'i', 'I':
				distribution[name] = distribution[name] + 2
				sum = sum + 2
			case 'o', 'O':
				distribution[name] = distribution[name] + 3
				sum = sum + 3
			case 'u', 'U':
				distribution[name] = distribution[name] + 4
				sum = sum + 4
			}
		}
	}
	return coins - sum
}

func main() {
	left := distriConut()
	fmt.Println(left)
	for k, v := range distribution {
		fmt.Println(k, v)
	}
}
