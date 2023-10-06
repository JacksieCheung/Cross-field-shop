package crawl

import "errors"

const (
	PlatformCCNU = iota
)

type PlatformType int

type ManagerBuilder interface {
	SetUid(uid string) ManagerBuilder
	SetPassword(password string) ManagerBuilder
	Build() (Manager, error)
}

type Manager interface {
	CrawlCourseInfo() (CourseInfo, error)
	CrawlGradeInfo() (GradesInfo, error)
	Reload() error
	Logout() error
}

type GradesInfo []GradeInfo

type CourseInfo struct {
	TotalCoursesNum int      `json:"total_courses_num"`
	Courses         []Course `json:"courses"`
}

const (
	UnCommitted = iota // 未提交
	UnScored           // 待审批
	Rejected           // 被驳回
	Finished           // 有得分
)

type Course struct {
	TotalHomeWorksNum int        `json:"total_home_works_num"` // 总作业数目
	SemesterYear      int        `json:"semester_year"`        // 学期（年） 2022 2021 ...
	SemesterSpan      int        `json:"semester_span"`        // 学期 1，2，3，4
	Name              string     `json:"name"`
	ID                string     `json:"id"` // 课程号
	HomeWorks         []HomeWork `json:"home_works"`
}

type HomeWork struct {
	Name   string  `json:"name"`
	Score  float64 `json:"score"`
	Status int     `json:"status"`
	Time   int64   `json:"time"` // ms
}

type GradeInfo struct {
	CourseName   string  `json:"course_name"`
	GPA          float64 `json:"gpa"`
	SemesterYear int     `json:"semester_year"` // 学期（年） 2022 2021 ...
	SemesterSpan int     `json:"semester_span"` // 学期 1，2，3，4
	Grade        float64 `json:"grade"`
	ID           string  `json:"id"` // 课程号
}

func GetTargetPlatformManagerBuilder(platform PlatformType) (ManagerBuilder, error) {
	switch platform {
	case PlatformCCNU:
		return &spocManagerBuilder{}, nil
	default:
		return nil, errors.New("no such manager builder")
	}
}
