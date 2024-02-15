package resources

type ServiceResultData struct {
	Data interface{}
}

type ServiceError struct {
	ErrorCode uint
	ErrorMsg  string
}

type ServiceResult struct {
	Code int
	ServiceResultData
	IsError bool
	Error   ServiceError
}