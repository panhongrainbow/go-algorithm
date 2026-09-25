package bpTree

// =====================================================================================================================
//                  ⚗️ Consistency Integrity Test ( [B Plus Tree] ) - B加树 主要测试
// =====================================================================================================================
// 🧪 B Plus Tree unit test validates structure via bulk insert and delete.
// 🧪 Inserts large data, then deletes all to check if tree resets to empty.
// 🧪 Indexing errors may cause data loss or deletion failures.
// 🧪 Ensures indexing accuracy for reliable operations.

// To run the test, run the following command:
//
// cd /home/panhong/go/src/github.com/panhongrainbow/go-algorithm/bptree
// go clean -cache
// go test -v . -timeout=0 -run Test_Check_BpTree_Accuracies

// =====================================================================================================================

import (
	"testing"

	"github.com/panhongrainbow/go-algorithm/utilhub"
	"github.com/stretchr/testify/require"
)

// Test_Check_BpTree_Accuracy 🧫 verifies B Plus Tree indexing correctness through (1)bulk insert/delete,
// (2)randomized boundary testing, and (3)single-node endurance testing.
func Test_Check_BpTree_Accuracies(t *testing.T) {
	// This is a set of path configurations shared by the automated testing.
	var (
		// 🧪 Create a config instance for B plus tree unit testing and parse default values.
		autoTestConfig = utilhub.GetAutoConfig()

		// 🧪 Navigate to the project dataSet directory for test record storage.
		ProjectDir = utilhub.FileNode{}.Goto(autoTestConfig.Record.TestRecordPath)

		// 🧪 Create a subdirectory named with the current date under the project.
		recordDir = ProjectDir.MkDir(autoTestConfig.Record.ManualRecordDate)
	)

	// Perform pre-test checks to ensure all required record paths are properly initialized before running the test.
	t.Run("Pre-Test Checks", func(t *testing.T) {
		// Record path must not be empty.
		require.NotEqual(t, "", ProjectDir.Path(), "record path is empty; check path creation")

		// Record subdirectory must not be empty.
		require.NotEqual(t, "", recordDir.Path(), "record date path is empty; check path creation")
	})

	// Bulk insert/delete performs a sudden bulk insertion followed by bulk deletion of data in the B Plus Tree.「简单的大量资料新增和」
	t.Run("Bulk Insert/Delete", func(t *testing.T) {
		// Prepare test data for BulkInsertDelete.
		prepareBulkInsertDelete(t, recordDir, autoTestConfig)

		// Verify test data for BulkInsertDelete.
		verifyBulkInsertDelete(t, recordDir, autoTestConfig)

		// Execute accuracy test for BulkInsertDelete.
		runBulkInsertDelete(t, recordDir, autoTestConfig)
	})

	// RandomizedBoundary performs repeated insertion and deletion of data at the boundaries of the B Plus Tree,
	// including randomized cases where newly inserted keys collide with existing index keys.「新增资料的键值与既有索引发生冲突」
	t.Run("Randomized Boundary Test", func(t *testing.T) {

		// Prepare test data for RandomizedBoundary.
		prepareRandomizedBoundary(t, recordDir)

		// Verify test data for RandomizedBoundary.
		verifyRandomizedBoundary(t, recordDir, autoTestConfig)

		// Execute accuracy test for RandomizedBoundary.
		runRandomizedBoundary(t, recordDir, autoTestConfig)
	})

	// RedundantOperation increases the B Plus Tree size across different scales and repeatedly inserts and deletes the same key,
	// ensuring that redundant operations do not cause structural errors or inconsistencies.「单点疲劳测试」
	t.Run("Single Node Endurance Test", func(t *testing.T) {

		// Prepare test data for SingleNodeEndurance.
		prepareSingleNodeEndurance(t, recordDir)

		// Verify test data for SingleNodeEndurance.
		verifySingleNodeEndurance(t, recordDir, autoTestConfig)

		// Execute accuracy test for SingleNodeEndurance.
		runSingleNodeEndurance(t, recordDir, autoTestConfig)

	})
}
