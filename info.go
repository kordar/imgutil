package imgutil

import (
	"image"
	"os"
)

// GetWH 获取图片宽高
func GetWH(source string) (int, int, error) {
	// 打开图片文件
	file, err := os.Open(source)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	// 输出宽度和高度
	return img.Width, img.Height, nil
}
