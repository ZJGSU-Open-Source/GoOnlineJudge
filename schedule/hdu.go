package schedule

import (
	"GoOnlineJudge/model"
	iconv "github.com/djimenez/iconv-go"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type HDUJudger struct {
	client    *http.Client
	testRx    *regexp.Regexp
	titleRx   *regexp.Regexp
	resLimtRx *regexp.Regexp
	ctxRx     *regexp.Regexp
	srcRx     *regexp.Regexp
	hintRx    *regexp.Regexp
}

var hdulogger *log.Logger

func (h *HDUJudger) Init() {
	h.client = &http.Client{Timeout: time.Second * 10}

	titlePat := `<h1 style='color:#1A5CC8'>(.*?)</h1>`
	h.titleRx = regexp.MustCompile(titlePat)

	resLimtPat := `Time Limit: \d+/(\d+) MS \(Java/Others\)&nbsp;&nbsp;&nbsp;&nbsp;Memory Limit: \d+/(\d+) K \(Java/Others\)<br>`
	h.resLimtRx = regexp.MustCompile(resLimtPat)

	ctxPat := `(?s)<div class=panel_content>(.*?)</div><div class=panel_bottom>`
	h.ctxRx = regexp.MustCompile(ctxPat)

	testPat := `(?s)<div style="font-family:Courier New,Courier,monospace;">(.*?)</??div`
	h.testRx = regexp.MustCompile(testPat)

	srcPat := `<a href="/search.php\?field=problem&key=.*?&source=1&searchmode=source"> (.*?) </a>`
	h.srcRx = regexp.MustCompile(srcPat)

	hintPat := `(?s)<i>Hint</i></div>(.*?)</div>`
	h.hintRx = regexp.MustCompile(hintPat)
}

func (h *HDUJudger) GetProblemPage(pid string) (string, error) {
	resp, err := h.client.Get("http://acm.hdu.edu.cn/showproblem.php?pid=" + pid)
	if err != nil {
		return "", ErrConnectFailed
	}
	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	return html, nil

}
func (h *HDUJudger) IsExist(page string) bool {
	return strings.Index(page, "No such problem") < 0
}
func (h *HDUJudger) ReplaceImg(text string) string {
	text = strings.Replace(text, `<img src=/data/images/`, `<img src=http://acm.hdu.edu.cn/data/images/`, -1)
	text = strings.Replace(text, `<img src=data/images/`, `<img src=http://acm.hdu.edu.cn/data/images/`, -1)
	text = strings.Replace(text, `<img src=../../../data/images/`, `<img src=http://acm.hdu.edu.cn/data/images/`, -1)
	text = strings.Replace(text, `<img src=../../data/images/`, `<img src=http://acm.hdu.edu.cn/data/images/`, -1)
	return text
}

func (h *HDUJudger) SetDetail(pid string, html string) error {
	log.Println(pid)
	pro := model.Problem{}
	pro.RPid, _ = strconv.Atoi(pid)
	pro.ROJ = "HDU"
	pro.Status = StatusAvailable

	titleMatch := h.titleRx.FindStringSubmatch(html)
	if len(titleMatch) < 1 {
		log.Println(titleMatch)
		return ErrMatchFailed
	}
	pro.Title = titleMatch[1]

	if strings.Index(html, "Special Judge") >= 0 {
		pro.Special = 1
	}

	resMatch := h.resLimtRx.FindStringSubmatch(html)
	if len(resMatch) < 3 {
		log.Println(resMatch)
		return ErrMatchFailed
	}
	pro.Time, _ = strconv.Atoi(resMatch[1])
	pro.Time /= 1000 //ms -> s
	pro.Memory, _ = strconv.Atoi(resMatch[2])

	cxtMatch := h.ctxRx.FindAllStringSubmatch(html, 3)
	if len(cxtMatch) < 3 {
		log.Println("ctx match error, hdu pid is", pid)
		return ErrMatchFailed
	}
	pro.Description = template.HTML(h.ReplaceImg(cxtMatch[0][1]))
	pro.Input = template.HTML(h.ReplaceImg(cxtMatch[1][1]))
	pro.Output = template.HTML(h.ReplaceImg(cxtMatch[2][1]))

	test := h.testRx.FindAllStringSubmatch(html, 2)
	if len(test) < 2 {
		log.Println("test data error, hdu pid is", pid)
		return ErrMatchFailed
	}
	pro.In = test[0][1]
	pro.Out = test[1][1]

	src := h.srcRx.FindStringSubmatch(html)
	if len(src) > 1 {
		pro.Source = src[1]
	}

	hint := h.hintRx.FindStringSubmatch(html)
	if len(hint) > 1 {
		pro.Hint = hint[1]
	}

	proModel := &model.ProblemModel{}
	proModel.Insert(pro)
	return nil
}

func (h *HDUJudger) GetProblems() {
	vidsModel := &model.VIdsModel{}
	StartId, _ := vidsModel.GetLastID("HDU")
	errCnt := 0
	lastId := StartId
	for i := 1; ; i++ {
		pid := strconv.Itoa(StartId + i)
		page, err := h.GetProblemPage(pid)
		if err != nil {
			hdulogger.Println("pid["+pid+"]: ", err, ".")
			return
		}
		cpage, err := iconv.ConvertString(page, "gb2312", "utf-8")
		if err != nil { //Although getting error, continue proccess it.
			hdulogger.Println("pid["+pid+"]: ", "encode convert error.")
			cpage = page
		}

		if h.IsExist(cpage) {
			err := h.SetDetail(pid, cpage)
			if err != nil {
				hdulogger.Println("pid["+pid+"]: ", "import error.")
			} else {
				lastId = StartId + i
			}
			errCnt = 0
		} else {
			hdulogger.Println("pid["+pid+"]: ", "not exist.")
			errCnt++
		}

		if errCnt >= 100 { //If "not exist" repeat 100 times, terminate it.
			break
		}
	}
	vidsModel.SetLastID("HDU", lastId)
	hdulogger.Println("import terminated. Last pid is ", lastId, ".")
}
