package utilhub

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// AutoFinalReport ⛏️ generates and prints a detailed progress report in a formatted table.
// rowWidth: The width of the value column in the table. It is based on the longest value length of each row.
func AutoFinalReport(pb *ProgressBar, rowWidth int) error {
	// If the progress is not finished, return an error message.
	if !pb.complete {
		return errors.New("progress is not yet complete")
	}

	// Calculate the total time that has elapsed between the start and the end.
	elapsed := pb.endTime.Sub(pb.startTime)

	// Define fixed widths for the table's fields and values to ensure proper alignment.
	fieldWidth := 20
	if rowWidth < 32 {
		rowWidth = 32 // The valueWidth should be at least 30 because the date and time are included.
	}
	totalWidth := fieldWidth + rowWidth + 7
	// Create a border for the table using a repeated pattern for visual clarity.
	border := BrightYellow + strings.Repeat("=", totalWidth) + Reset
	divider := BrightYellow + strings.Repeat("-", totalWidth) + Reset

	// Print the report title centered within the table, using padding to adjust its position.
	title := "Progress Bar Report"
	titleWidth := len(title)
	padding := (totalWidth - titleWidth) / 2
	fmt.Println(BrightMagenta + border + Reset)
	fmt.Printf("%s|%s%s%s|%s\n", BrightMagenta, strings.Repeat(" ", padding), title, strings.Repeat(" ", padding-1), Reset)
	fmt.Println(BrightMagenta + border + Reset)

	// Print the table header, highlighting the column titles for "Field" and "Value".
	fmt.Printf("%s| %-*s | %-*s |%s\n", BrightRed, fieldWidth, "Field", rowWidth, "Value", Reset)
	fmt.Println(BrightRed + divider + Reset)

	// Print each row of the table with the task's details, formatted to align fields and values.
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Task Name", rowWidth, pb.name, Reset) // %-*s ensures left alignment.
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Start Time", rowWidth, pb.startTime.Format(time.RFC1123), Reset)
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "End Time", rowWidth, pb.endTime.Format(time.RFC1123), Reset)
	fmt.Printf("%s| %-*s | %-*s |%s\n", DarkYellow, fieldWidth, "Elapsed Time", rowWidth, elapsed.String(), Reset)
	fmt.Printf("%s| %-*s | %-*d |%s\n", DarkYellow, fieldWidth, "Total Tasks", rowWidth, pb.total, Reset)
	fmt.Printf("%s| %-*s | %-*d |%s\n", DarkYellow, fieldWidth, "Completed Tasks", rowWidth, pb.currentProcess, Reset)

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
// rowWidth: The width of the value column in the table. It is based on the longest value length of each row.
func FailScenarioPreview(each TestProcessConfigType, rowWidth int) {
	// Test File
	fmt.Println(each.Record.ManualRecordFile)
	// Test Date
	// fmt.Println(each.Record.ManualRecordDate)
	// Test File Path
	// fmt.Println(each.Record.TestRecordPath)
	// Test Data Size
	// fmt.Println(each.Parameters.RandomTotalCount)
	// B Plus Tree Width
	// fmt.Println(each.Parameters.BpWidth[0])
}
