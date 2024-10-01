package calculate

import (
	"fmt"
)

func Calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("0으로 나눌 수 없습니다")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("지원하지 않는 연산자입니다: %s", operator)
	}
}
