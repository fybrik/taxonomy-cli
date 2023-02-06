// Copyright 2023 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"os"
	"os/exec"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("Taxonomy-CLI integration tests", Ordered, func() {
	var programPath string
	var command *exec.Cmd
	var errCmd error

	BeforeAll(func() {
		programPath = os.Getenv("ABS_EXEC_PATH")
	})

	Describe("Check validate command", func() {
		var jsonFilePath string
		var output []byte

		Context("on a proper json file", func() {
			BeforeEach(func() {
				jsonFilePath = "./testdata/base.json"
				command = exec.Command(programPath, "validate", jsonFilePath)
				output, errCmd = command.Output()
			})
			It("should pass validation", func() {
				Expect(errCmd).To(BeNil())
				Expect(output).To(ContainSubstring("validated successfully"))
			})
		})

		Context("on a corrupted json file", func() {
			BeforeEach(func() {
				jsonFilePath = "./testdata/corrupted.json"
				command = exec.Command(programPath, "validate", jsonFilePath)
				errCmd = command.Run()
			})
			It("should fail with error", func() {
				Expect(errCmd).To(Not(BeNil()))
			})
		})
	})

	Describe("Check compile command", func() {
		var basePath string
		var outputPath string
		var dataExpected, dataOutput interface{}
		var errReadExpected, errReadOutput error

		Context("on a proper base", func() {
			BeforeEach(func() {
				basePath = "./testdata/base.json"
				outputPath = "./testdata/output_only_base.json"
				command = exec.Command(programPath, "compile", "-b", basePath, "-o", outputPath)
				errCmd = command.Run()
				dataExpected, errReadExpected = os.ReadFile(basePath)
				dataOutput, errReadOutput = os.ReadFile(outputPath)
			})
			It("should create the same base file", func() {
				Expect(errCmd).To(BeNil())
				Expect(errReadExpected).To(BeNil())
				Expect(errReadOutput).To(BeNil())
				Expect(reflect.DeepEqual(dataExpected, dataOutput)).To(BeTrue())
			})
		})

		Context("on a proper base with a layer", func() {
			var layerPath, expectedPath string

			BeforeEach(func() {
				basePath = "./testdata/base.json"
				layerPath = "./testdata/appinfo.yaml"
				outputPath = "./testdata/output_base_with_layer.json"
				expectedPath = "./testdata/expected_base_and_appinfo.json"
				command = exec.Command(programPath, "compile", "-b", basePath, "-o", outputPath, layerPath)
				errCmd = command.Run()
				dataExpected, errReadExpected = os.ReadFile(expectedPath)
				dataOutput, errReadOutput = os.ReadFile(outputPath)
			})
			It("should produce the same file as expected", func() {
				Expect(errCmd).To(BeNil())
				Expect(errReadExpected).To(BeNil())
				Expect(errReadOutput).To(BeNil())
				Expect(reflect.DeepEqual(dataExpected, dataOutput)).To(BeTrue())
			})
		})

		Context("on a corrupted base", func() {
			BeforeEach(func() {
				basePath = "./testdata/corrupted.json"
				outputPath = "./testdata/output_corrupted_base.json"
				command = exec.Command(programPath, "compile", "-b", basePath, "-o", outputPath)
				errCmd = command.Run()
				_, errReadOutput = os.ReadFile(outputPath)
			})
			It("should create the same base file", func() {
				Expect(errCmd).To(Not(BeNil()))
				Expect(errReadOutput).To(Not(BeNil()))
			})
		})
	})
})
