package langs

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func connectDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS test_content (
			name TEXT NOT NULL PRIMARY KEY,
			value TEXT
		)
	`)
	require.NoError(t, err)

	return db
}

func TestContentValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	content := NewContentFromMap(map[Lang]string{
		ES: "es-content",
		EN: "en-content",
	})
	_, err := db.Exec("INSERT INTO test_content (name, value) VALUES (?, ?)", "foo", content)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value FROM test_content WHERE name = ?", "foo").Scan(&val))

	require.JSONEq(t, `{"en": "en-content", "es": "es-content"}`, val)
}

func TestContentValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var content Content
	_, err := db.Exec("INSERT INTO test_content (name, value) VALUES (?, ?)", "foo", content)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value FROM test_content WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "{}", val)
}

func TestContentScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_content (name, value) VALUES (?, ?)", "foo", `{"en": "en-content", "es": "es-content"}`)
	require.NoError(t, err)

	var content Content
	require.NoError(t, db.QueryRow("SELECT value FROM test_content WHERE name = ?", "foo").Scan(&content))

	require.Equal(t, content.Get(ES), "es-content")
	require.Equal(t, content.Get(EN), "en-content")
}

func TestContentScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_content (name, value) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var content Content
	require.NoError(t, db.QueryRow("SELECT value FROM test_content WHERE name = ?", "foo").Scan(&content))

	require.True(t, content.IsEmpty())
}

func TestContentSaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	content := NewContentFromMap(map[Lang]string{
		ES: "es-content",
		EN: "en-content",
	})
	_, err := db.Exec("INSERT INTO test_content (name, value) VALUES (?, ?)", "foo", content)
	require.NoError(t, err)

	var other Content
	require.NoError(t, db.QueryRow("SELECT value FROM test_content WHERE name = ?", "foo").Scan(&other))

	require.Equal(t, content.Get(ES), "es-content")
	require.Equal(t, content.Get(EN), "en-content")
}