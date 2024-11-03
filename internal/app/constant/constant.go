package constant

type ResponseStatus int
type Headers int

const (
	ADMIN_ROLE = "ADMIN"
	USER_ROLE  = "USER"
)

const (
	PAYMENT_PENDING = iota + 1
	PAYMENT_COMPLETED
	PAYMENT_CANCELLED
	PAYMENT_FAILED
)

const (
	PAYMODE_CASH = iota + 1
	PAYMODE_VNPAY
	PAYMODE_MOMO
)

const (
	CONTRACT_PROCESSING = iota + 1
	CONTRACT_ACTIVE
	CONTRACT_FINISHED
	CONTRACT_FAILED
	ONE_SIDE_CANCELLED
	AGREE_CANCELLED
	WAITING_FOR_LIQUIDITY
	LIQUIDITY_COMPLETED
)

const (
	ROOM_AVAILABLE = iota + 1
	ROOM_UNAVAILABLE
)

const (
	BOOKING_REQUEST_PROCESSING = "PROCESSING"
	BOOKING_REQUEST_ACCEPTED   = "ACCEPTED"
	BOOKING_REQUEST_REJECTED   = "REJECTED"
)

const (
	TRANSACTION_DEPOSIT = iota + 1
	TRANSACTION_WITHDRAW
	TRANSACTION_PAYMENT
	TRANSACTION_REFUND
)

const (
	TRANSACTION_SUCCESS = iota + 1
	TRANSACTION_FAILED
)

const (
	Success ResponseStatus = iota + 1
	DataNotFound
	InvalidRequest
	InternalServerError
	Unauthorized
	NoContent
	BadRequest
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "INVALID_REQUEST", "INTERNAL_SERVER_ERROR", "UNAUTHORIZED", "NO_CONTENT", "BAD_REQUEST"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data not found", "Invalid request", "Internal server error", "Unauthorized", "No content", "Bad request"}[r-1]
}
