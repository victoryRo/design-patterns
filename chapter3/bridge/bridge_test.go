package bridge

import (
	"errors"
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrintImple1{}

	err := api1.PrintMessage("Hello!")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrintImple2{}

	err := api2.PrintMessage("Hi !")
	if err != nil {
		expMsg := errPrint2.Error()
		if !strings.Contains(err.Error(), expMsg) {
			t.Errorf("Error msg was no correct.\n Actual: %s\nExpected: %s\n", err.Error(), expMsg)
		}
	}

	testWriter := TestWriter{}
	api2 = PrintImple2{
		Writer: &testWriter,
	}

	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer.\n Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expMsg := "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expMsg,
		Printer: &PrintImple1{},
	}

	err := normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expMsg,
		Printer: &PrintImple2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expMsg {
		t.Errorf("The expected message of the io.Writer doesn't match actual.\n Actual: %s\nExpected: %s\n", testWriter.Msg, expMsg)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	message := "Hello io.Writer again"
	expMsg := "Message from Packt: Hello io.Writer again"

	packt := PacktPrinter{
		Msg:     message,
		Printer: &PrintImple1{},
	}

	err := packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: message,
		Printer: &PrintImple2{
			Writer: &testWriter,
		},
	}

	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expMsg {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n  Actual: %s\nExpected: %s\n", testWriter.Msg, expMsg)
	}
}

// --------------------------------------------- implement interface io.Writer

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}
