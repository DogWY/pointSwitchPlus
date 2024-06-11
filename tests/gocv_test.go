package test

//import (
//	"fmt"
//	"gocv.io/x/gocv"
//	_ "gocv.io/x/gocv"
//	"testing"
//)
//
//func TestGocv(t *testing.T) {
//	capture, err2 := gocv.OpenVideoCapture()
//	capture.Close()
//	img := gocv.IMRead("static/img/1.png", gocv.IMReadGrayScale)
//	//buf := bytes.Buffer{}
//	//i, _, _ := image.Decode(bytes.NewReader(img.ToBytes()))
//	//jpeg.Encode(&buf, i, &jpeg.Options{
//	//	Quality: 40,
//	//})
//	b := img.ToBytes()
//	//window := gocv.NewWindow("hello")
//	//for {
//	//	window.IMShow(img)
//	//	window.WaitKey(1)
//	//}
//	newImg, err := gocv.NewMatFromBytes(342, 428, 0, b)
//	if err != nil {
//		fmt.Printf("生成图片失败")
//		return
//	}
//	gocv.IMWrite("./1.png", newImg)
//	fmt.Printf("生成图片成功")
//}
