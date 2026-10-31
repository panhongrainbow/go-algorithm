package utilhub

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Since AutoFinalReport and FailScenarioPreview are unlikely to be used elsewhere in the future,
// I would prefer to keep the logic in place rather than extract it just for the sake of reusability.
// (就这样，以后不太会重复使用到，不再优化或做成共用函式)

// AutoFinalReport ⛏️ generates and prints a detailed progress report in a formatted table.
// valueWidth: The width of the value column in the table. It is based on the longest value length of each row.
func AutoFinalReport(pb *ProgressBar, valueWidth int) error {
	// If the progress is not finished, return an error message.
	if !pb.complete {
		return errors.New("progress is not yet complete")
	}

	// Calculate the total time that has elapsed between the start and the end.
	elapsed := pb.endTime.Sub(pb.startTime)

	// Define fixed widths for the table's fields and values to ensure proper alignment.
	fieldWidth := 20
	if valueWidth < 32 {
		valueWidth = 32 // The valueWidth should be at least 32 for datetime-related field names.
	}

	barWidth := 7 // The bar between the field and value columns in the table is estimated to be 7 characters wide.

	totalWidth := fieldWidth + valueWidth + barWidth // totalWidth represents the combined width of the field, value, and separator.

	// Create a border for the table using a repeated pattern for visual clarity.
	border := BrightYellow + strings.Repeat("=", totalWidth) + Reset
	divider := BrightYellow + strings.Repeat("-", totalWidth) + Reset

	// Print the report title centered within the table, using padding to adjust its position.
	title := "Progress Bar Report"
	titleWidth := len(title)
	padding := (totalWidth - titleWidth) / 2
	fmt.Println(BrightMagenta + border + Reset)
	// The first and last %s are used to apply and reset the text color.
	fmt.Printf("%s|%s%s%s|%s\n", BrightMagenta, strings.Repeat(" ", padding), title, strings.Repeat(" ", padding-1), Reset)
	fmt.Println(BrightMagenta + border + Reset)

	// Print the table header, highlighting the column titles for "Field" and "Value".
	// %-*s ensures left alignment and requires two arguments: the field width and the string value.
	fmt.Printf("%s| %-*s | %-*s |%s\n", BrightRed, fieldWidth, "Field", valueWidth, "Value", Reset)
	fmt.Println(BrightRed + divider + Reset)

	// Print each row of the table with the task's details, formatted to align fields and values.
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Task Name", valueWidth, pb.name, Reset) // %-*s ensures left alignment.
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Start Time", valueWidth, pb.startTime.Format(time.RFC1123), Reset)
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "End Time", valueWidth, pb.endTime.Format(time.RFC1123), Reset)
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Elapsed Time", valueWidth, elapsed.String(), Reset)
	fmt.Printf("%s| %-*s | %-*d |%s\n", DarkYellow, fieldWidth, "Total Tasks", valueWidth, pb.total, Reset)
	fmt.Printf("%s| %-*s | %-*d |%s\n", DarkYellow, fieldWidth, "Completed Tasks", valueWidth, pb.currentProcess, Reset)

	// Print a closing border to signal the end of the report.
	fmt.Println(BrightMagenta + border + Reset)

	return nil
}

// FailScenarioPreview ⛏️ prints information about a recorded failure scenario
// to provide context for the error being tested.
// Failure scenarios focus on reproducibility, while performance testing
// is handled by automated tests that generate their own test data,
// allowing them to better reflect real-world conditions.
// (错误重现时不会记录时间)
// valueWidth: The width of the value column in the table. It is based on the longest value length of each row.
// (valueWidth 为每一行 栏位其值 的 长度)
func FailScenarioPreview(each TestProcessConfigType, valueWidth int) error {
	// Define fixed widths for the table's fields and values to ensure proper alignment.
	fieldWidth := 20
	if valueWidth < 32 {
		valueWidth = 32 // The valueWidth should be at least 32 for datetime-related field names.
	}

	barWidth := 7 // The bar between the field and value columns in the table is estimated to be 7 characters wide.

	totalWidth := fieldWidth + valueWidth + barWidth // totalWidth represents the combined width of the field, value, and separator.

	// Create a border for the table using a repeated pattern for visual clarity.
	border := BrightYellow + strings.Repeat("=", totalWidth) + Reset
	divider := BrightYellow + strings.Repeat("-", totalWidth) + Reset

	// Print the report title centered within the table, using padding to adjust its position.
	title := "Failure Scenario Preview"
	titleWidth := len(title)
	padding := (totalWidth - titleWidth) / 2
	fmt.Println(BrightMagenta + border + Reset)
	// The first and last %s are used to apply and reset the text color.
	fmt.Printf("%s|%s%s%s|%s\n", BrightMagenta, strings.Repeat(" ", padding), title, strings.Repeat(" ", padding-1), Reset)
	fmt.Println(BrightMagenta + border + Reset)

	// Print the table header, highlighting the column titles for "Field" and "Value".
	// %-*s ensures left alignment and requires two arguments: the field width and the string value.
	fmt.Printf("%s| %-*s | %-*s |%s\n", BrightRed, fieldWidth, "Field", valueWidth, "Value", Reset)
	fmt.Println(BrightRed + divider + Reset)

	// Print each row of the table with the task's details, formatted to align fields and values.
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Record File", valueWidth, each.Record.ManualRecordFile, Reset)
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Record Date", valueWidth, each.Record.ManualRecordDate, Reset)
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Record Path", valueWidth, each.Record.TestRecordPath, Reset)
	fmt.Printf("%s| %-*s | %-*d |%s\n", DarkYellow, fieldWidth, "Bptree Width", valueWidth, each.Parameters.BpWidth[0], Reset)
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Start Time", valueWidth, time.Now().Format(time.RFC1123), Reset)
	fmt.Printf("%s| %-*s | %-*d |%s\n", DarkYellow, fieldWidth, "Total Tasks", valueWidth, each.Parameters.RandomTotalCount, Reset)

	// Print a closing border to signal the end of the report.
	fmt.Println(BrightMagenta + border + Reset)

	return nil
}
