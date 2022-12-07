package user

type UserRepository interface {
	// HistoryReview(int, []interface{}) error
	// HistoryBooking(int, []interface{}) error
	FindByEmail(string, interface{}) error
	UpdateProfile(int, string) error
}
