package user

const (
	DeletedOff = 0
	DeletedOn  = 1

	RD    = 1
	FE    = 2
	UI    = 3
	UE    = 4
	OTHER = 0

	selectField = "id,uid,user_name,identity,email,phone,avatar,sn,given_name,deleted,c_time,m_time"
	insertField = "uid, user_name,identity, email, phone, avatar, sn, given_name, c_time, m_time"
)

var IdentityMap = map[uint64]string{
	RD:    "RD",
	FE:    "FE",
	UI:    "UI",
	UE:    "UE",
	OTHER: "Other",
}

type User struct {
	Id        uint64 `json:"id"`
	Uid       uint64 `json:"uid"`
	UserName  string `json:"user_name"`
	Identity  uint64 `json:"identity"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Sn        string `json:"sn"`
	GivenName string `json:"given_name"`
	Deleted   uint64 `json:"deleted"`
	CTime     uint64 `json:"c_time"`
	MTime     uint64 `json:"m_time"`
}
