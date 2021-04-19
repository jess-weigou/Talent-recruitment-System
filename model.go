package main

import "time"

type AccountTable struct {
    Uuid string `json:"uuid"`
    AccountPhone string    `json:"account_phone"`
    DingdingAccount string `json:"dingding_account"`
    Password string     `json:"password"`
    Position string    `json:"position"`
}
type StaffInterface struct{
    Id int
    StaffPhone string
}
type CompanyInterface struct {
    Id int
    CompanyId string
    CompanyName string
}
type SelfDetails struct {
    StaffPhone2 string `json:"staff_phone"`
    StaffName string   `json:"staff_name"`
    StaffIdentityCard string `json:"staff_identity_card"`
    StaffPhoto string `json:"staff_photo"`
    StaffSexuality string `json:"staff_sex"`
    StaffMarry string `json:"staff_marry"`
    StaffCertification string `json:"staff_certification"`
    StaffCharacter string `json:"staff_character"`
    StaffExperience string `json:"staff_experience"`
    StaffSchool string `json:"staff_school"`
    StaffEducation string `json:"staff_education"`
    StaffMajor string `json:"staff_major"`
}
type EmploymentStatus struct {
    FileId int `json:"file_id"`
    StaffPhone1 string `json:"staff_phone"`
    CompanyId string `json:"company_id"`
    DeptId string `json:"dept_id"`
    DeptName string `dept_name`
    StaffId string `json:"staff_id"`
    WorkAttendance float64 `json:"work_attem_dance"`
    WorkSalary int `json:"work_salary"`
    WorkInTime time.Time `json:"work_in_time"`
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