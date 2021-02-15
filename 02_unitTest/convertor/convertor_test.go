package convertor

import (
	"testing"

	mock_convertor "github.com/a-berahman/go-time/02_unitTest/convertor/mock"
	"github.com/golang/mock/gomock"
)

func TestPrint(t *testing.T) {
	// Create a controller to manage all our mock objects and make sure
	// that all expectations were met before completing the test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Obtain a mock instance that implements Printer and associate it with the controller.
	printer := mock_convertor.NewMockPrinter(ctrl)
	// Verify not just the call list but the call order as well
	gomock.InOrder(
		printer.EXPECT().MoveZeroes([]int{1, 0, 2, 0, 3}).Return([]int{1, 2, 3, 0, 0}),
		printer.EXPECT().MoveZeroes([]int{1, 0, 2, 0, 3, 4}).Return([]int{1, 2, 3, 4, 0, 0}),
	)
	handler := NewHandler(printer)

	got := handler.GetResult([]int{1, 0, 2, 0, 3})
	if exp := "1,2,3,0,0"; got != exp {
		t.Fatalf("expected strings to be:\n%v\n got:\n%v", exp, got)
	}
	got = handler.GetResult([]int{1, 0, 2, 0, 3, 4})
	if exp := "1,2,3,4,0,0"; got != exp {
		t.Fatalf("expected strings to be:\n%v\n got:\n%v", exp, got)
	}
}
