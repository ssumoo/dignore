package list_backend_test

import (
	"testing"

	"github/ssumoo/dignore/list_backend"
)

func testCheck(path string, line string, ans list_backend.MatchMode, t *testing.T) {
	ret := list_backend.CheckPathAgainstLine(path, line)
	if ret != ans {
		t.Errorf("CheckPathAgainstLine(%s, %s) -> %s != %s", path, line, ret.String(), ans.String())
	}
}

func TestCheckExclude(t *testing.T) {
	testCheck(
		"main.go",
		"*.go",
		list_backend.MatchExclude,
		t,
	)
}

func TestCheckInclude(t *testing.T) {
	testCheck(
		"main.go",
		"!main.go",
		list_backend.MatchInclude,
		t,
	)
}

func TestCheckNone(t *testing.T) {
	testCheck(
		"main.go",
		"abcde",
		list_backend.MatchNone,
		t,
	)
}

func TestOneLevelStarExclude(t *testing.T) {
	testCheck(
		"aaa/bbb/ccc.txt",
		"*/bbb",
		list_backend.MatchExclude,
		t,
	)
}

func TestOneLevelStarInclude(t *testing.T) {
	testCheck(
		"aaa/bbb/ccc.txt",
		"*bbb",
		list_backend.MatchNone,
		t,
	)
}

func TestImmediateDirExclude(t *testing.T) {
	testCheck(
		".git/logs/refs/heads/25-my-lovely-branch",
		".git/",
		list_backend.MatchExclude,
		t,
	)
}

func TestImmediateStarExclude(t *testing.T) {
	testCheck(
		".git",
		"*",
		list_backend.MatchExclude,
		t,
	)
}

func TestNestedStarExclude(t *testing.T) {
	testCheck(
		".git/logs/refs/heads/25-my-lovely-branch",
		"*",
		list_backend.MatchExclude,
		t,
	)
}

func TestImplicitDirInclude(t *testing.T) {
	testCheck(
		"services/my_random_service/file.foo",
		"my_random_service/",
		list_backend.MatchNone,
		t,
	)
}

func TestImplicitStarInclude(t *testing.T) {
	testCheck(
		"services/my_random_service/file.foo",
		"*my_random_service",
		list_backend.MatchNone,
		t,
	)
}

func TestNaiveDoubleStarMatch(t *testing.T) {
	testCheck(
		"services/dir1/name/some_files.txt",
		"**/name/",
		list_backend.MatchExclude,
		t,
	)
}

func TestTwoDoubleStarMatch(t *testing.T) {
	testCheck(
		"services/dir1/name/some_files.txt",
		"**/name/**",
		list_backend.MatchExclude,
		t,
	)
}

func TestTwoDoubleStarMatch2(t *testing.T) {
	testCheck(
		"services/dir1/name/some_files.txt",
		"**/non_existent/**",
		list_backend.MatchNone,
		t,
	)
}

func TestDoubleStarMatchAtRoot(t *testing.T) {
	testCheck(
		"services/dir1/name/some_files.txt",
		"**/some_files.txt",
		list_backend.MatchExclude,
		t,
	)
}

func TestDoubleStarFlieExtensionMatch(t *testing.T) {
	testCheck(
		"services/dir1/name/some_files.txt",
		"**/*.txt",
		list_backend.MatchExclude,
		t,
	)
}

func TestDeepInSubdirExclude(t *testing.T) {
	testCheck(
		"services/dir1/dir2/aaa_sub_string_abc.bar",
		"**/*sub_string*",
		list_backend.MatchExclude,
		t,
	)
}

func TestImplicitFileExtensionInclude(t *testing.T) {
	testCheck(
		"services/dir1/dir2/aaa_sub_string_abc.bar",
		"*.bar",
		list_backend.MatchNone,
		t,
	)
}

func TestOneLetterExclude(t *testing.T) {
	testCheck(
		"abcde.txt",
		"abcd?.txt",
		list_backend.MatchExclude,
		t,
	)
}
