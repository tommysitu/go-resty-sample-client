package interview_accountapi

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInterviewAccountapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InterviewAccountapi Suite")
}
