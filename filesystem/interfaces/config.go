package interfaces

/**
 * Config interface
 *
 * @create 2021-8-1
 * @author deatil
 */
type Config interface {
	// With data
	With(map[string]any) Config

	// Set kv dta
	Set(string, any) Config

	// Has
	Has(string) bool

	// Get one data
	Get(string, ...any) any
}
