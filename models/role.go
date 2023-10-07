package models

type Role struct {
	Id                         string `bson:"id"`
	Name                       string
	Description                string
	CanApproveVisaApplications bool `bson:"can_approve_visa_application"`
	CanViewVisaApplications    bool `bson:"can_view_visa_application"`
}

// 0 = player
// 1 = admin
