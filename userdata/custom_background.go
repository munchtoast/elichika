package userdata

import (
	"elichika/utils"
)

func customBackgroundFinalizer(session *Session) {
	for _, userCustomBackground := range session.UserModel.UserCustomBackgroundByID.Objects {
		affected, err := session.Db.Table("u_custom_background").Where("user_id = ? AND custom_background_master_id = ?",
			session.UserStatus.UserID, userCustomBackground.CustomBackgroundMasterID).AllCols().Update(userCustomBackground)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_custom_background").Insert(userCustomBackground)
			utils.CheckErr(err)
		}
	}
}
func init() {
	addFinalizer(customBackgroundFinalizer)
	addGenericTableFieldPopulator("u_custom_background", "UserCustomBackgroundByID")
}
