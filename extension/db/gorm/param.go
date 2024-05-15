/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Param struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	DBName      string `yaml:"dbname"`
	MaxOpenConn int    `yaml:"max-open-conn"`
	MaxIdleConn int    `yaml:"max-idle-conn"`
}

func (c *Param) New(mysqlImpl *GORMDB) (*GORMDB, error) {
	rawDB, err := gorm.Open(mysql.Open(getMysqlLinkStr(c)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db, err := rawDB.DB()
	if err != nil {
		return nil, err
	}

	if c.MaxOpenConn != 0 {
		db.SetMaxOpenConns(c.MaxOpenConn)
	}
	if c.MaxIdleConn != 0 {
		db.SetMaxIdleConns(c.MaxIdleConn)
	}
	db.SetConnMaxLifetime(time.Hour)

	mysqlImpl.db = rawDB
	return mysqlImpl, err
}

func getMysqlLinkStr(conf *Param) string {
	return conf.Username + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.DBName +
		"?charset=utf8&parseTime=True&loc=Local"
}
