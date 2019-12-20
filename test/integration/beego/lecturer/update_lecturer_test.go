package lecturer

//import (
//	"fmt"
//	"net/http"
//	"strconv"
//
//	"github.com/gavv/httpexpect"
//	"github.com/go-pg/pg"
//	pG "github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//)
//
//var _ = Describe("更新讲师", func() {
//	var lecturerId int64
//	BeforeEach(func() {
//		_, err := pG.DB.QueryOne(
//			pg.Scan(&lecturerId),
//			"INSERT INTO lecturers (id, name, introduction, topic) VALUES (DEFAULT, ?, ?, ?) RETURNING id",
//			"testName", "testIntroduction", "testTopic")
//		Expect(err).NotTo(HaveOccurred())
//	})
//	Describe("提交数据更新讲师", func() {
//		Context("提交正确的敏捷之旅的分享讲师数据", func() {
//			It("返回更新后的敏捷之旅的分享讲师数据", func() {
//				httpExpect := httpexpect.New(GinkgoT(), server.URL)
//				body := map[string]interface{}{
//					"introduction": "update-introduction",
//					"topic":        "update-topic",
//				}
//				httpExpect.PUT(fmt.Sprintf(`/lecturers/%s`, strconv.FormatInt(lecturerId, 10))).
//					WithJSON(body).
//					Expect().
//					Status(http.StatusOK).
//					JSON().
//					Object().
//					ContainsKey("code").ValueEqual("code", 0).
//					ContainsKey("msg").ValueEqual("msg", "ok").
//					ContainsKey("data").Value("data").Object().
//					ContainsKey("id").ValueEqual("id", lecturerId)
//			})
//		})
//	})
//	AfterEach(func() {
//		_, err := pG.DB.Exec("DELETE FROM lecturers WHERE true")
//		Expect(err).NotTo(HaveOccurred())
//	})
//})
