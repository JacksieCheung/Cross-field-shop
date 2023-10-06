package search

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "Data-acquisition-subsystem/handler"
	"Data-acquisition-subsystem/log"
	"Data-acquisition-subsystem/model"
	"Data-acquisition-subsystem/pkg/errno"
	"Data-acquisition-subsystem/util"
)

type HomeworkQueryReq struct {
	Course       string `json:"course"`     // 可以为空
	StudentID    string `json:"student_id"` // 可以为空
	TimeStart    int64  `json:"time_start"`
	TimeEnd      int64  `json:"time_end"`
	SemesterYear int    `json:"semester_year"`
	SemesterSpan int    `json:"semester_span"`
}

type HomeworkInfo struct {
	CourseName   string  `json:"course_name"`
	CourseID     string  `json:"course_id"`
	StudentID    string  `json:"student_id"`
	Status       int     `json:"status"` // 作业完成状态
	Score        float64 `json:"score"`  // 成绩
	Time         int64   `json:"time"`   // 时间
	Name         string  `json:"name"`
	SemesterYear int     `json:"semester_year"`
	SemesterSpan int     `json:"semester_span"`
}

type HomeworkResp []HomeworkInfo

func QueryHomework(c *gin.Context) {
	log.Info("User LoginCCNU function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var err error
	// 解析 body
	req := HomeworkQueryReq{
		Course:       c.Query("course"),
		StudentID:    c.Query("student_id"),
		TimeStart:    0,
		TimeEnd:      0,
		SemesterYear: 0,
		SemesterSpan: 0,
	}
	req.TimeStart, _ = strconv.ParseInt(c.Query("time_start"), 10, 64)
	req.TimeEnd, _ = strconv.ParseInt(c.Query("time_end"), 10, 64)
	req.SemesterYear, _ = strconv.Atoi(c.Query("semester_year"))
	req.SemesterSpan, _ = strconv.Atoi(c.Query("semester_span"))

	// 查询相关的所有作业成绩
	// 1. 课程为空，学号为空：查所有
	// 2. 课程不为空，学号不为空：查该课程，返回该学号数据
	// 3. 课程不为空，学号为空：查课程，返回所有学生
	// 4. 课程为空，学号不为空：查该学生的所有课成绩
	var homeworks []model.HomeWorkModel
	if len(req.Course) != 0 {
		homeworks, err = model.QueryHomeworksByCourseNameAndSid(req.Course, req.StudentID, req.SemesterYear, req.SemesterSpan, req.TimeStart, req.TimeEnd)
		if err != nil {
			SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
			return
		}
	} else if len(req.StudentID) != 0 {
		// len(req.Course) == 0 && len(req.StudentID) != 0
		homeworks, err = model.QueryHomeworksByStudentID(req.StudentID, req.SemesterYear, req.SemesterSpan, req.TimeStart, req.TimeEnd)
		if err != nil {
			SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
			return
		}
	} else {
		// len(req.Course) == 0 && len(req.StudentID) == 0
		homeworks, err = model.QueryAllHomeworks(req.SemesterYear, req.SemesterSpan, req.TimeStart, req.TimeEnd)
		if err != nil {
			SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
			return
		}
	}

	res := make(HomeworkResp, len(homeworks))
	for idx, homework := range homeworks {
		res[idx].CourseName = homework.CourseName
		res[idx].CourseID = homework.CourseID
		res[idx].StudentID = homework.StuId
		res[idx].SemesterYear = homework.SemesterYear
		res[idx].SemesterSpan = homework.SemesterSpan
		res[idx].Time, _ = strconv.ParseInt(homework.Time, 10, 64)
		res[idx].Score = homework.Score
		res[idx].Status = homework.Status
		res[idx].Name = homework.HomeWorkName
	}

	SendResponse(c, errno.OK, res)
}
