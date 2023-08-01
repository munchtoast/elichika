package serverdb

import (
	"elichika/config"
	"elichika/model"
	"elichika/utils"

	"encoding/json"
	"fmt"
	// "github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func DbGetUserData(fileName string) string {
	userDataFile := config.UserDataPath + fileName
	if utils.PathExists(userDataFile) {
		return utils.ReadAllText(userDataFile)
	}

	presetDataFile := config.PresetDataPath + fileName
	if !utils.PathExists(presetDataFile) {
		panic("File not exists")
	}

	userData := utils.ReadAllText(presetDataFile)
	utils.WriteAllText(userDataFile, userData)

	return userData
}

// A session is a complete transation between server and client
// so 1 session per request
// A session fetch the data needs to be modified.
type Session struct {
	UserInfo            model.UserInfo
	CardDiffs           map[int]model.CardInfo
	UserMemberDiffs     map[int]model.UserMemberInfo
	UserLessonDeckDiffs map[int]model.UserLessonDeck
	UserLiveDeckDiffs   map[int]model.UserLiveDeck
	UserLivePartyDiffs  map[int]model.UserLiveParty
	CardGradeUpTriggers []any
}

// Push update into the db and create the diff
// The actual response depend on the API, but they often contain the diff somewhere
// The mainKey is the key to the diff
func (session *Session) Finalize(jsonBody string, mainKey string) string {
	// jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_status", session.UserInfo)
	// modelDiff, _ = sjson.Set(modelDiff, mainKey+".user_status.gdpr_version", 4)
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_member_by_member_id", session.FinalizeUserMemberDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_card_by_card_id", session.FinalizeCardDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_info_trigger_card_grade_up_by_trigger_id", session.FinalizeCardGradeUpTrigger())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_lesson_deck_by_id", session.FinalizeUserLessonDeckDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_live_deck_by_id", session.FinalizeUserLiveDeckDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_live_party_by_id", session.FinalizeUserLivePartyDiffs())
	return jsonBody
}

// fetch the user, this is always sent back to client
func (session *Session) InitUser(userID int) {
	session.UserInfo.UserID = userID
	exists, err := Engine.Table("s_user_info").Where("user_id = ?", userID).Get(&session.UserInfo)
	if err != nil {
		panic(err)
	}
	if !exists { // create one
		fmt.Println("Insert new user: ", userID)
		data := utils.ReadAllText("assets/userdata/userStatus.json")
		if err := json.Unmarshal([]byte(data), &session.UserInfo); err != nil {
			panic(err)
		}

		// insert into the db
		_, err := Engine.Table("s_user_info").AllCols().Insert(&session.UserInfo)
		if err != nil {
			panic(err)
		}
	}
}

func GetSession(userId int) Session {
	s := Session{}
	s.CardDiffs = make(map[int]model.CardInfo)
	s.UserMemberDiffs = make(map[int]model.UserMemberInfo)
	s.UserLessonDeckDiffs = make(map[int]model.UserLessonDeck)
	s.UserLiveDeckDiffs = make(map[int]model.UserLiveDeck)
	s.UserLivePartyDiffs = make(map[int]model.UserLiveParty)
	s.CardGradeUpTriggers = make([]any, 0)
	s.InitUser(userId)
	return s
}
