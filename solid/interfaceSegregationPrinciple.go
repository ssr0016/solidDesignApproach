package solid

// The Interface Segregation Principle states that clients should not be forced to depend on interfaces they do not use.
// This principle encourages creating smaller, more focused interfaces rather than large, monolithic ones.

type Document interface {
	Read() string
	Write(content string)
	Print() string
}

type TextDocument struct {
	content string
}

func (d *TextDocument) Read() string {
	return d.content
}

func (d *TextDocument) Write(content string) {
	d.content = content
}

func (d *TextDocument) Print() string {
	return "Printing: " + d.content
}

type ReadOnlyDocument struct {
	content string
}

func (d *ReadOnlyDocument) Read() string {
	return d.content
}

func (d *ReadOnlyDocument) Write(content string) {
	// Not supported
}

func (d *ReadOnlyDocument) Print() string {
	// Not supported
	return ""
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(content string)
}

type Printer interface {
	Print() string
}

type TextDocument1 struct {
	content string
}

// Implement Reader, Writer, and Printer for TextDocument

type ReadOnlyDocument1 struct {
	content string
}

// Implement only Reader for ReadOnlyDocument
