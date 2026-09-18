package bpTree

import (
	"strings"
	"testing"

	"github.com/panhongrainbow/go-algorithm/utilhub"
	"github.com/stretchr/testify/require"
)

func Test_Check_Manual_Accuracies(t *testing.T) {
	// This is a set of path configurations shared by the automated testing.
	var (
		// 🧪 Create a config instance for B plus tree unit testing and parse default values.
		manualTestConfig = utilhub.GetManualConfig()
	)

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
		case strings.Contains(each.Record.ManualRecordFile, "BulkInsertDelete"):
			// Basic test: bulk insert a large amount of data into the B Plus tree,
			// then delete a large amount of data.

			// Verify test data for BulkInsertDelete.
			verifyBulkInsertDelete(t, recordDir, each)

			// Execute accuracy test for BulkInsertDelete.
			runBulkInsertDelete(t, recordDir, each)

		case strings.Contains(each.Record.ManualRecordFile, "RandomizedBoundary"):
			// Boundary test: test cases where data keys may collide with index keys.

			// Verify test data for RandomizedBoundary.
			verifyRandomizedBoundary(t, recordDir, each)

			// Execute accuracy test for RandomizedBoundary.
			runRandomizedBoundary(t, recordDir, each)

		case strings.Contains(each.Record.ManualRecordFile, "SingleNodeEndurance"):
			// Endurance test: repeatedly delete the same sequence of data in a randomized order,
			// verifying that the B Plus tree does not encounter node failures or become corrupted.

			// Verify test data for SingleNodeEndurance.
			verifySingleNodeEndurance(t, recordDir, each)

			// Execute accuracy test for SingleNodeEndurance.
			runSingleNodeEndurance(t, recordDir, each)
		}
	}

	return
}
