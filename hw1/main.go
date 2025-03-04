package main

import (
	"fmt"
)

const N = 25

func main() {
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			pics_01(x, y)
			pics_02(x, y)
			pics_03(x, y)
			pics_04(x, y)
			pics_05(x, y)
			pics_06(x, y)
			pics_07(x, y)
			pics_08(x, y)
			pics_09(x, y)
			pics_10(x, y)
			pics_11(x, y)
			pics_12(x, y)
			pics_13(x, y)
			pics_14(x, y)
			pics_15(x, y)
			pics_16(x, y)
			pics_18(x, y)
			pics_19(x, y)
			pics_20(x, y)
			pics_21(x, y)
			pics_23(x, y)
			pics_24(x, y)
			pics_25(x, y)
		}
		fmt.Println("")
	}
}

func pics_01(x, y int) {
	show(x < y)
}

func pics_02(x, y int) {
	show(x == y)
}

func pics_03(x, y int) {
	show(y == N-x-1)
}

func pics_04(x, y int) {
	show(y <= N-x+4)
}

func pics_05(x, y int) {
	show(y == 2*x+1 || y == 2*(x-1)+2)
}

func pics_06(x, y int) {
	show(x < (N-4)/2 || y < (N-4)/2)
}

func pics_07(x, y int) {
	show(x > (N+5)/2 && y > (N+5)/2)
}

func pics_08(x, y int) {
	show(x == 0 || y == 0)
}

func pics_09(x, y int) {
	show(x-(N-4)/2 > y || x+(N-4)/2 < y)
}

func pics_10(x, y int) {
	show(x < y && y < 2*x+2)
}

func pics_11(x, y int) {
	show((!(x == 0 && y == 0) && (x == 1 || y == 1)) || (!(x == N-1 && y == N-1) && (x == N-2 || y == N-2)))
}

func pics_12(x, y int) {
	show(y*y <= 400-x*x) // hint
}

func pics_13(x, y int) {
	show(x+y >= 20 && x+y <= 28) // hint
}

func pics_14(x, y int) {
	show(x*y <= 100) // hint
}

func pics_15(x, y int) {
	show((x-y >= 10 && y-x > -21) || (y-x >= 10 && x-y > -21))
}

func pics_16(x, y int) {
	show(x < y+10 && x > y-10 && y < N-x+9 && y > N-x-11)
}

func pics_18(x, y int) {
	show(((x != 0 || y != 0) && (x == 0 || y == 0)) || (x == 1 || y == 1))
}

func pics_19(x, y int) {
	show((x == 0 || y == 0) || (x == N-1 || y == N-1))
}

func pics_20(x, y int) {
	show(x%2 == y%2)
}

func pics_21(x, y int) {
	show(y%(x+1) == 0)
}

func pics_23(x, y int) {
	show(x%3 == 0 && y%2 == 0)
}

func pics_24(x, y int) {
	show(x == y || y == N-x-1)
}

func pics_25(x, y int) {
	show(x%6 == 0 || y%6 == 0)
}

func show(condition bool) {
	if condition {
		fmt.Print("* ")
	} else {
		fmt.Print(". ")
	}
}
