package main

type Input struct {
    Id int
    AccountPhone string    `json:"account_phone"`
    PassWord string     `json:"password"`
    DingDingAccount string `json:"dingding_account"`
}
type StaffInterface struct{
    Id int
    StaffPhone string
    StaffStatus string
}
type CompanyInterface struct {
    Id int
    CompanyId string
    CompanyName string
}
type SelfDetail struct {
    StaffPhone string `json:"staff_phone"`
    StaffName string   `json:"staff_name"`
    StaffIdentityCard string `json:"staff_identity_card"`
    StaffPhoto string `json:"staff_photo"`
    StaffSex bool `json:"staff_sex"`
    StaffMarry bool `json:"staff_marry"`
    StaffCertification string `json:"staff_certification"`
    StaffCharacter string `json:"staff_character"`
    StaffExperience string `json:"staff_experience"`
    StaffSchool string `json:"staff_school"`
}
type EmploymentStatus struct {
    FileId int `json:"file_id"`
    StaffPhone string `json:"staff_phone"`
    CompanyId string `json:"company_id"`
    DeptId string `json:"dept_id"`
    StaffId string `json:"staff_id"`
    WorkAttendance string `json:"work_attem_dance"`
    WorkPerformance string `json:"work_performance"`
    WorkSalary int `json:"work_salary"`
    WorkInTime string `json:"work_in_time"`
    WorkPost string `json:"work_post"`
    CreaterName string `json:"creater_name"`
}
type Performance struct {
    Id int `json:"id"`
    FileId int `json:"file_id"`
    PerformData string
    PerformScore int
}

type Attendance struct {
    Id int `json:"id"`
    FileId int `json:"file_id"`
    AttendanceData string
    AttendanceSituation string
}

type Mistake struct {
    Id int `json:"id"`
    FileId int `json:"file_id"`
    MistakeData string
    MistakeDetail string
}