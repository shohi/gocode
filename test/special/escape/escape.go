package escape

type Displayer interface {
	Display() string
}

type Field struct {
	Name  string
	Value interface{}
}

func (f Field) Display() string {
	return f.Name
}

func appendField(buf []Field, f Field) []Field {
	buf = append(buf, f)
	return buf
}
