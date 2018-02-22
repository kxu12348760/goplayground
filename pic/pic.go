package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		var picrow []uint8
		for j := 0; j < dx; j++ {
			picrow = append(picrow, uint8((i + j) / 2))
		}
		pic = append(pic, picrow)
	}
	return pic
}

funct main() {
	pc.Show(Pic)
}
