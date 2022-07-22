package main

import "fmt"

type HUMAN struct {
	Name            string
	SelfDescription string
	Like            []string
	DisLike         []string
}

func main() {
	BGKang := &HUMAN{
		Name:            "강보권",
		SelfDescription: "많이 부족한 주니어입니다. 잘 부탁드립니다.",
		Like:            []string{"커피", "카페인", "투썸 스트로베리피치 프라페"},
		DisLike:         []string{"더위", "바깥..?"},
	}
	_ = BGKang
	fmt.Println("Please View Code")
	return
}
