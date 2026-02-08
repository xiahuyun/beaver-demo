package device

import "strings"

const (
	DesktopWindows  = "beaver_desktop_windows"
	DesktopMacOS    = "beaver_desktop_macos"
	DesktopLinux    = "beaver_desktop_linux"
	MobileAndroid   = "beaver_mobile_android"
	MobileIOS       = "beaver_mobile_ios"
	MobileHarmonyOS = "beaver_mobile_harmonyos"
)

const (
	Desktop = "desktop"
	Mobile  = "mobile"
	Web     = "web"
)

// GetDeviceType 根据 User-Agent 字符串识别设备类型
func GetDeviceType(userAgent string) string {
	// 识别桌面端
	if strings.Contains(userAgent, DesktopWindows) {
		return Desktop
	} else if strings.Contains(userAgent, DesktopMacOS) {
		return Desktop
	} else if strings.Contains(userAgent, DesktopLinux) {
		return Desktop
	}

	// 识别移动端
	if strings.Contains(userAgent, MobileAndroid) {
		return Mobile
	} else if strings.Contains(userAgent, MobileIOS) {
		return Mobile
	} else if strings.Contains(userAgent, MobileHarmonyOS) {
		return Mobile
	}

	// 不是桌面端和移动端的，剩下的都是web端
	return Web
}
