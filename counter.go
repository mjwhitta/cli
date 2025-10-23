package cli

import "fmt"

// Counter allows incrementing a value by counting the number of times
// a flag was passed, as in: -v -v -v
type Counter uint

// Inc will increment the Counter.
func (c *Counter) Inc() error {
	*c++
	return nil
}

// IsBoolFlag allows Counter to be called as --flag rather than
// --flag=true.
func (c *Counter) IsBoolFlag() bool {
	return true // Allow shorthand
}

// String will return a string representation of the Counter.
func (c *Counter) String() string {
	return fmt.Sprintf("%d", c)
}

// Set will call Inc().
func (c *Counter) Set(_ string) error {
	return c.Inc()
}
