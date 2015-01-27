package vjudger

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	// "os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type HDUJudger struct {
	client   *http.Client
	token    string
	pat      *regexp.Regexp
	username string
	userpass string
}

var HDURes = map[string]int{"Queuing": 0,
	"Compiling": 1, "Running": 1,
	"Compilation Error":                         2,
	"Accepted":                                  3,
	"Runtime Error<br>(STACK_OVERFLOW)":         4,
	"Runtime Error<br>(ACCESS_VIOLATION)":       4,
	"Runtime Error<br>(ARRAY_BOUNDS_EXCEEDED)":  4,
	"Runtime Error<br>(FLOAT_DENORMAL_OPERAND)": 4,
	"Runtime Error<br>(FLOAT_DIVIDE_BY_ZERO)":   4,
	"Runtime Error<br>(FLOAT_OVERFLOW)":         4,
	"Runtime Error<br>(FLOAT_UNDERFLOW )":       4,
	"Runtime Error<br>(INTEGER_OVERFLOW)":       4,
	"Runtime Error<br>(INTEGER_DIVIDE_BY_ZERO)": 4,
	"Wrong Answer":                              5,
	"Time Limit Exceeded":                       6,
	"Memory Limit Exceeded":                     7,
	"Output Limit Exceeded":                     8,
	"Presentation Error":                        9,
	"System Error":                              10}

func (h *HDUJudger) Init(_ *User) error {
	jar, _ := cookiejar.New(nil)
	h.client = &http.Client{Jar: jar}
	h.token = "HDU"
	pattern := `(\d+)</td><td>(.*?)</td><td>(?s:.*?)<font color=.*?>(.*?)</font>.*?<td>(\d+)MS</td><td>(\d+)K</td><td><a href="/viewcode.php\?rid=\d+"  target=_blank>(\d+) B</td><td>(.*?)</td>`
	h.pat = regexp.MustCompile(pattern)
	h.username = "mysake"
	h.userpass = "123456"
	return nil
}

func (h *HDUJudger) Match(token string) bool {
	if token == h.token {
		return true
	}
	return false
}
func (h *HDUJudger) Login(_ *User) error {

	h.client.Get("http://acm.hdu.edu.cn")

	uv := url.Values{}
	uv.Add("username", h.username)
	uv.Add("userpass", h.userpass)
	uv.Add("login", "Sign In")

	req, err := http.NewRequest("POST", "http://acm.hdu.edu.cn/userloginex.php?action=login", strings.NewReader(uv.Encode()))
	if err != nil {
		return BadInternet
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println("err", err)
		return BadInternet
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	if strings.Index(html, "No such user or wrong password.") >= 0 {
		return LoginFailed
	}

	return nil
}

func (h *HDUJudger) Submit(u *User) error {

	uv := url.Values{}
	uv.Add("check", "0")
	uv.Add("problemid", strconv.Itoa(u.Vid))
	uv.Add("language", strconv.Itoa(u.Lang))
	uv.Add("usercode", u.Code)

	req, err := http.NewRequest("POST", "http://acm.hdu.edu.cn/submit.php?action=submit", strings.NewReader(uv.Encode()))
	if err != nil {
		log.Println(err)
		return BadInternet
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	u.SubmitTime = time.Now()
	resp, err := h.client.Do(req)
	if err != nil {
		log.Println(err)
		return BadInternet
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	if strings.Index(html, "No such problem") >= 0 {
		return NoSuchProblem
	}
	if strings.Index(html, "One or more following ERROR(s) occurred.") >= 0 {
		return SubmitFailed
	}
	//TODO check whether submit success
	return nil
}

func (h *HDUJudger) GetStatus(u *User) error {

	statusUrl := "http://acm.hdu.edu.cn/status.php?first=&" +
		"pid=" + strconv.Itoa(u.Vid) +
		"&user=" + h.username +
		"&lang=" + strconv.Itoa(u.Lang) + "&status=0"

	endTime := time.Now().Add(MAX_WaitTime * time.Second)

	for true {
		if time.Now().After(endTime) {
			return BadStatus
		}
		resp, err := h.client.Get(statusUrl)
		if err != nil {
			log.Println(err)
			return BadInternet
		}
		defer resp.Body.Close()

		b, _ := ioutil.ReadAll(resp.Body)
		AllStatus := h.pat.FindAllStringSubmatch(string(b), -1)

		layout := "2006-01-02 15:04:05 (MST)" //parse time
		for i := len(AllStatus) - 1; i >= 0; i-- {
			status := AllStatus[i]
			t, _ := time.Parse(layout, status[2]+" (CST)")
			t = t.Add(40 * time.Second) //HDU server's time is less 36s.
			log.Println(t, u.SubmitTime)
			log.Println(status[1:])
			if t.After(u.SubmitTime) {
				rid := status[1]
				u.Result = HDURes[status[3]]

				if u.Result >= 2 {
					if u.Result == 2 {
						u.CE, err = h.GetCEInfo(rid)
						if err != nil {
							log.Println(err)
						}
					}

					u.Time, _ = strconv.Atoi(status[4])
					u.Mem, _ = strconv.Atoi(status[5])
					u.Length, _ = strconv.Atoi(status[6])
					return nil
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("here")
	return nil
}

func (h *HDUJudger) GetCEInfo(rid string) (string, error) {
	resp, err := h.client.Get("http://acm.hdu.edu.cn/viewerror.php?rid=" + rid)
	if err != nil {
		log.Println(err)
		return "", BadInternet
	}

	b, _ := ioutil.ReadAll(resp.Body)
	pre := "(?s)<pre>(.*?)</pre>"
	re := regexp.MustCompile(pre)
	match := re.FindStringSubmatch(string(b))
	return match[1], nil
}

func (h *HDUJudger) Run(u *User) error {
	for _, apply := range []func(*User) error{h.Init, h.Login, h.Submit, h.GetStatus} {
		t := time.Now()
		if err := apply(u); err != nil {
			log.Println(err)
			return err
		}
		log.Println("time:", time.Now().Sub(t))
	}
	return nil
}
