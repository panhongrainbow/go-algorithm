package bpTree

import (
	"strings"
	"testing"

	"github.com/panhongrainbow/go-algorithm/utilhub"
	"github.com/stretchr/testify/require"
)

// Test_Check_Manual_Accuracies 🧫 is a collection of previously failed cases from automated testing,
// used to repeatedly reproduce and verify the reported issues.
func Test_Check_Manual_Accuracies(t *testing.T) {

	// 🧪 Load the configuration containing manually collected failed test cases for re-verification.
	var manualTestConfig = utilhub.GetManualConfig()

	// Iterate over failed test cases and ignore the index.
	for _, each := range manualTestConfig {
		var (
			// 🧪 Navigate to the project dataSet directory for test record storage.
			ProjectDir = utilhub.FileNode{}.Goto(each.Record.TestRecordPath)

			// 🧪 Create a subdirectory named with the date under the project.
			recordDir = ProjectDir.MkDir(each.Record.ManualRecordDate)
		)

		// Verify that the record file exists.
		if _, err := recordDir.CheckFile(each.Record.ManualRecordFile); err != nil {
			require.NoError(t, err)
		}

		// The following tests cover different aspects of B Plus tree correctness and stability:
		switch {

		// Bulk insert/delete performs a sudden bulk insertion followed by bulk deletion of data in the B Plus Tree.「简单的大量资料新增和」
		case strings.Contains(each.Record.ManualRecordFile, "Bulk Insert/Delete"):

			// Verify test data for BulkInsertDelete.
			verifyBulkInsertDelete(t, recordDir, each)

			// Execute accuracy test for BulkInsertDelete.
			runBulkInsertDelete(t, recordDir, each)

		// RandomizedBoundary performs repeated insertion and deletion of data at the boundaries of the B Plus Tree,
		// including randomized cases where newly inserted keys collide with existing index keys.「新增资料的键值与既有索引发生冲突」
		case strings.Contains(each.Record.ManualRecordFile, "Randomized Boundary Test"):

			// Verify test data for RandomizedBoundary.
			verifyRandomizedBoundary(t, recordDir, each)

			// Execute accuracy test for RandomizedBoundary.
			runRandomizedBoundary(t, recordDir, each)

		// RandomizedBoundary performs repeated insertion and deletion of data at the boundaries of the B Plus Tree,
		// including randomized cases where newly inserted keys collide with existing index keys.「新增资料的键值与既有索引发生冲突」
		case strings.Contains(each.Record.ManualRecordFile, "Single Node Endurance Test"):

			// Verify test data for SingleNodeEndurance.
			verifySingleNodeEndurance(t, recordDir, each)

			// Execute accuracy test for SingleNodeEndurance.
			runSingleNodeEndurance(t, recordDir, each)
		}
	}

	return
}
