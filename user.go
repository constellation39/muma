package muma

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type User struct {
	*Request
	*Config

	UUID  string
	Token string

	UserId   int
	Phone    string
	CourseId int
}

func (user *User) Fields() []zap.Field {
	return []zap.Field{
		zap.String("UUID", user.UUID),
		zap.String("Token", user.Token),
		zap.Int("UserId", user.UserId),
		zap.String("Phone", user.Phone),
		zap.Int("CourseId", user.CourseId),
	}
}

func NewUser(config *Config) *User {
	return &User{
		Request: NewRequest(config.Host),
		Config:  config,
		UUID:    uuid.NewString(),
	}
}

type CSLogin struct {
	Equipment LoginEquipment `json:"equipment"`
	UserName  string         `json:"userName"`
	Password  string         `json:"password"`
}

type LoginEquipment struct {
	DeviceBrand            string `json:"deviceBrand"`
	DeviceModel            string `json:"deviceModel"`
	OperationSystem        string `json:"operationSystem"`
	OperationSystemVersion string `json:"operationSystemVersion"`
	ResolutionHeight       int    `json:"resolutionHeight"`
	ResolutionWidth        int    `json:"resolutionWidth"`
	ServiceProvider        string `json:"serviceProvider"`
}

// Login 登录
func (user *User) Login() bool {
	body, err := user.Post("web-gateway/token", CSLogin{
		Equipment: LoginEquipment{
			DeviceBrand:            "Chrome",
			DeviceModel:            fmt.Sprintf("100.0.%d.127", 4000+rand.Intn(2000)),
			OperationSystem:        "Windows 10",
			OperationSystemVersion: "10.0",
			ResolutionHeight:       1920,
			ResolutionWidth:        1080,
			ServiceProvider:        "",
		},
		UserName: user.Username,
		Password: user.Password,
	})
	if err != nil {
		Logger.Error("Login", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	token := fmt.Sprintf("Bearer %s", body)
	user.Token = token
	user.SetHeader("authorization", token)
	Logger.Debug("Login Success", user.Fields()...)
	return true
}

type SCRecentlyWork struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data SCRecentlyWorkData `json:"data"`
}

type SCRecentlyWorkData struct {
	Time       int64       `json:"time"`
	ClassName  string      `json:"className"`
	CourseName interface{} `json:"courseName"`
}

// RecentlyWork 最近工作
func (user *User) RecentlyWork() bool {
	body, err := user.Get("web-gateway/t/student/recently-work")
	if err != nil {
		Logger.Error("RecentlyWork Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scRecentlyWork := new(SCRecentlyWork)
	err = json.Unmarshal(body, scRecentlyWork)
	if err != nil {
		Logger.Error("RecentlyWork Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("RecentlyWork Success", zap.String("id", user.UUID), zap.Reflect("SCRecentlyWork", scRecentlyWork))
	return true
}

type SCUserInfo struct {
	UserBaseDetailInfoDTO UserBaseDetailInfoDTO `json:"userBaseDetailInfoDTO"`
	UserCollegeInfoDTO    []UserCollegeInfoDTO  `json:"userCollegeInfoDTO"`
}

type UserBaseDetailInfoDTO struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Nickname         string      `json:"nickname"`
	Username         string      `json:"username"`
	Number           interface{} `json:"number"`
	Card             string      `json:"card"`
	Qq               string      `json:"qq"`
	Phone            string      `json:"phone"`
	Birthday         string      `json:"birthday"`
	Region           interface{} `json:"region"`
	Email            string      `json:"email"`
	Sex              int         `json:"sex"`
	Industry         interface{} `json:"industry"`
	Hobby            interface{} `json:"hobby"`
	Sign             interface{} `json:"sign"`
	Education        string      `json:"education"`
	Graduated        interface{} `json:"graduated"`
	Secret           interface{} `json:"secret"`
	Type             int         `json:"type"`
	Avatar           string      `json:"avatar"`
	Status           int         `json:"status"`
	Webchat          string      `json:"webchat"`
	MicroBlog        interface{} `json:"microBlog"`
	Address          interface{} `json:"address"`
	WxBindSign       int         `json:"wxBindSign"`
	Nation           string      `json:"nation"`
	Politics         string      `json:"politics"`
	IsDouble         interface{} `json:"isDouble"`
	IsVerifyIdentity int         `json:"isVerifyIdentity"`
	IsInstructor     interface{} `json:"isInstructor"`
	PatrolList       interface{} `json:"patrolList"`
	PatrolDormList   interface{} `json:"patrolDormList"`
	Ad               interface{} `json:"ad"`
}

type UserCollegeInfoDTO struct {
	CollegeRelationID        int         `json:"collegeRelationId"`
	UserID                   int         `json:"userId"`
	ClassID                  int         `json:"classId"`
	ClassName                string      `json:"className"`
	SchoolName               string      `json:"schoolName"`
	DepartmentName           string      `json:"departmentName"`
	MajorName                string      `json:"majorName"`
	SchoolID                 int         `json:"schoolId"`
	DepartmentID             int         `json:"departmentId"`
	MajorID                  int         `json:"majorId"`
	Number                   string      `json:"number"`
	StudentStatus            interface{} `json:"studentStatus"`
	CollegeType              int         `json:"collegeType"`
	StudentStatusInformation int         `json:"studentStatusInformation"`
	DataType                 interface{} `json:"dataType"`
	AttributionFacultyID     interface{} `json:"attributionFacultyId"`
	AttributionFacultyIDName interface{} `json:"attributionFacultyIdName"`
	GradeID                  interface{} `json:"gradeId"`
	GradeName                interface{} `json:"gradeName"`
	ClassType                int         `json:"classType"`
	PreviousClassID          interface{} `json:"previousClassId"`
	DesensitizedStatus       int         `json:"desensitizedStatus"`
}

// UserInfo 用户信息
func (user *User) UserInfo() bool {
	body, err := user.Get("web-gateway/user/userInfo")
	if err != nil {
		Logger.Error("UserInfo Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scUserInfo := new(SCUserInfo)
	err = json.Unmarshal(body, scUserInfo)
	if err != nil {
		Logger.Error("UserInfo Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	user.UserId = scUserInfo.UserBaseDetailInfoDTO.ID
	user.Phone = scUserInfo.UserBaseDetailInfoDTO.Phone
	Logger.Debug("UserInfo Success", zap.String("id", user.UUID), zap.Reflect("SCUserInfo", scUserInfo))
	return true
}

type SCInstructorCheck struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data bool   `json:"data"`
}

// UserInfo 辅导员信息检查
func (user *User) InstructorCheck() bool {
	body, err := user.Get(fmt.Sprintf("web-gateway/instructor/instructor-check?userId=%d", user.UserId))
	if err != nil {
		Logger.Error("InstructorCheck Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scInstructorCheck := new(SCInstructorCheck)
	err = json.Unmarshal(body, scInstructorCheck)
	if err != nil {
		Logger.Error("InstructorCheck Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("InstructorCheck Success", zap.String("id", user.UUID), zap.Reflect("SCInstructorCheck", scInstructorCheck))
	return true
}

type SCGetUserStatus struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// GetUserStatus 获取用户状态
func (user *User) GetUserStatus() bool {
	body, err := user.Get(fmt.Sprintf("web-gateway/teacherAbnormalBehavior/getUserStatus?name=%s", user.Phone))
	if err != nil {
		Logger.Error("GetUserStatus Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scGetUserStatus := new(SCGetUserStatus)
	err = json.Unmarshal(body, scGetUserStatus)
	if err != nil {
		Logger.Error("GetUserStatus Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("GetUserStatus Success", zap.String("id", user.UUID), zap.Reflect("SCGetUserStatus", scGetUserStatus))
	return true
}

type SCCourseRecordInfo struct {
	Total             int    `json:"total"`
	List              []List `json:"list"`
	PageNum           int    `json:"pageNum"`
	PageSize          int    `json:"pageSize"`
	Size              int    `json:"size"`
	StartRow          int    `json:"startRow"`
	EndRow            int    `json:"endRow"`
	Pages             int    `json:"pages"`
	PrePage           int    `json:"prePage"`
	NextPage          int    `json:"nextPage"`
	IsFirstPage       bool   `json:"isFirstPage"`
	IsLastPage        bool   `json:"isLastPage"`
	HasPreviousPage   bool   `json:"hasPreviousPage"`
	HasNextPage       bool   `json:"hasNextPage"`
	NavigatePages     int    `json:"navigatePages"`
	NavigatepageNums  []int  `json:"navigatepageNums"`
	NavigateFirstPage int    `json:"navigateFirstPage"`
	NavigateLastPage  int    `json:"navigateLastPage"`
}
type List struct {
	ID                         int           `json:"id"`
	CourseName                 string        `json:"courseName"`
	TeacherID                  interface{}   `json:"teacherId"`
	TeacherName                interface{}   `json:"teacherName"`
	TeacherHeadImage           interface{}   `json:"teacherHeadImage"`
	TeacherDesc                interface{}   `json:"teacherDesc"`
	TeacherRemark              interface{}   `json:"teacherRemark"`
	Teacherurl                 interface{}   `json:"teacherurl"`
	CollegeID                  interface{}   `json:"collegeId"`
	CollegeName                interface{}   `json:"collegeName"`
	Score                      interface{}   `json:"score"`
	ImageURL                   string        `json:"imageUrl"`
	Tags                       interface{}   `json:"tags"`
	Summary                    interface{}   `json:"summary"`
	Duration                   interface{}   `json:"duration"`
	Description                interface{}   `json:"description"`
	AppDescription             interface{}   `json:"appDescription"`
	Recommend                  interface{}   `json:"recommend"`
	Grounding                  int           `json:"grounding"`
	State                      int           `json:"state"`
	CreateTime                 string        `json:"createTime"`
	CreateUser                 interface{}   `json:"createUser"`
	UpdateTime                 interface{}   `json:"updateTime"`
	UpdateUser                 interface{}   `json:"updateUser"`
	Status                     int           `json:"status"`
	Liked                      interface{}   `json:"liked"`
	Share                      interface{}   `json:"share"`
	Learned                    interface{}   `json:"learned"`
	TotalSubject               int           `json:"totalSubject"`
	Progress                   interface{}   `json:"progress"`
	ProgressRatio              interface{}   `json:"progressRatio"`
	CourseScore                interface{}   `json:"courseScore"`
	Like                       interface{}   `json:"like"`
	CategoryName               interface{}   `json:"categoryName"`
	DateSize                   interface{}   `json:"dateSize"`
	Join                       interface{}   `json:"join"`
	DetailList                 interface{}   `json:"detailList"`
	UseClassList               []interface{} `json:"useClassList"`
	ClassCoursePermissionCount interface{}   `json:"classCoursePermissionCount"`
	ClassCoursePermission      interface{}   `json:"classCoursePermission"`
	TotalTime                  interface{}   `json:"totalTime"`
	TotalCount                 interface{}   `json:"totalCount"`
	Categories                 interface{}   `json:"categories"`
}

// GetUserStatus 获取用户所有的课程
func (user *User) CourseRecordInfo() bool {
	body, err := user.Get("web-gateway/course/courseRecordInfo?pageSize=10&pageNum=1&total=0&orderBy=1&type=1")
	if err != nil {
		Logger.Error("CourseRecordInfo Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scCourseRecordInfo := new(SCCourseRecordInfo)
	err = json.Unmarshal(body, scCourseRecordInfo)
	if err != nil {
		Logger.Error("CourseRecordInfo Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("CourseRecordInfo Success", zap.String("id", user.UUID), zap.Reflect("SCCourseRecordInfo", scCourseRecordInfo))
	return true
}

type SCAuthorization struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data bool   `json:"data"`
}

// Authorization 验证用户能否学习该视频
func (user *User) Authorization(courseId int) bool {
	body, err := user.Get(fmt.Sprintf("web-gateway/course-permission/user/authorization?courseId=%d&userId=%d", courseId, user.UserId))
	if err != nil {
		Logger.Error("Authorization Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scAuthorization := new(SCAuthorization)
	err = json.Unmarshal(body, scAuthorization)
	if err != nil {
		Logger.Error("Authorization Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("Authorization Success", zap.String("id", user.UUID), zap.Reflect("SCAuthorization", scAuthorization))
	return true
}

type SCCourseDetail []struct {
	ID             int              `json:"id"`
	CourseID       int              `json:"courseId"`
	Name           string           `json:"name"`
	SeqNo          int              `json:"seqNo"`
	SectionNo      string           `json:"sectionNo"`
	Description    string           `json:"description"`
	CreateTime     int64            `json:"createTime"`
	UpdateTime     int64            `json:"updateTime"`
	Status         int              `json:"status"`
	SubjectDTOList []SubjectDTOList `json:"subjectDTOList"`
}
type PointVideos struct {
	ID            int         `json:"id"`
	PointID       int         `json:"pointId"`
	SeqNo         interface{} `json:"seqNo"`
	Name          string      `json:"name"`
	FileURL       string      `json:"fileUrl"`
	Duration      string      `json:"duration"`
	Status        int         `json:"status"`
	Description   interface{} `json:"description"`
	CheckPoint    interface{} `json:"checkPoint"`
	IsLearning    bool        `json:"isLearning"`
	QuestionDTO   interface{} `json:"questionDTO"`
	PointExercise interface{} `json:"pointExercise"`
	SubjectLike   interface{} `json:"subjectLike"`
	TotalTime     string      `json:"totalTime"`
	Collect       interface{} `json:"collect"`
}
type SubjectPointList struct {
	ID          int           `json:"id"`
	SectionID   interface{}   `json:"sectionId"`
	SubjectID   int           `json:"subjectId"`
	SeqNo       interface{}   `json:"seqNo"`
	Name        string        `json:"name"`
	Duration    int           `json:"duration"`
	Progress    interface{}   `json:"progress"`
	VideoID     int           `json:"videoId"`
	PointVideos []PointVideos `json:"pointVideos"`
}
type SubjectDTOList struct {
	ID               int                `json:"id"`
	CourseID         int                `json:"courseId"`
	SectionID        int                `json:"sectionId"`
	SeqNo            int                `json:"seqNo"`
	Name             string             `json:"name"`
	Duration         int                `json:"duration"`
	Progress         interface{}        `json:"progress"`
	Status           int                `json:"status"`
	LearnStatus      interface{}        `json:"learnStatus"`
	SubjectPointList []SubjectPointList `json:"subjectPointList"`
	SubjectExercise  interface{}        `json:"subjectExercise"`
	SubjectData      interface{}        `json:"subjectData"`
}

// CourseDetail 获得目标课程的详细章节
func (user *User) CourseDetail(courseId int) bool {
	body, err := user.Get(fmt.Sprintf("web-gateway/course/courseDetail?courseId=%d", courseId))
	if err != nil {
		Logger.Error("CourseDetail Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scCourseDetail := new(SCCourseDetail)
	err = json.Unmarshal(body, scCourseDetail)
	if err != nil {
		Logger.Error("CourseDetail Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("CourseDetail Success", zap.String("id", user.UUID), zap.Reflect("SCCourseDetail", scCourseDetail))
	return true
}

type CSLearnedVideo struct {
	CourseID  int    `json:"courseId"`
	SubjectID string `json:"subjectId"`
	VideoID   int    `json:"videoId"`
}

type SCLearnedVideo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// LearnedVideo 学习这个视频
func (user *User) LearnedVideo(courseID int, subjectID string, videoID int) bool {
	body, err := user.Post("web-gateway/user/learnedVideo", CSLearnedVideo{
		CourseID:  courseID,
		SubjectID: subjectID,
		VideoID:   videoID,
	})
	if err != nil {
		Logger.Error("LearnedVideo Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scLearnedVideo := new(SCLearnedVideo)
	err = json.Unmarshal(body, scLearnedVideo)
	if err != nil {
		Logger.Error("LearnedVideo Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("LearnedVideo Success", zap.String("id", user.UUID), zap.Reflect("SCLearnedVideo", scLearnedVideo))
	return true
}

type SCLoginOut struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// LoginOut 登出
func (user *User) LoginOut() bool {
	body, err := user.Get("web-gateway/user/loginOut")
	if err != nil {
		Logger.Error("LoginOut Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	scLoginOut := new(SCLoginOut)
	err = json.Unmarshal(body, scLoginOut)
	if err != nil {
		Logger.Error("LoginOut Failed", append(user.Fields(), zap.Reflect("err", err))...)
		return false
	}
	Logger.Debug("LoginOut Success", zap.String("id", user.UUID), zap.Reflect("SCLoginOut", scLoginOut))
	return true
}
