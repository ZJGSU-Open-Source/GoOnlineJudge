package schedule

import (
    "GoOnlineJudge/model"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
)

type PKUJudger struct {
    client    *http.Client
    testRx    *regexp.Regexp
    titleRx   *regexp.Regexp
    resLimtRx *regexp.Regexp
    ctxRx     *regexp.Regexp
    srcRx     *regexp.Regexp
    hintRx    *regexp.Regexp
}

var PKUlogger *log.Logger

func (p *PKUJudger) Host() string {
    return "PKU"
}
func (p *PKUJudger) Ping() error {
    p.client = &http.Client{Timeout: time.Second * 10}
    resp, err := p.client.Get("http://poj.org/problem?id=1000")
    if err != nil {
        return ErrConnectFailed
    }
    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        return nil
    }
    return ErrResponse
}

func (h *PKUJudger) Init() {
    h.client = &http.Client{Timeout: time.Second * 10}

    titlePat := `<div class="ptt" lang=".*?">(.*?)</div>`

    h.titleRx = regexp.MustCompile(titlePat)

    resLimtPat := `<td><b>Time Limit:</b> (\d+)MS</td><td width="10px"></td><td><b>Memory Limit:</b> (\d+)K</td>`
    h.resLimtRx = regexp.MustCompile(resLimtPat)

    ctxPat := `(?s)<div class="ptx" lang=".*?">(.*?)</div><p class="pst">`
    h.ctxRx = regexp.MustCompile(ctxPat)

    testPat := `(?s)<pre class="sio">(.*?)</pre>`
    h.testRx = regexp.MustCompile(testPat)

    srcPat := `<a href="searchproblem?field=source&key=.*?">(.*?)</a>`
    h.srcRx = regexp.MustCompile(srcPat)

    hintPat := `(?s)<p class="pst">Hint</p><div class="ptx" lang=".*?">(.*?)</div>`
    h.hintRx = regexp.MustCompile(hintPat)

    PKULogfile, err := os.Create("log/pku.log")
    if err != nil {
        log.Println(err)
        return
    }
    PKUlogger = log.New(PKULogfile, "[PKU]", log.Ldate|log.Ltime)
}

func (h *PKUJudger) GetProblemPage(pid string) (string, error) {
    resp, err := h.client.Get("http://poj.org/problem?id=" + pid)
    if err != nil {
        return "", ErrConnectFailed
    }
    b, _ := ioutil.ReadAll(resp.Body)
    html := string(b)
    return html, nil

}
func (h *PKUJudger) IsExist(page string) bool {
    return strings.Index(page, "Can not find problem") < 0
}
func (h *PKUJudger) ReplaceImg(text string) string {

    if strings.Index(text, `<img src="`) >= 0 {
        text = strings.Replace(text, `<img src="`, `<img src="http://poj.org/`, -1)
    } else {
        text = strings.Replace(text, `<img src=`, `<img src=http://poj.org/`, -1)
    }
    text = strings.Replace(text, `<IMG src="`, `<img src="http://poj.org/`, -1)

    return text
}

func (h *PKUJudger) SetDetail(pid string, html string) error {
    log.Println(pid)
    pro := model.Problem{}
    pro.RPid, _ = strconv.Atoi(pid)
    pro.ROJ = "PKU"
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
        log.Println("ctx match error, PKU pid is", pid)
        return ErrMatchFailed
    }
    pro.Description = template.HTML(h.ReplaceImg(cxtMatch[0][1]))
    pro.Input = template.HTML(h.ReplaceImg(cxtMatch[1][1]))
    pro.Output = template.HTML(h.ReplaceImg(cxtMatch[2][1]))

    test := h.testRx.FindAllStringSubmatch(html, 2)
    if len(test) < 2 {
        log.Println("test data error, PKU pid is", pid)
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
        pro.Hint = template.HTML(hint[1])
    }

    proModel := &model.ProblemModel{}
    proModel.Insert(pro)
    return nil
}

func (h *PKUJudger) GetProblems() error {
    vidsModel := &model.VIdsModel{}
    StartId, err := vidsModel.GetLastID("PKU")
    if err == model.DBErr {
        return err
    } else if StartId < 1000 {
        StartId = 999
    }
    errCnt := 0
    lastId := StartId
    for i := 1; ; i++ {
        pid := strconv.Itoa(StartId + i)
        page, err := h.GetProblemPage(pid)
        if err != nil { //offline
            PKUlogger.Println("pid["+pid+"]: ", err, ".")
            return err
        }
        if h.IsExist(page) {
            err := h.SetDetail(pid, page)
            if err != nil {
                PKUlogger.Println("pid["+pid+"]: ", "import error.")
            } else {
                lastId = StartId + i
                vidsModel.SetLastID("PKU", lastId)
            }
            errCnt = 0
        } else {
            PKUlogger.Println("pid["+pid+"]: ", "not exist.")
            errCnt++
        }

        if errCnt >= 100 { //If "not exist" continuously repeat 100 times, terminate it.
            break
        }
    }
    PKUlogger.Println("import terminated. Last pid is ", lastId, ".")
    return nil
}
