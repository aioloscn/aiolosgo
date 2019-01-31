package parser

import (
	"aiolosgo/crawler/engine"
	"aiolosgo/crawler/model"
	"bytes"
	"regexp"
)

var re = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^>]+)</div>`)
var re1 = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^>]+)</div>`)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	//age, err := strconv.Atoi(extractString(contents, ageRe))
	//if err == nil {
	//	profile.Age = age
	//}
	//
	//height, err := strconv.Atoi(extractString(contents, heightRe))
	//if err == nil {
	//	profile.Height = height
	//}
	//
	//weight, err := strconv.Atoi(extractString(contents, weightRe))
	//if err == nil {
	//	profile.Weight = weight
	//}
	//
	//profile.Income = extractString(contents, incomeRe)
	//profile.Gender = extractString(contents, genderRe)
	//profile.Car = extractString(contents, carRe)
	//profile.Education = extractString(contents, educationRe)
	//profile.Hokou = extractString(contents, hokouRe)
	//profile.House = extractString(contents, houseRe)
	//profile.Marriage = extractString(contents, marriageRe)
	//profile.Occupation = extractString(contents, occupationRe)
	//profile.Xinzuo = extractString(contents, xinzuoRe)

	var buffer bytes.Buffer
	submatch := re.FindAllSubmatch(contents, -1)
	for _, v := range submatch {
		buffer.WriteString(string(v[1]))
		buffer.WriteString(" ")
	}

	submatch = re1.FindAllSubmatch(contents, -1)
	for _, v := range submatch {
		buffer.WriteString(string(v[1]))
		buffer.WriteString(" ")
	}

	profile.Details = buffer.String()

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}