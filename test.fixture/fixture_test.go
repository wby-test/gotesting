// 测试固件
// 测试数据的初始化和销毁，测试结果的可重复性
// 使用setUp/设置 tearDown/销毁
// 使用场景: 特定数据加载到DB，完成后清除； 复制一组特定文件，完成后删除； 数据mock，设置期望结果
package test_fixture

import (
	"fmt"
	"testing"
)

func setUp(testName string) func() {
	fmt.Printf("\tsetUp fixture for %s\n", testName)
	return func() {
		fmt.Printf("\ttearDown fixture for %s\n", testName)
	}
}

func TestFunc1(t *testing.T) {
	defer setUp(t.Name())()
	fmt.Printf("\tExecute test: %s\n", t.Name())
}

// 测试固件销毁的原生支持
func TestFunc2(t *testing.T) {
	t.Cleanup(setUp(t.Name()))
	fmt.Printf("\tExecute test: %s\n", t.Name())
}

// ----------------包级别测试固件--------------
func pkgSetup(pkgName string) func() {
	fmt.Printf("package setUp fixture for %s\n", pkgName)
	return func() {
		fmt.Printf("package  TearDown fixture for %s\n", pkgName)
	}
}

func TestMain(m *testing.M) {
	defer pkgSetup("package test_fixture")()
	m.Run()
}

// --------------测试套件 + 测试固件------------
func suiteSetup(suitName string) func() {
	fmt.Printf("\tsetUp fixture  for suite %s\n", suitName)
	return func() {
		fmt.Printf("\ttearDown fixture for suite %s\n", suitName)
	}
}

func testCase1(t *testing.T) {
	fmt.Printf("\tExecute test: %s\n", t.Name())
}

func testCase2(t *testing.T) {
	fmt.Printf("\tExecute test: %s\n", t.Name())
}

func TestFuncSuite(t *testing.T) {
	t.Cleanup(suiteSetup(t.Name()))
	t.Run("testCase1", testCase1)
	t.Run("testCase2", testCase2)
}
