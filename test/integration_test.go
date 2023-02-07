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
	programPath := os.Getenv("ABS_EXEC_PATH")
	if programPath == "" {
		Fail("The program path is not set, the tests should be run by 'make test'.")
	}

	testdataDir := "./testdata/"
	var command *exec.Cmd
	var errCmd error

	Describe("Check validate command", func() {
		var jsonFilePath string
		var output []byte

		Context("on a valid json file", func() {
			BeforeEach(func() {
				jsonFilePath = testdataDir + "base.json"
				command = exec.Command(programPath, "validate", jsonFilePath)
				output, errCmd = command.Output()
			})
			It("should pass validation", func() {
				Expect(errCmd).NotTo(HaveOccurred())
				Expect(output).To(ContainSubstring("validated successfully"))
			})
		})

		Context("on a corrupted json file", func() {
			BeforeEach(func() {
				jsonFilePath = testdataDir + "corrupted.json"
				command = exec.Command(programPath, "validate", jsonFilePath)
				errCmd = command.Run()
			})
			It("should fail with error", func() {
				Expect(errCmd).To(HaveOccurred())
			})
		})
	})

	Describe("Check compile command", func() {
		var basePath string
		var outputPath string
		var dataExpected, dataOutput interface{}
		var errReadExpected, errReadOutput error

		Context("on a valid base", func() {
			BeforeEach(func() {
				basePath = testdataDir + "base.json"
				outputPath = testdataDir + "output_only_base.json"
				command = exec.Command(programPath, "compile", "-b", basePath, "-o", outputPath)
				errCmd = command.Run()
				dataExpected, errReadExpected = os.ReadFile(basePath)
				dataOutput, errReadOutput = os.ReadFile(outputPath)
			})
			It("should create the same base file", func() {
				Expect(errCmd).NotTo(HaveOccurred())
				Expect(errReadExpected).NotTo(HaveOccurred())
				Expect(errReadOutput).NotTo(HaveOccurred())
				Expect(reflect.DeepEqual(dataExpected, dataOutput)).To(BeTrue())
			})
		})

		Context("on a valid base with a layer", func() {
			var layerPath, expectedPath string

			BeforeEach(func() {
				basePath = testdataDir + "base.json"
				layerPath = testdataDir + "appinfo.yaml"
				outputPath = testdataDir + "output_base_with_layer.json"
				expectedPath = testdataDir + "expected_base_and_appinfo.json"
				command = exec.Command(programPath, "compile", "-b", basePath, "-o", outputPath, layerPath)
				errCmd = command.Run()
				dataExpected, errReadExpected = os.ReadFile(expectedPath)
				dataOutput, errReadOutput = os.ReadFile(outputPath)
			})
			It("should produce the same file as expected", func() {
				Expect(errCmd).NotTo(HaveOccurred())
				Expect(errReadExpected).NotTo(HaveOccurred())
				Expect(errReadOutput).NotTo(HaveOccurred())
				Expect(reflect.DeepEqual(dataExpected, dataOutput)).To(BeTrue())
			})
		})

		Context("on a corrupted base", func() {
			BeforeEach(func() {
				basePath = testdataDir + "corrupted.json"
				outputPath = testdataDir + "output_corrupted_base.json"
				command = exec.Command(programPath, "compile", "-b", basePath, "-o", outputPath)
				errCmd = command.Run()
				_, errReadOutput = os.ReadFile(outputPath)
			})
			It("should create the same base file", func() {
				Expect(errCmd).To(HaveOccurred())
				Expect(errReadOutput).To(HaveOccurred())
			})
		})
	})
})
