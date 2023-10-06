package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation   = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase     = &Errno{Code: 20002, Message: "Database error."}
	ErrToken        = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrBase64Decode = &Errno{Code: 20004, Message: "Error occurred while decode base64"}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrTokenExpired      = &Errno{Code: 20104, Message: "The token is expired."}
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "The password was incorrect."}
	ErrAuthFailed        = &Errno{Code: 20106, Message: "The sid or password was incorrect."}
	ErrMissingHeader     = &Errno{Code: 20107, Message: "The length of the `Authorization` header is zero."}

	ErrCrawlerBuilder = &Errno{Code: 20201, Message: "method calls mangerBuilder.Build() error"}

	// upload errors
	ErrFileNotFound = &Errno{Code: 20301, Message: "File not found"}
	ErrFileInvalid  = &Errno{Code: 20302, Message: "Invalid file: only support file.mht"}
	ErrUploadFailed = &Errno{Code: 20303, Message: "Fail to upload file"}

	ErrQueryUserMht = &Errno{Code: 20401, Message: "Error happended when QueryUserMht"}
)
