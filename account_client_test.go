package interview_accountapi

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
)

var _ = Describe("AccountsClient", func() {

	var client *AccountClient

	BeforeEach(func() {
		client = NewAccountClient()
	})

	AfterEach(func() {
		_ = client.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
	})

	Describe("Create", func() {
		Context("With mandatory attributes", func() {
			It("should return a new account without error", func() {
				resource, err := client.Create(&AccountParams{
					Attributes: Attributes{
						Country: "GB",
					},
					ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
					OrganizationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Type:           "accounts",
				})
				Expect(err).To(BeNil())
				Expect(resource).NotTo(BeNil())
			})
		})
	})

	Describe("Delete", func(){
		Context("An empty account", func() {
			It("should not fail", func() {
				err := client.Delete(uuid.NewV4().String(), 0)
				Expect(err).To(BeNil())
			})
		})

		Context("An existing account", func() {

			JustBeforeEach(func() {
				_, _ = client.Create(&AccountParams{
					Attributes: Attributes{
						Country: "GB",
					},
					ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
					OrganizationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Type:           "accounts",
				})
			})

			It("should not fail", func() {
				err := client.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
				Expect(err).To(BeNil())
			})
		})
	})
})
