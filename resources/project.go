package resources

type ProjectInput struct {
	Name  					string  `json:"name" validate:"required"`
	Status 					string  `json:"status"`
	Description			string 	`json:"description"`
	CreatedBy  			string  `json:"created_by"`
	CreatedAt  		  int  		`json:"created_at"`
	UpdatedBy  			string  `json:"updated_by"`
	UpdatedAt  		  int  		`json:"updated_at"`
}
