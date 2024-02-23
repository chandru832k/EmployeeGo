package resources

type AddEmployeeInput struct {
	EmployeeName  	string  `json:"employee_name" validate:"required"`
	EmployeeAge 		int     `json:"employee_age"`
	EmployeeGender  string  `json:"employee_gender"`
	EmployeeRole  	string  `json:"employee_role"`
	EmployeeMail 		string  `json:"employee_mail"`
}
