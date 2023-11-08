package imgutil

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
)

// // 定义目标宽度和高度（这里是 300 像素）
//	width := 300
//	height := 0 // 0 表示按照比例自动计算高度

func Compression(input string, width int, height int, output string, interp resize.InterpolationFunction) error {
	// 打开要压缩的图片文件
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 使用 resize 库进行缩放
	resizedImg := resize.Resize(uint(width), uint(height), img, interp)

	// 创建输出文件
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	// 将缩放后的图像写入文件
	jpeg.Encode(out, resizedImg, nil)
	return nil
}
