package ircity

// Province represents a province with a code and name.
type Province struct {
	Code uint   `db:"code" bson:"code" json:"code"`
	Name string `db:"name" bson:"name" json:"name"`
}

// Provinces returns a slice of all provinces.
func Provinces() []Province {
	res := make([]Province, 0, len(provinces))
	for _, prov := range provinces {
		res = append(res, prov)
	}
	return res
}

// FindProvince searches for a province by its code.
// Returns a pointer to the Province if found, otherwise nil.
func FindProvince(code uint) *Province {
	if prov, ok := provinces[code]; ok {
		return &prov
	}
	return nil
}

// provinces is a map of province codes to Province structs.
var provinces = map[uint]Province{
	3:  {Code: 3, Name: "آذربایجان شرقی"},
	4:  {Code: 4, Name: "آذربایجان غربی"},
	24: {Code: 24, Name: "اردبیل"},
	10: {Code: 10, Name: "اصفهان"},
	30: {Code: 30, Name: "البرز"},
	16: {Code: 16, Name: "ایلام"},
	18: {Code: 18, Name: "بوشهر"},
	23: {Code: 23, Name: "تهران"},
	14: {Code: 14, Name: "چهارمحال وبختیاری"},
	29: {Code: 29, Name: "خراسان جنوبی"},
	9:  {Code: 9, Name: "خراسان رضوی"},
	28: {Code: 28, Name: "خراسان شمالی"},
	6:  {Code: 6, Name: "خوزستان"},
	19: {Code: 19, Name: "زنجان"},
	20: {Code: 20, Name: "سمنان"},
	11: {Code: 11, Name: "سیستان وبلوچستان"},
	7:  {Code: 7, Name: "فارس"},
	26: {Code: 26, Name: "قزوین"},
	25: {Code: 25, Name: "قم"},
	12: {Code: 12, Name: "کردستان"},
	8:  {Code: 8, Name: "کرمان"},
	5:  {Code: 5, Name: "کرمانشاه"},
	17: {Code: 17, Name: "کهگیلویه وبویراحمد"},
	27: {Code: 27, Name: "گلستان"},
	1:  {Code: 1, Name: "گیلان"},
	15: {Code: 15, Name: "لرستان"},
	2:  {Code: 2, Name: "مازندران"},
	0:  {Code: 0, Name: "مرکزی"},
	22: {Code: 22, Name: "هرمزگان"},
	13: {Code: 13, Name: "همدان"},
	21: {Code: 21, Name: "یزد"},
}
