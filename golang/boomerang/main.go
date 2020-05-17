package main

func isBoomerang(points [][]int) bool {
	a := points[0]
	b := points[1]
	c := points[2]
	area := a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1])
	return area != 0
}

func main() {

}
