// Code generated by entc, DO NOT EDIT.

package schedule

const (
	// Label holds the string label denoting the schedule type in the database.
	Label = "schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldStartAt holds the string denoting the startat field in the database.
	FieldStartAt = "start_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the schedule in the database.
	Table = "schedules"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_schedules"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
)

// Columns holds all SQL columns for schedule fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldStartAt,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "schedule_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)