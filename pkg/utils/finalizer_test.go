// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRemoveFromFinalizer(t *testing.T) {
	g := NewGomegaWithT(t)

	finalizers := []string{"1", "2", "3", "4"}

	finalizers = RemoveFromFinalizer(finalizers, "2")
	g.Expect(finalizers).Should(Equal([]string{"1", "3", "4"}))
	finalizers = RemoveFromFinalizer(finalizers, "1")
	g.Expect(finalizers).Should(Equal([]string{"3", "4"}))
	finalizers = RemoveFromFinalizer(finalizers, "4")
	g.Expect(finalizers).Should(Equal([]string{"3"}))
	finalizers = RemoveFromFinalizer(finalizers, "3")
	g.Expect(finalizers).Should(Equal([]string{}))
}

func TestInsertFinalizer(t *testing.T) {
	g := NewGomegaWithT(t)

	finalizers := []string{"1"}
	InsertFinalizer(finalizers, "1")

	g.Expect(finalizers).Should(Equal([]string{"1"}))
}
