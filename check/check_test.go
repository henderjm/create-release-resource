package check_test

import (
	"github.com/henderjm/create-release-resource/check"
	"github.com/henderjm/create-release-resource/concourse"
	"github.com/henderjm/create-release-resource/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CheckCommand", func() {
	var (
		fakeGithubClient *fakes.GithubClient
		checkRequest     check.CheckRequest
	)

	It("Returns the HEAD commit sha", func() {
		checkRequest.Version = concourse.Version{Number: "111111"}
		checkCommand := check.NewCheckCommand(fakeGithubClient)
		_, err := checkCommand.Execute(checkRequest)
		Ω("124").Should(Equal("123"))
	})
})
