package db

type (
	// DB struct
	DB struct {
		Name string
	}
)

// NewDB returns db instance
func NewDB(hint string) *DB {
	return &DB{
		Name: hint,
	}
}

// GetSession get db session
func (d *DB) GetSession() string {
	// 実際はDB接続情報などを返す
	return d.Name
}
