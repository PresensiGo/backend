package constants

type AttendanceStatus string
type AttendanceStatusType string

const (
	AttendanceStatusTypePresentOnTime AttendanceStatusType = "present-on-time"
	AttendanceStatusTypePresentLate   AttendanceStatusType = "present-late"
	AttendanceStatusTypeSick          AttendanceStatusType = "sick"
	AttendanceStatusTypePermission    AttendanceStatusType = "permission"
	AttendanceStatusTypeAlpha         AttendanceStatusType = "alpha"
)

const (
	AttendanceStatusPresent    AttendanceStatus = "hadir"
	AttendanceStatusAlpha      AttendanceStatus = "alpha"
	AttendanceStatusSick       AttendanceStatus = "sakit"
	AttendanceStatusPermission AttendanceStatus = "izin"
)

func (a AttendanceStatusType) ToAttendanceStatus() AttendanceStatus {
	switch a {
	case AttendanceStatusTypePresentOnTime:
		return AttendanceStatusPresent
	case AttendanceStatusTypePresentLate:
		return AttendanceStatusPresent
	case AttendanceStatusTypeSick:
		return AttendanceStatusSick
	case AttendanceStatusTypePermission:
		return AttendanceStatusPermission
	case AttendanceStatusTypeAlpha:
		return AttendanceStatusAlpha
	default:
		return ""
	}
}
