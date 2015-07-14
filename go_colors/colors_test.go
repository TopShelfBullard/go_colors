package colors_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/TopShelfBullard/go_colors/colors"
)

func TestRed(t *testing.T) {
	Convey("The Red method returns the expected string", t, func() {
	expected :=  "\033[31mRed\033[0m"
		actual := colors.Red("Red")
		So(actual, ShouldEqual, expected)
	})
}

func TestGreen(t *testing.T) {
	Convey("The Green method returns the expected string", t, func() {
	expected :=  "\033[32mGreen\033[0m"
		actual := colors.Green("Green")
		So(actual, ShouldEqual, expected)
	})
}

func TestYellow(t *testing.T) {
	Convey("The Yellow method returns the expected string", t, func() {
	expected :=  "\033[33mYellow\033[0m"
		actual := colors.Yellow("Yellow")
		So(actual, ShouldEqual, expected)
	})
}

func TestPurple(t *testing.T) {
	Convey("The Purple method returns the expected string", t, func() {
	expected :=  "\033[34mPurple\033[0m"
		actual := colors.Purple("Purple")
		So(actual, ShouldEqual, expected)
	})
}

func TestViolet(t *testing.T) {
	Convey("The Violet method returns the expected string", t, func() {
	expected :=  "\033[35mViolet\033[0m"
		actual := colors.Violet("Violet")
		So(actual, ShouldEqual, expected)
	})
}

func TestCyan(t *testing.T) {
	Convey("The Cyan method returns the expected string", t, func() {
	expected :=  "\033[36mCyan\033[0m"
		actual := colors.Cyan("Cyan")
		So(actual, ShouldEqual, expected)
	})
}

func TestBrightRed(t *testing.T) {
	Convey("The BrightRed method returns the expected string", t, func() {
	expected :=  "\033[91mBrightRed\033[0m"
		actual := colors.BrightRed("BrightRed")
		So(actual, ShouldEqual, expected)
	})
}

func TestBrightGreen(t *testing.T) {
	Convey("The BrightGreen method returns the expected string", t, func() {
	expected :=  "\033[92mBrightGreen\033[0m"
		actual := colors.BrightGreen("BrightGreen")
		So(actual, ShouldEqual, expected)
	})
}

func TestBrightYellow(t *testing.T) {
	Convey("The BrightYellow method returns the expected string", t, func() {
	expected :=  "\033[93mBrightYellow\033[0m"
		actual := colors.BrightYellow("BrightYellow")
		So(actual, ShouldEqual, expected)
	})
}

func TestBrightPurple(t *testing.T) {
	Convey("The BrightPurple method returns the expected string", t, func() {
	expected :=  "\033[94mBrightPurple\033[0m"
		actual := colors.BrightPurple("BrightPurple")
		So(actual, ShouldEqual, expected)
	})
}

func TestBrightViolet(t *testing.T) {
	Convey("The BrightViolet method returns the expected string", t, func() {
		expected := "\033[95mBrightViolet\033[0m"
		actual := colors.BrightViolet("BrightViolet")
		So(actual, ShouldEqual, expected)
	})
}

func TestBrightCyan(t *testing.T) {
	Convey("The BrightCyan method returns the expected string", t, func() {
		expected := "\033[96mBrightCyan\033[0m"
		actual := colors.BrightCyan("BrightCyan")
		So(actual, ShouldEqual, expected)
	})
}