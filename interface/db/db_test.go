package db

import "testing"

func TestDB(t *testing.T) {

	t.Run("mysql", func(t *testing.T) {
		driver := "mysql"
		dsn := "root:12345@tcp(localhost:3306)/test"

		mysql := MySQL{driver, dsn}

		_, err := mysql.Connect()
		if err != nil {
			t.Error("connect failed: ", err)
		}
	})

	t.Run("postgres", func(t *testing.T) {
		driver := "postgres"
		dsn := "postgres://postgres:12345@localhost:5432/test?sslmode=disable"

		postgres := Postgres{driver, dsn}

		_, err := postgres.Connect()
		if err != nil {
			t.Error("connect failed: ", err)
		}

	})

}

func TestDBInterface(t *testing.T) {

	testcases := []struct {
		name string
		sql  SQL
	}{
		{
			name: "mysql",
			sql:  MySQL{driver: "mysql", dsn: "root:12345@tcp(localhost:3306)/test"},
		},
		{
			name: "postgres",
			sql:  Postgres{driver: "postgres", dsn: "postgres://postgres:12345@localhost:5432/test?sslmode=disable"},
		},
	}

	for _, test := range testcases {

		t.Run(test.name, func(t *testing.T) {
			_, err := test.sql.Connect()
			if err != nil {
				t.Error("connect failed: ", err)
			}
		})
	}

}
