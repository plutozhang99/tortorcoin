package utils

// PairStatus 定义 PairStatus 类型，方便类型检查和自动补全
type PairStatus string

// 定义可能的 PairStatus 值
const (
	Success  PairStatus = "success"
	Pending  PairStatus = "pending"
	Rejected PairStatus = "rejected"
)

// ErrorType 定义 ErrorType 类型，方便类型检查和自动补全
type ErrorType string

func (e ErrorType) Error() string {
	return string(e)
}

const (
	ErrAccountExists          ErrorType = "account already exists"
	ErrUserNotFound           ErrorType = "user not found"
	ErrPairNotFound           ErrorType = "pair not found"
	ErrPaymentFailed          ErrorType = "payment failed"
	ErrPaymentHistoryNotFound ErrorType = "payment history not found"
	ErrInsertFailed           ErrorType = "insert failed"
	ErrUpdateFailed           ErrorType = "update failed"
	ErrGetPaymentFailed       ErrorType = "get payment failed"
	ErrGetPairFailed          ErrorType = "get pair failed"
	ErrPairAlreadyExists      ErrorType = "pair already exists"
	ErrGetUserFailed          ErrorType = "get user failed"
	ErrGetRecordFailed        ErrorType = "get record failed"
)
