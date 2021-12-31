package validator

import "fmt"

func intsJoin(arr []int, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func uintsJoin(arr []uint, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func int8sJoin(arr []int8, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func uint8sJoin(arr []uint8, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func int16sJoin(arr []int16, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func uint16sJoin(arr []uint16, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func int32sJoin(arr []int32, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func uint32sJoin(arr []uint32, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func int64sJoin(arr []int64, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func uint64sJoin(arr []uint64, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%d", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func float32sJoin(arr []float32, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%g", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
func float64sJoin(arr []float64, sep string) (str string) {
	for k, v := range arr {
		str += fmt.Sprintf("%g", v)
		if k+1 < len(arr) {
			str += sep
		}
	}
	return str
}
