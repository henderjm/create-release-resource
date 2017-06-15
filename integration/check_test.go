package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/henderjm/create-release-resource/check"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Check", func() {
	var (
		checkPath string
		err       error
		session   *gexec.Session
	)

	BeforeEach(func() {
		checkPath, err = gexec.Build("github.com/henderjm/create-release-resource/cmd/check")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return an empty array when latest version matches", func() {
		command := exec.Command(checkPath)
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, 5*time.Second).Should(gexec.Exit(0))
		reader := bytes.NewBuffer(session.Buffer().Contents())

		var response check.CheckResponse
		err := json.NewDecoder(reader).Decode(&response)
		fmt.Println(response)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(response).Should(BeEmpty())
	})
})
