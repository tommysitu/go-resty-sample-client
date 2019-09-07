package interview_accountapi

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
)

var _ = Describe("AccountsClient", func() {

	const testAccountId = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	var client *AccountClient
	testAccount := AccountParams{
		Attributes: Attributes{
			Country: "GB",
		},
		ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganizationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
	}

	BeforeEach(func() {
		client = NewAccountClient()
	})

	AfterEach(func() {
		_ = client.Delete(testAccountId, 0)
	})

	Describe("Create", func() {
		Context("With mandatory attributes", func() {
			It("should return a new account without error", func() {
				resource, err := client.Create(testAccount)
				Expect(err).To(BeNil())
				Expect(resource.Data.ID).To(Equal(testAccountId))
				Expect(resource.Data.OrganizationID).To(Equal("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"))
				Expect(resource.Data.Attributes.Country).To(Equal("GB"))
				Expect(resource.Links).To(Equal(map[string]string{"self": "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}))
			})
		})
	})

	Describe("Fetch", func() {
		Context("A non-existent account", func() {
			It("should return error", func() {
				anyUUID := uuid.NewV4().String()
				resource, err := client.Fetch(anyUUID)

				Expect(err).To(Equal(fmt.Errorf("request failed []: record %s does not exist", anyUUID)))
				Expect(resource).To(BeNil())
			})
		})

		Context("An existing account", func() {
			JustBeforeEach(func() {
				_, _ = client.Create(testAccount)
			})

			It("should return account data", func() {
				resource, err := client.Fetch(testAccountId)

				Expect(err).To(BeNil())
				Expect(resource.Data.ID).To(Equal(testAccountId))
			})
		})
	})

	Describe("Delete", func(){
		Context("A non-existent account", func() {
			It("should not fail", func() {
				err := client.Delete(uuid.NewV4().String(), 0)
				Expect(err).To(BeNil())
			})
		})

		Context("An existing account", func() {

			JustBeforeEach(func() {
				_, _ = client.Create(testAccount)
			})

			It("should not fail", func() {
				err := client.Delete(testAccountId, 0)
				Expect(err).To(BeNil())
			})
		})
	})
})
