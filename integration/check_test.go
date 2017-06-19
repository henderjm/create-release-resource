package integration_test

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"time"

	"github.com/henderjm/create-release-resource/check"
	"github.com/henderjm/create-release-resource/concourse"
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
		checkRequest := check.CheckRequest{
			Version: concourse.Version{Number: "999d8f8554db612b00cf26637686a062f28fcee4"}}

		stdin := &bytes.Buffer{}
		err := json.NewEncoder(stdin).Encode(checkRequest)
		Expect(err).NotTo(HaveOccurred())

		command := exec.Command(checkPath)
		command.Stdin = stdin

		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, 5*time.Second).Should(gexec.Exit(0))
		reader := bytes.NewBuffer(session.Buffer().Contents())

		response := check.CheckResponse{}
		err = json.NewDecoder(reader).Decode(&response)

		Ω(err).ShouldNot(HaveOccurred())
		Ω(response).Should(BeEmpty())
	})

	It("should return an array of versions when there are currently none", func() {
		checkRequest := check.CheckRequest{
			Version: concourse.Version{Number: "736b8f1dda544789fd194fd3775d71ea58095e6d"},
		}
		stdin := &bytes.Buffer{}
		err := json.NewEncoder(stdin).Encode(checkRequest)
		Expect(err).NotTo(HaveOccurred())

		command := exec.Command(checkPath)
		command.Stdin = stdin

		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, 5*time.Second).Should(gexec.Exit(0))
		reader := bytes.NewBuffer(session.Buffer().Contents())

		response := check.CheckResponse{}
		err = json.NewDecoder(reader).Decode(&response)

		Expect(err).NotTo(HaveOccurred())
		Ω(response).Should(Equal(check.CheckResponse{concourse.Version{Number: "999d8f8554db612b00cf26637686a062f28fcee4"}}))
	})
})
