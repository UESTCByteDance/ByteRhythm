package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Favorite struct {
	Id         int       `orm:"column(id);pk;auto" description:"喜欢id"`
	UserId     *User     `orm:"column(user_id);rel(fk)" description:"点赞用户id"`
	VideoId    *Video    `orm:"column(video_id);rel(fk)" description:"点赞视频id"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
}

func (t *Favorite) TableName() string {
	return "favorite"
}

func init() {
	orm.RegisterModel(new(Favorite))
}

func FavoriteData(favorites []Favorite) (data []interface{}) {

	for _, favorite := range favorites {
		data = append(data, map[string]interface{}{
			"id":          favorite.Id,
			"user_id":     favorite.UserId,
			"video_id":    favorite.VideoId,
			"create_time": favorite.CreateTime.Format("2006-1-2 15:04"),
		})
	}
	return
}

// AddFavorite insert a new Favorite into database and returns
// last inserted Id on success.
func AddFavorite(m *Favorite) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFavoriteById retrieves Favorite by Id. Returns error if
// Id doesn't exist
func GetFavoriteById(id int) (v *Favorite, err error) {
	o := orm.NewOrm()
	v = &Favorite{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFavorite retrieves all Favorite matches certain condition. Returns empty list if
// no records exist
func GetAllFavorite(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Favorite))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Favorite
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateFavorite updates Favorite by Id and returns error if
// the record to be updated doesn't exist
func UpdateFavoriteById(m *Favorite) (err error) {
	o := orm.NewOrm()
	v := Favorite{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFavorite deletes Favorite by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFavorite(id int) (err error) {
	o := orm.NewOrm()
	v := Favorite{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Favorite{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
