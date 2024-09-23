package constant

type ResponseStatus int
type Headers int

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
	CONTRACT_AVAILABLE = iota + 1
	CONTRACT_UNAVAILABLE
)

const (
	ROOM_AVAILABLE = iota + 1
	ROOM_UNAVAILABLE
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
