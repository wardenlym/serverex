// +build windows

package log

// 在linux和mac平台下，支持使用 kill -USR2 pid 来动态 enable debug log
func EnableSwitchBySIGUSR2(f func(debugEnabled bool)) {

}
