/**
 * Created by Michael on 2015/8/6.
 */
package db
import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"vo"
	"strconv"
	"time"
)

/**
 * 检测用户是否拥有指定世界
 */
func (this *Slave) CheckUserWorld(account_num int,worldID int) (id int, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	var userWorld int
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return userWorld, err
	}

	err = this.db.QueryRow("SELECT world_id FROM tb_world_user_new WHERE account_num = ? AND world_id = ?",
		account_num,worldID).Scan(&userWorld)
	if err != nil {
		log.Errorln(err)
		return userWorld, err
	}
	return userWorld, nil
}

/**
 * 为用户创建新的世界
 */
func (this *Master) OpenUserWorld(worldVO *vo.UserWorld, buildPartCloud []*vo.BuildPartCloud) (e interface{}){
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return err
	}

	create_time := time.Now().Unix()
	stmt, err := this.db.Prepare("INSERT INTO tb_world_user_new(account_num, world_id,create_time) VALUES(?,?,?)")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(worldVO.AccountNum,worldVO.WorldID, create_time)
	if err != nil {
		log.Errorln(err)
		return err
	}

	valueStr := ""
	for _,value := range buildPartCloud{
		if valueStr == ""{
			valueStr += "(" + strconv.Itoa(value.Buildid) + ","+strconv.Itoa(value.AccountNum) + ",'" +value.Part +"','" +value.Cloud  +"')"
		}else{
			valueStr += ",(" + strconv.Itoa(value.Buildid) + ","+strconv.Itoa(value.AccountNum) + ",'" +value.Part +"','" +value.Cloud  +"')"
		}
	}

	stmt1, err1 := this.db.Prepare("INSERT INTO tb_build_user(build_id, account_num, part, cloud) VALUES " + valueStr)
	if err1 != nil {
		log.Errorln(err1)
		return err1
	}
	defer stmt1.Close()
	_, err1 = stmt1.Exec()
	if err1 != nil {
		log.Errorln(err1)
		return err1
	}

	return nil
}

/**
 * 获取用户世界列表(并标识当前世界)
 */
func (this *Slave) GetWorldList(account_num int) (world []*vo.UserWorld, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	worldList := make([]*vo.UserWorld, 0, 10)
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return worldList, err
	}

	rows, err := this.db.Query("SELECT n.world_id, c.id FROM tb_world_user_new AS n "+
		"LEFT JOIN tb_world_currently AS c ON c.account_num = n.account_num  AND c.world_id = n.world_id WHERE n.account_num = ?", account_num)
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			woldVO := &vo.UserWorld{}
			rows.Scan(&woldVO.WorldID,&woldVO.Currently)
			worldList = append(worldList, woldVO)
		}
	}

	if err != nil {
		log.Errorln(err)
		return worldList, err
	}
	return worldList, nil
}

//世界大楼展示位置（包括大楼生长位置，云消散位置）
func (this *Slave) GetWorldBuildsShow(account_num int, buildArr []string) (showList []*vo.BuildPartCloud, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	partCloudList := make([]*vo.BuildPartCloud, 0, 3)
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return partCloudList, err
	}
	first:= true
	where := ""
	part := "build_id = "
	for _,value := range buildArr{
		if first {
			where += " AND (" + part + value
			first = false;
		}else{
			where += " OR " + part + value
		}
	}
	if where != ""{
		where += ")"
	}

	rows, err := this.db.Query("SELECT account_num, build_id, part, cloud FROM tb_build_user WHERE account_num = ? " + where , account_num)
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			data := &vo.BuildPartCloud{}
			err := rows.Scan(&data.AccountNum, &data.Buildid, &data.Part, &data.Cloud)
			if err != nil {
				log.Errorln(err)
			}
			log.Infoln(data.AccountNum, data.Buildid, data.Part, data.Cloud)
			partCloudList = append(partCloudList, data)
		}
	}

	if err != nil {
		log.Errorln(err)
		return partCloudList, err
	}
	return partCloudList, nil
}

// 获取大楼伸展位置
func (this *Slave)GetBuildPart(account_num int, build_id int) (part string, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	var partStr string
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return partStr, err
	}

	err = this.db.QueryRow("SELECT part FROM tb_build_user WHERE account_num = ? AND build_id = ?",
		account_num,build_id).Scan(&partStr)
	if err != nil {
		log.Errorln(err)
		return partStr,err
	}
	return partStr, nil
}

// 获取大楼云消散位置
func (this *Slave)GetBuildCloud(account_num int, buildArr []string) (showList []*vo.BuildPartCloud, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	partCloudList := make([]*vo.BuildPartCloud, 0, 3)
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return partCloudList, err
	}
	first:= true
	where := ""
	part := "build_id = "
	for _,value := range buildArr{
		if first {
			where += " AND (" + part + value
			first = false;
		}else{
			where += " OR " + part + value
		}
	}
	if where != ""{
		where += ")"
	}
	rows, err := this.db.Query("SELECT build_id, cloud FROM tb_build_user WHERE account_num = ? " + where , account_num)
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			data := &vo.BuildPartCloud{}
			err := rows.Scan(&data.Buildid, &data.Cloud)
			if err != nil {
				log.Errorln(err)
			}
			partCloudList = append(partCloudList, data)
		}
	}

	if err != nil {
		log.Errorln(err)
		return partCloudList, err
	}
	return partCloudList, nil
}

// 更改大楼伸展位置
func (this *Master) UpdateBuildPart(account_num int, build_id int, part string) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	err := this.db.Ping()
	if err != nil {
		return err
	}

	stmt, err := this.db.Prepare("UPDATE tb_build_user SET part = ? WHERE account_num = ? AND build_id = ?")

	if stmt != nil {
		defer stmt.Close()
		_, err = stmt.Exec(part, account_num, build_id)
		if err != nil {
			return err
		}
	}else {
		return err
	}

	return nil
}

// 更改大楼云消散位置
func (this *Master) UpdateBuildCloud(account_num int, cloudList []*vo.BuildPartCloud) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	err := this.db.Ping()
	if err != nil {
		return err
	}

	sqlStr := "UPDATE tb_build_user SET cloud = CASE build_id "
	where := " END WHERE account_num = " + strconv.Itoa(account_num) + " AND "
	paramStr := ""
	filter := "("
	for _,value:= range cloudList{
		buildID := strconv.Itoa(value.Buildid)
		paramStr += " WHEN " + buildID + " THEN '" + value.Cloud + "' "
		if filter == "(" {
			filter += " build_id = " + buildID
		}else{
			filter += " OR build_id = " + buildID
		}

	}
	filter += ")"
	stmt, err := this.db.Prepare(sqlStr + paramStr + where + filter)

	if stmt != nil {
		defer stmt.Close()
		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	}else {
		return err
	}
	return nil
}

