package convertor

import (
	"strconv"
	"strings"
)

type Printer interface {
	MoveZeroes(nums []int) []int
}

type Handler struct {
	printer Printer
}

func NewHandler(printer Printer) *Handler {
	return &Handler{printer: printer}
}

func (h *Handler) GetResult(item []int) string {
	arr := h.printer.MoveZeroes(item)
	res := []string{}
	for _, v := range arr {
		res = append(res, strconv.Itoa(v))
	}

	return strings.Join(res, ",")
}
