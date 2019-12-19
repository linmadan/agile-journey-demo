package lecturer

import (
	"net/http"

	"github.com/gavv/httpexpect"
	"github.com/go-pg/pg"
	pG "github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("返回讲师列表", func() {
	var lecturerId int64
	BeforeEach(func() {
		_, err := pG.DB.QueryOne(
			pg.Scan(&lecturerId),
			"INSERT INTO lecturers (id, name, introduction, topic) VALUES (DEFAULT, ?, ?, ?) RETURNING id",
			"testName", "testIntroduction", "testTopic")
		Expect(err).NotTo(HaveOccurred())
	})
	Describe("根据参数返回敏捷之旅的分享讲师列表", func() {
		Context("传入有效的参数", func() {
			It("返回敏捷之旅的分享讲师数据列表", func() {
				httpExpect := httpexpect.New(GinkgoT(), server.URL)
				httpExpect.GET("/lecturers/").
					WithQuery("offset", 0).
					WithQuery("limit", 0).
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().
					ContainsKey("code").ValueEqual("code", 0).
					ContainsKey("msg").ValueEqual("msg", "ok").
					ContainsKey("data").Value("data").Object().
					ContainsKey("count").ValueEqual("count", 1).
					ContainsKey("lecturers").Value("lecturers").Array()
			})
		})
	})
	AfterEach(func() {
		_, err := pG.DB.Exec("DELETE FROM lecturers WHERE true")
		Expect(err).NotTo(HaveOccurred())
	})
})
