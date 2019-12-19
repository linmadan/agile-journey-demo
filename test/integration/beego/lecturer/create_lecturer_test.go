package lecturer

import (
	"net/http"

	"github.com/gavv/httpexpect"
	pG "github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("创建新讲师", func() {
	Describe("提交数据创建新讲师", func() {
		Context("提交正确的新敏捷之旅的分享讲师数据", func() {
			It("返回敏捷之旅的分享讲师数据", func() {
				httpExpect := httpexpect.New(GinkgoT(), server.URL)
				body := map[string]interface{}{
					"name":         "test-name",
					"introduction": "test-introduction",
					"topic":        "test-topic",
				}
				httpExpect.POST("/lecturers/").
					WithJSON(body).
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().
					ContainsKey("code").ValueEqual("code", 0).
					ContainsKey("msg").ValueEqual("msg", "ok").
					ContainsKey("data").Value("data").Object().
					ContainsKey("id").ValueNotEqual("id", BeZero())
			})
		})
	})
	AfterEach(func() {
		_, err := pG.DB.Exec("DELETE FROM lecturers WHERE true")
		Expect(err).NotTo(HaveOccurred())
	})
})
