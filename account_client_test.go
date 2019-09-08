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

	BeforeEach(func() {
		client = NewAccountClient()
	})

	AfterEach(func() {
		_ = client.Delete(testAccountId, 0)
	})

	Describe("Create", func() {
		Context("With mandatory attributes", func() {
			It("should return a new account without error", func() {
				resource, err := client.Create(getTestAccountParams(testAccountId))
				Expect(err).To(BeNil())
				Expect(resource.Data.ID).To(Equal(testAccountId))
				Expect(resource.Data.OrganizationID).To(Equal("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"))
				Expect(resource.Data.Attributes.Country).To(Equal("GB"))
				Expect(resource.Links).To(Equal(map[string]string{"self": "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}))
			})
		})

		Context("Without mandatory attributes", func() {
			It("should return an error", func() {
				resource, err := client.Create(AccountParams{
					OrganizationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Type:           "accounts",
				})
				Expect(err).To(Equal(fmt.Errorf("request failed []: validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body is required\nid in body is required")))
				Expect(resource).To(BeNil())
			})
		})
	})

	Describe("List", func() {
		Context("When no account is available", func() {
			It("should return empty data", func() {
				resources, err := client.List(PagingParams{})
				Expect(err).To(BeNil())
				Expect(resources.Data).To(HaveLen(0))
			})
		})

		Context("When accounts are present", func() {
			uuids := make([]string, 3)
			BeforeEach(func() {
				for i := 0; i < 3; i++ {
					uuids[i] = uuid.NewV4().String()
					_, _ = client.Create(getTestAccountParams(uuids[i]))
				}
			})

			AfterEach(func() {
				for i := 0; i < 3; i++ {
					_ = client.Delete(uuids[i], 0)
				}
			})

			It("should return a list of account data", func() {
				resources, err := client.List(PagingParams{})
				Expect(err).To(BeNil())
				Expect(resources.Data).To(HaveLen(3))
				Expect(resources.Data[0].ID).To(Equal(uuids[0]))
				Expect(resources.Data[1].ID).To(Equal(uuids[1]))
				Expect(resources.Data[2].ID).To(Equal(uuids[2]))
				Expect(resources.Links).To(Equal(map[string]string{
					"self":  "/v1/organisation/accounts",
					"first": "/v1/organisation/accounts?page%5Bnumber%5D=first",
					"last":  "/v1/organisation/accounts?page%5Bnumber%5D=last",
				}))
			})

			It("should support paging", func() {
				resources, err := client.List(PagingParams{ number: "0" , size: "2"})
				Expect(err).To(BeNil())
				Expect(resources.Data).To(HaveLen(2))
				Expect(resources.Data[0].ID).To(Equal(uuids[0]))
				Expect(resources.Data[1].ID).To(Equal(uuids[1]))

				resources, err = client.List(PagingParams{ number: "1" , size: "2"})
				Expect(err).To(BeNil())
				Expect(resources.Data).To(HaveLen(1))
				Expect(resources.Data[0].ID).To(Equal(uuids[2]))
			})
		})
	})

	Describe("Fetch", func() {
		Context("With an invalid account ID", func() {
			It("should return a validation error", func() {
				resource, err := client.Fetch("not-an-uuid")

				Expect(err).To(Equal(fmt.Errorf("account ID must be a valid v4 UUID")))
				Expect(resource).To(BeNil())
			})
		})

		Context("A non-existent account", func() {
			It("should return an error", func() {
				anyUUID := uuid.NewV4().String()
				resource, err := client.Fetch(anyUUID)

				Expect(err).To(Equal(fmt.Errorf("request failed []: record %s does not exist", anyUUID)))
				Expect(resource).To(BeNil())
			})
		})

		Context("An existing account", func() {
			BeforeEach(func() {
				_, _ = client.Create(getTestAccountParams(testAccountId))
			})

			It("should return account data", func() {
				resource, err := client.Fetch(testAccountId)

				Expect(err).To(BeNil())
				Expect(resource.Data.ID).To(Equal(testAccountId))
			})
		})
	})

	Describe("Delete", func() {
		Context("A non-existent account", func() {
			It("should not fail", func() {
				err := client.Delete(uuid.NewV4().String(), 0)
				Expect(err).To(BeNil())
			})
		})

		Context("An existing account", func() {

			BeforeEach(func() {
				_, _ = client.Create(getTestAccountParams(testAccountId))
			})

			It("should not fail", func() {
				err := client.Delete(testAccountId, 0)
				Expect(err).To(BeNil())

				resource, err := client.Fetch(testAccountId)
				Expect(resource).To(BeNil())
			})
		})
	})
})

func getTestAccountParams(id string) AccountParams {
	return AccountParams{
		Attributes: Attributes{
			Country: "GB",
		},
		ID:             id,
		OrganizationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
	}
}
