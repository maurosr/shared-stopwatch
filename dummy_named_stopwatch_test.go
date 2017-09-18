/*
Copyright (c) 2017, Mauro Schilman, courtesy Booking.com
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Package stopwatch implements a shared stopwatch
package stopwatch

import (
	"testing"
	"time"
)

func TestNewNamedDummy(t *testing.T) {
	s := NewDummyNamedStopwatch()
	if s == nil {
		t.Errorf("Error: NewStopwatch() returned: %v, expected: non-nil value", s)
	}
}

func TestNamedDummyIsBaseNamed(t *testing.T) {
	var s BaseNamedStopwatch
	s = NewDummyNamedStopwatch()

	if s == nil {
		t.Errorf("Error: NewStopwatch() returned: %v, expected: non-nil value", s)
	}
}

func TestAddAndDeleteNamedDummy(t *testing.T) {
	s := NewDummyNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	for _, name := range names {
		if s.Exists(name) {
			t.Errorf("Stopwatch %v should not have been added", name)
		}
	}

	s.Delete(names[0])
	if s.Exists(names[0]) {
		t.Errorf("Stopwatch %v was not successfully deleted", names[0])
	}
}

func TestStartNamedDummy(t *testing.T) {
	s := NewDummyNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	s.Start(names[0])
	elapsed := s.Elapsed(names[0])
	if elapsed != time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: 0", elapsed)
	}
}

func TestElapsedNamedDummy(t *testing.T) {
	s := NewDummyNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	elapsed0 := s.Elapsed(names[0])

	if elapsed0 != time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: 0", elapsed0)
	}

	stopCallback := s.Start(names[0])
	elapsed1 := s.Elapsed(names[0])

	if elapsed1 != time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: 0", elapsed0)
	}

	stopCallback()
	elapsed2 := s.Elapsed(names[0])
	elapsed3 := s.Elapsed(names[0])

	if elapsed2 < elapsed1 {
		t.Errorf("Error: Elapsed time is not monotonically incresing: elapsed2(%v) <= elapsed1(%v)", elapsed2, elapsed1)
	}

	if elapsed2 != elapsed3 {
		t.Errorf("Error: Elapsed time continues increasing after stop: elapsed2(%v) != elapsed3(%v)", elapsed2, elapsed3)
	}
}

func TestStopNamedDummy(t *testing.T) {
	s := NewDummyNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	stopCallback := s.Start(names[0])
	err := stopCallback()

	if err != nil {
		t.Error(err)
	}

	elapsed1 := s.Elapsed(names[0])
	elapsed2 := s.Elapsed(names[0])
	if elapsed1 != elapsed2 {
		t.Errorf("Error: Elapsed time continues increasing after stop: elapsed2(%v) != elapsed3(%v)", elapsed1, elapsed2)
	}

	err = stopCallback()

	if err != nil {
		t.Error(err)
	}
}

func TestResetNamedDummy(t *testing.T) {
	s := NewDummyNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	stopC := s.Start(names[0])

	err := s.Reset(names[0])
	if err != nil {
		t.Error(err)
	}

	stopC()
	err = s.Reset(names[0])
	if err != nil {
		t.Error(err)
	}
	elapsed := s.Elapsed(names[0])
	if elapsed != time.Duration(0) {
		t.Errorf("Error: elapsed time after reset %v, expected 0", elapsed)
	}
}

func TestIsRunningNamedDummy(t *testing.T) {
	s := NewDummyNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	if s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch running before started")
	}
	stopC := s.Start(names[0])
	if s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch should not be running")
	}

	stopC2 := s.Start(names[0])

	stopC()
	if s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch should not be running")
	}
	stopC2()
	if s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch did not stop after all stop callbacks executed")
	}
}
