package content

import (
	"database/sql"
	"errors"
	"fmt"
	"thor/service/data"
	"time"
)

type ContentModel struct {
	Db *sql.DB
	// Logger *log.Logger
}

func NewContentModel() *ContentModel {
	return &ContentModel{
		Db: data.GetBladeDb(),
	}
}

func (t *ContentModel) getTableName() string {
	return "content"
}

// 根据uid查询记录
func (t *ContentModel) GetContentListByUid(uid, contentType, limit, offset uint64) ([]*Content, error) {
	if uid < 0 || contentType < 1 {
		return nil, errors.New("params fail!")
	}

	if offset == 0 {
		offset = Offset
	}

	queryFmt := "select %s from %s where uid = %d and content_type = %d and deleted = %d order by m_time desc limit %d,%d"
	querySql := fmt.Sprintf(queryFmt, selectField, t.getTableName(), uid, contentType, DeletedOff, limit, offset)

	result, err := t.contentQuery(querySql)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *ContentModel) GetListById(id uint64) ([]*Content, error) {
	if id < 1 {
		return nil, errors.New("params fail!")
	}

	queryFmt := "select %s from %s where id=%d"
	querySql := fmt.Sprintf(queryFmt, selectField, t.getTableName(), id)

	result, err := t.contentQuery(querySql)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *ContentModel) GetAllContentList(contentType, limit, offset uint64) ([]*Content, error) {
	if contentType < 1 {
		return nil, errors.New("params fail!")
	}

	if offset == 0 {
		offset = Offset
	}

	queryFmt := "select %s from %s where content_type = %d and deleted = %d order by m_time desc limit %d,%d"
	querySql := fmt.Sprintf(queryFmt, selectField, t.getTableName(), contentType, DeletedOff, limit, offset)

	result, err := t.contentQuery(querySql)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *ContentModel) UpdateContentTextById(id, uid uint64, contentText string) (uint64, error) {
	if id < 1 || uid < 1 || contentText == "" {
		return 0, errors.New("params fail!")
	}

	nowTime := time.Now().Unix()
	updateFmt := "update %s set content_text=?, m_time=%d where id=%d and uid=%d limit 1"
	execSql := fmt.Sprintf(updateFmt, t.getTableName(), nowTime, id, uid)

	result, err := t.Db.Exec(execSql, contentText)
	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(row), nil
}

// 根据时间查询记录
func (t *ContentModel) GetContentListByDate(contentDate, contentType, limit, offset uint64) ([]*Content, error) {
	if contentDate < 0 || contentType < 1 {
		return nil, errors.New("params fail!")
	}

	if offset == 0 {
		offset = Offset
	}

	queryFmt := "select %s from %s where content_date = %d and content_type = %d and deleted = %d order by m_time desc limit %d, %d"
	querySql := fmt.Sprintf(queryFmt, selectField, t.getTableName(), contentDate, contentType, DeletedOff, limit, offset)

	result, err := t.contentQuery(querySql)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 插入一条记录
func (t *ContentModel) InsertContent(uid, contentTag, contentType uint64, contentText string) (uint64, error) {
	if uid < 1 || contentText == "" {
		return 0, errors.New("params fail")
	}

	nowTime := time.Now().Unix()
	dayTime := time.Now().Format("20060102")

	// uid,content_date,content_tag,content_text,content_type,deleted,c_time,m_time
	insertFmt := "insert into %s (%s) values (?,?,?,?,?,?,?,?)"
	execSql := fmt.Sprintf(insertFmt, t.getTableName(), insertField)

	result, err := t.Db.Exec(execSql, uid, dayTime, contentTag, contentText, contentType, DeletedOff, nowTime, nowTime)
	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(row), nil
}

func (t *ContentModel) contentQuery(querySql string) ([]*Content, error) {

	contentList := make([]*Content, 0)
	rows, err := t.Db.Query(querySql)
	if err != nil {
		return nil, err
	}

	var id, uid, content_date, content_tag, content_type, deleted, c_time, m_time uint64
	var content_text string

	for rows.Next() {
		if err := rows.Scan(&id, &uid, &content_date, &content_tag, &content_text, &content_type, &deleted, &c_time, &m_time); err != nil {
			return nil, err
		}

		content := &Content{
			Id:          id,
			Uid:         uid,
			ContentDate: content_date,
			ContentTag:  content_tag,
			ContentText: content_text,
			ContentType: content_type,
			Deleted:     deleted,
			CTime:       c_time,
			MTime:       m_time,
		}

		contentList = append(contentList, content)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contentList, nil
}
