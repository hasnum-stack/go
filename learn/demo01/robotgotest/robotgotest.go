package robotgotest

import (
	"github.com/go-vgo/robotgo"
)

// 大写是导出
func MainRobotgotest() {
	// color := robotgo.GetPixelColor(800, 400)
	// robotgo.Move(237, 724)
	// robotgo.Click("left")
	// robotgo.Move(224, 620)
	// robotgo.Click("left")
	// time.Sleep(1000)
	// robotgo.Move(325, 618)
	// robotgo.Click("left")

	for i := 0; i < 1000; i++ {
		robotgo.Click("left")
	}
	// robotgo.
	// fmt.Printf("color: %v\n", color)

	// x, y := robotgo.GetMousePos()
	// fmt.Println("pos: ", x, y)
}
