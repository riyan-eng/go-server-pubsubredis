package util

type convert struct{}

func Convert() *convert {
	return &convert{}
}

func (c *convert) AnyToString(arg any) string {
	if arg == nil {
		return ""
	}
	return arg.(string)
}

func (c *convert) AnyToInt(arg any) int {
	return arg.(int)
}

func (c *convert) AnyToFloat32(arg any) float32 {
	return arg.(float32)
}

func (c *convert) AnyToFloat64(arg any) float64 {
	return arg.(float64)
}

func (c *convert) AnyToBool(arg any) bool {
	return arg.(bool)
}
