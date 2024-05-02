package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello world."

	t.Run("old printer", func(t *testing.T) {
		adapter := PrinterAdapter{
			OldPrinter: &MyLegacyPrinter{},
			Msg:        msg,
		}

		exp := "Legacy printer: adapter: Hello world."

		returnMsg := adapter.PrintStore()
		if returnMsg != exp {
			t.Errorf("Message didn't match: got %q and want %q\n", returnMsg, exp)
		}
	})

	t.Run("adapter new printer", func(t *testing.T) {
		adapter := PrinterAdapter{
			OldPrinter: nil,
			Msg:        msg,
		}

		returnMsg := adapter.PrintStore()
		if returnMsg != msg {
			t.Errorf("Message didn't match: %s\n", returnMsg)
		}
	})
}
