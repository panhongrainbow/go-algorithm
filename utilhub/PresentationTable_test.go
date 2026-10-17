package utilhub

import (
	"io"
	"os"
	"testing"
)

func TestPresentationTable(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	title := TestProcessConfigType{}
	title.Record.ManualRecordFile = "test"

	FailScenarioPreview(title, 1)

	w.Close()
	os.Stdout = old

	output, _ := io.ReadAll(r)
	r.Close()

	got := string(output)
	want := "test\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
