package vjudger

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

type HDUJudger struct {
	client *http.Client
	token  string
}

func (h *HDUJudger) Init(_ *User) error {
	jar, _ := cookiejar.New(nil)
	h.client = &http.Client{Jar: jar}
	h.token = "HDU"
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
	return nil
}

func (h *HDUJudger) Run(u *User) error {
	for _, apply := range []func(*User) error{h.Init, h.Login, h.Submit, h.GetStatus} {
		if err := apply(u); err != nil {
			return err
		}
	}
	return nil
}
