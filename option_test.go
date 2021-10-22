package main

import (
	"cliMemoApp/app"
	"testing"
)

func TestIsContainsArgs(t *testing.T) {
	res := app.IsContains([]string{
		"cmemo", "--height", "40%", "-v",
	}, []string{
		"-v",
		"--height",
	})
	if res {
		t.Errorf("Abs(true) = %t; want true", res)
	}
}

func TestIsContainsArgsLessArgs(t *testing.T) {
	res := app.IsContains([]string{
		"cmemo", "-v",
	}, []string{
		"-v",
		"--height",
	})
	if res {
		t.Errorf("Abs(true) = %t; want true", res)
	}
}

func TestIsContainsArgsFalse(t *testing.T) {
	res := app.IsContains([]string{
		"cmemo", "--height", "40%", "-v",
	}, []string{
		"--height",
	})
	if !res {
		t.Errorf("Abs(false) = %t; want false", res)
	}
}

func TestIsContainsArgsNowords(t *testing.T) {
	res := app.IsContains([]string{
		"cmemo", "--height", "40%", "-v",
	}, []string{})
	if !res {
		t.Errorf("Abs(false) = %t; want false", res)
	}
}
