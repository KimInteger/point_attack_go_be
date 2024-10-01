package main

import (
	"fmt"
	"myapp/calculate" // 패키지 경로 (여기서 myapp은 모듈 이름)
)

func main() {
	var num1, num2 float64
	var operator string

	// 사용자로부터 입력을 받음
	fmt.Println("간단한 계산기입니다.")
	fmt.Print("첫 번째 숫자를 입력하세요: ")
	fmt.Scanln(&num1)
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("두 번째 숫자를 입력하세요: ")
	fmt.Scanln(&num2)

	// calculate 패키지의 Calculate 함수 호출
	result, err := calculate.Calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("오류:", err)
	} else {
		fmt.Printf("결과: %.2f\n", result)
	}
}
