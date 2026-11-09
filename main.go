package main

func main() {}

type formatter interface {
	format() string
}
type plainText struct {
	message string
}
type bold struct {
	message string
}
type code struct {
	message string
}

func format(f formatter) string {
	switch v := f.(type) {
	case plainText:
		return v.format()
	case bold:
		return v.format()
	case code:
		return v.format()
	default:
		return ""
	}
}
func (p plainText) format() string {
	return p.message
}
func (b bold) format() string {
	return "**" + b.message + "**"
}
func (c code) format() string {
	return "`" + c.message + "`"
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
