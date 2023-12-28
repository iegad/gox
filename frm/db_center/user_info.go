package dbcenter

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/iegad/gox/frm/biz"
	"github.com/iegad/gox/frm/utils"
)

type UserInfo struct {
	UserID     int64  `json:"user_id"`
	ChannelID  int32  `json:"channel_id"`
	UserType   byte   `json:"user_type"`
	UserName   string `json:"user_name"`
	Password   string `json:"-"`
	Nickname   string `json:"nickname"`
	RegistTime int64  `json:"regist_time"`
	LastUpdate int64  `json:"-"`
	State      byte   `json:"state"`
	Token      string `json:"token"`
}

func maxUserID(db *sql.DB) (int64, error) {
	row := db.QueryRow("SELECT IFNULL(MAX(F_USER_ID), 100000) FROM DB_CENTER.T_USER_INFO")
	maxID := int64(0)
	err := row.Scan(&maxID)
	if err != nil {
		return -1, err
	}

	return maxID, nil
}

func NewGuest(channelID int32, db *sql.DB) (*UserInfo, error) {
	maxID, err := maxUserID(db)
	if err != nil {
		return nil, err
	}

	userID := maxID + 1
	tnow := time.Now().Unix()
	return &UserInfo{
		UserID:     userID,
		ChannelID:  channelID,
		UserType:   0,
		UserName:   fmt.Sprintf("guest_%v", userID),
		Password:   utils.MD5Hex("888888"),
		Nickname:   fmt.Sprintf("GUEST_%v", userID),
		RegistTime: tnow,
		LastUpdate: tnow,
		State:      1,
	}, nil
}

func (this_ *UserInfo) String() string {
	jstr, _ := json.Marshal(this_)
	return string(jstr)
}

func (this_ *UserInfo) Insert(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO DB_CENTER.T_USER_INFO(F_USER_ID,F_CHANNEL_ID,F_USER_TYPE,F_USER_NAME,F_PASSWORD,F_NICKNAME,F_REGIST_TIME,F_LAST_UPDATE,F_STATE)VALUES(?,?,?,?,?,?,?,?,?)",
		this_.UserID, this_.ChannelID, this_.UserType, this_.UserName, this_.Password, this_.Nickname, this_.RegistTime, this_.LastUpdate, this_.State)
	return err
}

func (this_ *UserInfo) SetSession(r *redis.Client) error {
	this_.Token = uuid.NewString()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()
	err := r.Set(ctx, biz.GetUserSessionKey(this_.UserID), this_.String(), time.Hour*24*30).Err()
	if err != nil {
		return err
	}

	return nil
}
