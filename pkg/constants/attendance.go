package constants

type AttendanceStatus string

const (
	AttendanceStatusPresent    AttendanceStatus = "hadir"
	AttendanceStatusAlpha      AttendanceStatus = "alpha"
	AttendanceStatusSick       AttendanceStatus = "sakit"
	AttendanceStatusPermission AttendanceStatus = "izin"
)
