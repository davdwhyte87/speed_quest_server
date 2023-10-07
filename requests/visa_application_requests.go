package requests



type VisaApplicationRequest struct{
    Name string             `json:"name,omitempty" validate:"required"`
    Phone    string             `json:"phone,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Location    string             `json:"location,omitempty" validate:"required"`
	Profession    string  `json:"profession,omitempty" validate:"required"`
	ApplicationAnswers [] ApplicationAnswersRequest 
}

type ApplicationAnswersRequest struct {
    QuestionId string             
    YesNoAnswer    bool            
	TextAnswer    string                       	
}