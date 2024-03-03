package utils

import "fmt"

type Color string

const (
	// ColorPurple 紫色
	ColorPurple = "\x1b[35m"
	// ColorGreen 绿色
	ColorGreen = "\x1b[32m"
	// ColorYellow 黄色
	ColorYellow = "\x1b[33m"
	// ColorRed 红色
	ColorRed = "\x1b[31m"
)

// SetColor 设置颜色
func SetColor(color Color, msg string) string {
	return fmt.Sprintf("%v%v%v", color, msg, color)
}

// SetColorPurple 设置紫色
func SetColorPurple(msg string) string {
	return fmt.Sprintf("%v%v%v", ColorPurple, msg, ColorPurple)
}

// SetColorColorGreen 设置绿色
func SetColorColorGreen(msg string) string {
	return fmt.Sprintf("%v%v%v", ColorGreen, msg, ColorGreen)
}

// SetColorColorYellow 设置黄色
func SetColorColorYellow(msg string) string {
	return fmt.Sprintf("%v%v%v", ColorYellow, msg, ColorYellow)
}

// SetColorColorRed 设置红色
func SetColorColorRed(msg string) string {
	return fmt.Sprintf("%v%v%v", ColorRed, msg, ColorRed)
}
