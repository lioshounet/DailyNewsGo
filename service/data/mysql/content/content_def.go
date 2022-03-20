package content

const (
	Offset = 20

	DeletedOff = 0
	DeletedOn  = 1

	// 内容类型
	TypeDaily = 1 // 日报

	selectField = "id,uid,content_date,content_tag,content_text,content_type,deleted,c_time,m_time"
	insertField = "uid,content_date,content_tag,content_text,content_type,deleted,c_time,m_time"
)

type Content struct {
	Id          uint64 `json:"id"`
	Uid         uint64 `json:"uid"`
	ContentDate uint64 `json:"content_date"`
	ContentTag  uint64 `json:"content_tag"`
	ContentText string `json:"content_text"`
	ContentType uint64 `json:"content_type"`
	Deleted     uint64 `json:"deleted"`
	CTime       uint64 `json:"c_time"`
	MTime       uint64 `json:"m_time"`
}
