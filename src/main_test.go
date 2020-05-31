/*
Copyright 2020 The Worker-Monitor-Client Author.
Licensed under the GNU General Public License v3.0.
    https://github.com/MfsTeller/worker-monitor-client/blob/master/LICENSE
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	ValidCase   = "[ValidCase] "
	InvalidCase = "[InvalidCase] "
)

func SetExiter(testName string, t *testing.T) {
	if strings.Contains(testName, ValidCase) {
		exiter = func(exitCode int) {
			t.Errorf("Test exited")
		}
	} else {
		exiter = func(exitCode int) {
			t.Skip("Test skipped")
		}
	}
}

// TestMain
func TestMain(m *testing.M) {
	// Initial setting
	exiter = func(exitCode int) {
		fmt.Println("Exit code =", exitCode)
	}

	fmt.Println("=== Start testing")
	exitCode := m.Run()
	fmt.Println("=== End testing")
	os.Exit(exitCode)
}

// 01: -d option is specified
// 02: -d option is not specified
func Test_parseFlag(t *testing.T) {
	tests := []struct {
		name   string
		osArgs []string
	}{
		{
			name:   ValidCase + "-d option",
			osArgs: []string{"worker-monitor", Run, "-d", "2020/04/30"},
		},
		{
			name:   ValidCase + "without -d option",
			osArgs: []string{"worker-monitor", Run},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.osArgs
			parseFlag()
		})
	}
}

// 01: setup subcommand is specified
// 02: unsetup subcommand is specified
// 03: run subcommand is specified
// 04: get subcommand is specified
// 05: post subcommand is specified
// 06: subcommand is not specified
func Test_osArgsValidation(t *testing.T) {
	tests := []struct {
		name   string
		osArgs []string
	}{
		{
			name:   ValidCase + Setup,
			osArgs: []string{"worker-monitor", Setup},
		},
		{
			name:   ValidCase + Unsetup,
			osArgs: []string{"worker-monitor", Unsetup},
		},
		{
			name:   ValidCase + Run,
			osArgs: []string{"worker-monitor", Run},
		},
		{
			name:   ValidCase + Get,
			osArgs: []string{"worker-monitor", Get},
		},
		{
			name:   ValidCase + Post,
			osArgs: []string{"worker-monitor", Post},
		},
		{
			name:   InvalidCase + "subcommand is not specified",
			osArgs: []string{"worker-monitor"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetExiter(tt.name, t)
			os.Args = tt.osArgs
			osArgsValidation()
		})
	}
}

// 01: Setup
func Test_setup(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: ValidCase + Setup,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetExiter(tt.name, t)
			setup()
		})
	}
}

// 01: Unsetup
func Test_unsetup(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: ValidCase + Unsetup,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetExiter(tt.name, t)
			unsetup()
		})
	}
}

func Test_get(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get()
		})
	}
}

// 01: run
func Test_run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: ValidCase + Run,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run()
		})
	}
}

func Test_post(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name   string
		osArgs []string
	}{
		{
			name:   ValidCase + Setup,
			osArgs: []string{"worker-monitor", Setup},
		},
		{
			name:   ValidCase + Unsetup,
			osArgs: []string{"worker-monitor", Unsetup},
		},
		{
			name:   ValidCase + Run,
			osArgs: []string{"worker-monitor", Run},
		},
		{
			name:   ValidCase + "run with -d option",
			osArgs: []string{"worker-monitor", Run, "-d", "2020/04/30"},
		},
		{
			name:   ValidCase + "invalid argument",
			osArgs: []string{"worker-monitor", "xxx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.osArgs
			main()
		})
	}
}
