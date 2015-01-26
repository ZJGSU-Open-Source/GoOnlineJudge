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
	client *http.Client
	token  string
	pat    *regexp.Regexp
}

var HDURes = map[string]int{"Accepted": 1}

func (h *HDUJudger) Init(_ *User) error {
	jar, _ := cookiejar.New(nil)
	h.client = &http.Client{Jar: jar}
	h.token = "HDU"
	pattern := `\d+</td><td>(.*?)</td><td><font color=.*?>(.*?)</font>.*?<td>(\d+)MS</td><td>(\d+)K</td><td><a href="/viewcode.php\?rid=\d+"  target=_blank>(\d+) B</td><td>(.*?)</td>`
	h.pat = regexp.MustCompile(pattern)
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

	username := "mysake"
	userpass := "123456"
	uv := url.Values{}
	uv.Add("username", username)
	uv.Add("userpass", userpass)
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

	//TODO check login
	return nil
	// b, _ := ioutil.ReadAll(resp.Body)
	// log.Println(resp.Header)
	// log.Println(resp.Status)

	// f, err := os.Create("/Users/sake/hdu.html")
	// if err != nil {
	// 	return
	// }
	// defer f.Close()
	// f.Write(b)
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

	//TODO check whether submit success
	return nil
}

func (h *HDUJudger) GetStatus(u *User) error {
	time.Sleep(1 * time.Second)
	username := "mysake"
	statusUrl := "http://acm.hdu.edu.cn/status.php?first=&" +
		"pid=" + strconv.Itoa(u.Vid) +
		"&user=" + username +
		"&lang=" + strconv.Itoa(u.Lang) + "&status=0"
	resp, err := h.client.Get(statusUrl)
	if err != nil {
		log.Println(err)
		return BadInternet
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	s := string(b)
	find := h.pat.FindAllStringSubmatch(s, -1)

	layout := "2006-01-02 15:04:05 (MST)" //parse time

	for i := len(find) - 1; i >= 0; i-- {
		str := find[i]
		t, _ := time.Parse(layout, str[1]+" (CST)")
		t = t.Add(36 * time.Second) //HDU server's time is less 36s.
		if t.After(u.SubmitTime) {
			log.Println(str[1:])
			u.Result = HDURes[str[2]]
			u.Time, _ = strconv.Atoi(str[3])
			u.Mem, _ = strconv.Atoi(str[4])
			u.Length, _ = strconv.Atoi(str[5])
			break
		}
	}
	return nil
}

func (h *HDUJudger) Run(u *User) error {
	for _, apply := range []func(*User) error{h.Init, h.Login, h.Submit, h.GetStatus} {
		t := time.Now()
		if err := apply(u); err != nil {
			return err
		}
		log.Println("time:", time.Now().Sub(t))
	}
	return nil
}
