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

	BeforeEach(func() {
		fakeGithubClient = &fakes.GithubClient{
			VersionsToReturn: "222222",
		}
	})

	It("Returns an empty response when current version is latest SHA", func() {
		checkRequest.Version = concourse.Version{Number: "222222"}
		checkCommand := check.NewCheckCommand(fakeGithubClient)
		checkResponse, err := checkCommand.Execute(checkRequest)

		Ω(err).ShouldNot(HaveOccurred())
		Ω(checkResponse).Should(BeEmpty())

	})

	It("Returns the HEAD commit sha", func() {
		checkRequest.Version = concourse.Version{Number: "111111"}
		checkCommand := check.NewCheckCommand(fakeGithubClient)
		checkResponse, err := checkCommand.Execute(checkRequest)
		Expect(err).ToNot(HaveOccurred())
		Ω(checkResponse).Should(Equal(check.CheckResponse{
			concourse.Version{Number: "222222"}}))
	})
})
