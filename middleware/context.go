package middleware

import (
	"GoOnlineJudge/model"

	"github.com/zenazn/goji/web"
)

func UserToC(c *web.C, user *model.User) {
	if c.Env == nil {
		c.Env = make(map[interface{}]interface{})
	}

	c.Env["user"] = user
}

// ToUser returns the User from the current
// request context. If the User does not exist
// a nil value is returned.
func ToUser(c web.C) *model.User {
	var v = c.Env["user"]
	if v == nil {
		return nil
	}
	u, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return u
}

func ProblemToC(c *web.C, problem *model.Problem) {
	if c.Env == nil {
		c.Env = make(map[interface{}]interface{})
	}

	c.Env["problem"] = problem
}

// ToProblem returns the problem from the current
// request context. If the problem does not exist
// a nil value is returned.
func ToProblem(c web.C) *model.Problem {
	var v = c.Env["problem"]
	if v == nil {
		return nil
	}
	u, ok := v.(*model.Problem)
	if !ok {
		return nil
	}
	return u
}

func ContestToC(c *web.C, contest *model.Contest) {
	if c.Env == nil {
		c.Env = make(map[interface{}]interface{})
	}

	c.Env["contest"] = contest
}

// ToContest returns the contest from the current
// request context. If the contest does not exist
// a nil value is returned.
func ToContest(c web.C) *model.Contest {
	var v = c.Env["contest"]
	if v == nil {
		return nil
	}
	u, ok := v.(*model.Contest)
	if !ok {
		return nil
	}
	return u
}

func NewsToC(c *web.C, news *model.News) {
	if c.Env == nil {
		c.Env = make(map[interface{}]interface{})
	}

	c.Env["news"] = news
}

// ToNews returns the news from the current
// request context. If the news does not exist
// a nil value is returned.
func ToNews(c web.C) *model.News {
	var v = c.Env["news"]
	if v == nil {
		return nil
	}
	u, ok := v.(*model.News)
	if !ok {
		return nil
	}
	return u
}
