package lecturer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gavv/httpexpect"
	"github.com/go-pg/pg"
	pG "github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("移除讲师", func() {
	var lecturerId int64
	BeforeEach(func() {
		_, err := pG.DB.QueryOne(
			pg.Scan(&lecturerId),
			"INSERT INTO lecturers (id, name, introduction, topic) VALUES (DEFAULT, ?, ?, ?) RETURNING id",
			"testName", "testIntroduction", "testTopic")
		Expect(err).NotTo(HaveOccurred())
	})
	Describe("根据参数移除讲师", func() {
		Context("传入有效的lecturerId", func() {
			It("返回被移除敏捷之旅的分享讲师的数据", func() {
				httpExpect := httpexpect.New(GinkgoT(), server.URL)
				httpExpect.DELETE(fmt.Sprintf(`/lecturers/%s`, strconv.FormatInt(lecturerId, 10))).
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().
					ContainsKey("code").ValueEqual("code", 0).
					ContainsKey("msg").ValueEqual("msg", "ok").
					ContainsKey("data").Value("data").Object()
			})
		})
	})
	AfterEach(func() {
		_, err := pG.DB.Exec("DELETE FROM lecturers WHERE true")
		Expect(err).NotTo(HaveOccurred())
	})
})
