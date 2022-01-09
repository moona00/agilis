package agilis

// Data is the database data
type Data map[string]interface{}

// Database is a simple Agilis database struct
type Database struct {
	Name string
	Data map[string]interface{}
}
