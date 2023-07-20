package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ilmsg/cafe/di/runner"
)

type mockDatabase struct {
	setErr error
	getErr error
	getStr string
}

func (m *mockDatabase) Get(string) (string, error) {
	return m.getStr, m.getErr
}

func (m *mockDatabase) Set(string, string) error {
	return m.setErr
}

func TestRunnerArgsErr(t *testing.T) {
	r := runner.NewRunner(&mockDatabase{})
	if err := r.Run(&bytes.Buffer{}, []string{}); err == nil {
		t.Error("expected error on empty slice for args, got nil")
	}
}
func TestRunnerUsageErr(t *testing.T) {
	r := runner.NewRunner(&mockDatabase{})
	if err := r.Run(&bytes.Buffer{}, []string{"./kv", "help", "123"}); err == nil {
		t.Error("expected error on empty slice for args, got nil")
	}
}

func TestRunnerSetMissingArgErr(t *testing.T) {
	r := runner.NewRunner(&mockDatabase{})
	if err := r.Run(&bytes.Buffer{}, []string{"./kv", "set", "bob"}); err == nil {
		t.Error("expected error on empty slice for args, go nil")
	}
}

func TestRunnerReturnErrOnSet(t *testing.T) {
	setErr := errors.New("set err")
	r := runner.NewRunner(&mockDatabase{setErr: setErr})
	err := r.Run(&bytes.Buffer{}, []string{"./kv", "set", "bob", "10"})
	if err == nil {
		t.Error("expected error on empty slice for args, go nil")
	}
	if err.Error() != setErr.Error() {
		t.Errorf("expected error to be %v, got %v", setErr, err)
	}
}

func TestRunnerReturnErrOnGet(t *testing.T) {
	getErr := errors.New("get err")
	r := runner.NewRunner(&mockDatabase{getErr: getErr, getStr: "10"})
	err := r.Run(&bytes.Buffer{}, []string{"./kv", "get", "bob"})
	if err == nil {
		t.Error("expected error on empty slice for args, got nil")
	}
	if err.Error() != getErr.Error() {
		t.Errorf("expected error to be %v, got %v", getErr, err)
	}
}

func TestRunnerExpectedOutput(t *testing.T) {
	r := runner.NewRunner(&mockDatabase{getStr: "10"})
	buf := &bytes.Buffer{}
	err := r.Run(buf, []string{"./kv", "get", "bob"})
	if err != nil {
		t.Error("expected error to be nil on mock db get returning strings")
	}
	if buf.String() != "10\n" {
		t.Errorf("expected buffer to be 10 got %s", buf.String())
	}
}
