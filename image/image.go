package main

import (
	"fmt"
	"github.com/fogleman/gg"
)

func drawBackground(fileName string) {
	// 加载图片
	img, err := gg.LoadJPG(fileName)
	if err != nil {
		panic(err)
	}

	// 获取图片尺寸
	size := img.Bounds().Size()

	width := size.X
	height := size.Y

	// 以加载图片的宽高作为新图片的大小,生成一张画布
	dc := gg.NewContext(width, height)

	// 画图  此步骤是将一张图画在新生成的画布上
	dc.DrawImage(img, 0, 0)

	//在图像上画矩形需要的点
	x, y, w, h := 10.0, 10.0, 100.0, 100.0

	// 画矩形
	dc.SetLineWidth(1) // 设置画线的宽度
	dc.SetRGB(1, 0, 0) // 设置画线的颜色，取值范围[0,1]，其实就是x/255，比如红色是255,0,0，除以255是1,0,0
	dc.DrawRectangle(x, y, w, h)
	dc.Stroke()

	// 保存新图片，一般quality设置为75即可，最高可设置为100，值越高，质量越好，但是占空间大
	err = dc.SaveJPG("after.jpeg", 75)
	if err != nil {
		fmt.Println("SaveJPG err:", err)
		panic(err)
	}
}
