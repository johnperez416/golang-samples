// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func init() {
	os.Setenv("SLEEP_MS", "0")
	os.Setenv("TASK_NUM", "1")
	os.Setenv("ATTEMPT_NUM", "1")
	os.Setenv("FAIL_RATE", "")
}

func TestSuccessfulJob(t *testing.T) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	main()

	writer.Flush()
	output := buf.String()

	start := fmt.Sprintf("Starting Task #1, Attempt #1 ...")
	finish := fmt.Sprintf("Completed Task #1, Attempt #1")

	if !(strings.Contains(start, output) && strings.Contains(finish, output)) {
		t.Fatalf("\nExpected string to contain:\n%s\n%s\nGot: %s", start, finish, output)
	}
}

// func TestFailingJob(t *testing.T) {
// 	os.Setenv("FAIL_RATE", "0.9999999")
// 	main()
// }
