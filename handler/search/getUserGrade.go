package search

//
//import (
//	"sort"
//	"strconv"
//
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//
//	. "Data-acquisition-subsystem/handler"
//	"Data-acquisition-subsystem/log"
//	"Data-acquisition-subsystem/model"
//	"Data-acquisition-subsystem/pkg/errno"
//	"Data-acquisition-subsystem/util"
//)
//
//type GradeQueryReq struct {
//	Course       string `json:"course"`
//	StudentID    string `json:"student_id"`
//	SemesterYear int    `json:"semester_year"`
//	SemesterSpan int    `json:"semester_span"`
//}
//
//type GradeInfo struct {
//	CourseName string  `json:"course_name"`
//	StudentID  string  `json:"student_id"`
//	Rank       int     `json:"rank"`
//	GPA        float64 `json:"gpa"`
//	Grade      float64 `json:"grade"`
//}
//
//type QueryResp []GradeInfo
//
//func QueryGrade(c *gin.Context) {
//	log.Info("User LoginCCNU function called.",
//		zap.String("X-Request-Id", util.GetReqID(c)))
//
//	var err error
//	// 解析 body
//	req := GradeQueryReq{
//		Course:       c.Query("course"),
//		StudentID:    c.Query("student_id"),
//	}
//	req.SemesterYear, _ = strconv.Atoi(c.Query("semester_year"))
//	req.SemesterSpan, _ = strconv.Atoi(c.Query("semester_span"))
//
//	// 查询该课程相关的所有成绩
//	// 1. 课程为空，学号为空：查所有
//	// 2. 课程不为空，学号不为空：查该课程，返回该学号数据
//	// 3. 课程不为空，学号为空：查课程，返回所有学生
//	// 4. 课程为空，学号不为空：查该学生的所有课成绩
//	m := make(map[string][]GradeInfo)
//	var grades []model.GradeModel
//	if len(req.Course) != 0 {
//		grades, err = model.QueryCoursesByCourse(req.Course, req.SemesterYear, req.SemesterSpan)
//		if err != nil {
//			SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
//			return
//		}
//	} else if len(req.StudentID) != 0 {
//		// len(req.Course) == 0 && len(req.StudentID) != 0
//		courses, err := model.QuerySidCourses(req.StudentID, req.SemesterYear, req.SemesterSpan)
//		if err != nil {
//			SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
//			return
//		}
//		grades, err = model.QueryCoursesByCoursesID(courses, req.SemesterYear, req.SemesterSpan)
//		if err != nil {
//			SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
//			return
//		}
//	} else {
//		// len(req.Course) == 0 && len(req.StudentID) == 0
//		grades, err = model.QueryAllCourses(req.SemesterYear, req.SemesterSpan)
//		if err != nil {
//			SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
//			return
//		}
//	}
//	for _, grade := range grades {
//		m[grade.CourseID] = append(m[grade.CourseID], GradeInfo{
//			CourseName: grade.CourseName,
//			StudentID:  grade.StuId,
//			Rank:       0,
//			GPA:        grade.GPA,
//			Grade:      grade.Grade,
//		})
//	}
//
//	for k, v := range m {
//		sort.Slice(v, func(i, j int) bool {
//			return v[i].Grade > v[j].Grade
//		})
//		m[k] = v
//		for idx := range m[k] {
//			m[k][idx].Rank = idx + 1
//		}
//	}
//
//	var res QueryResp
//	if len(req.StudentID) == 0 {
//		for _, course := range m {
//			res = append(res, course...)
//		}
//	} else {
//		for courseIdx, course := range m {
//			for stuIdx, stu := range course {
//				if stu.StudentID == req.StudentID {
//					res = append(res, m[courseIdx][stuIdx])
//					break
//				}
//			}
//		}
//	}
//
//	SendResponse(c, errno.OK, res)
//}
