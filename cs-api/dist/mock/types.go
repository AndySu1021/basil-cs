package mock

type MockSqlResult struct {
}

func (m MockSqlResult) LastInsertId() (int64, error) {
	return 1, nil
}

func (m MockSqlResult) RowsAffected() (int64, error) {
	return 1, nil
}
