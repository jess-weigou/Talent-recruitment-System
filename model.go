package main

type Input struct {
    UserName string    `json:"user_name"`
    PassWord string     `json:"pass_word"`
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