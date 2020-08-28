// Promotion is the model
type Go struct {
	ID int `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Priority int `json:"priority"`
	Exclusive bool `json:"exclusive"`
	Used int `json:"used"`
	CouponBased bool `json:"couponBased"`
	Rules []interface{} `json:"rules"`
	Actions []interface{} `json:"actions"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Channels []interface{} `json:"channels"`
	Links GoLinks `json:"_links"`
}

type GoSelf struct {
	Href string `json:"href"`
}

type GoLinks struct {
	Self GoSelf `json:"self"`
}