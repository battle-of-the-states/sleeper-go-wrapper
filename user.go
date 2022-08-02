package sleeper

import "fmt"

type UserSimpleJSON struct {
	Username    string `json:"username"`
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email,omitempty"`
}

// UsersJSON is the return type for all users in a league
type UsersJSON []UserJSON

// UserJSON is a single user from the league users API.
type UserJSON struct {
	UserID   string      `json:"user_id"`
	Settings interface{} `json:"settings"`
	Metadata struct {
		TeamName                string `json:"team_name"`
		TeamNameUpdate          string `json:"team_name_update"`
		MentionPn               string `json:"mention_pn"`
		AllowPn                 string `json:"allow_pn"`
		TradeBlockPn            string `json:"trade_block_pn"`
		JoinVoicePn             string `json:"join_voice_pn"`
		UserMessagePn           string `json:"user_message_pn"`
		TransactionCommissioner string `json:"transaction_commissioner"`
		Archived                string `json:"archived"`
		TransactionFreeAgent    string `json:"transaction_free_agent"`
		TransactionTrade        string `json:"transaction_trade"`
		TransactionWaiver       string `json:"transaction_waiver"`
		PlayerLikePn            string `json:"player_like_pn"`
		MascotMessage           string `json:"mascot_message"`
		PlayerNicknameUpdate    string `json:"player_nickname_update"`

		// tf are these
		MascotItemTypeIDLeg1      string `json:"mascot_item_type_id_leg_1"`
		MascotItemTypeIDLeg2      string `json:"mascot_item_type_id_leg_2"`
		MascotItemTypeIDLeg3      string `json:"mascot_item_type_id_leg_3"`
		MascotItemTypeIDLeg4      string `json:"mascot_item_type_id_leg_4"`
		MascotItemTypeIDLeg5      string `json:"mascot_item_type_id_leg_5"`
		MascotItemTypeIDLeg6      string `json:"mascot_item_type_id_leg_6"`
		MascotItemTypeIDLeg7      string `json:"mascot_item_type_id_leg_7"`
		MascotItemTypeIDLeg8      string `json:"mascot_item_type_id_leg_8"`
		MascotItemTypeIDLeg9      string `json:"mascot_item_type_id_leg_9"`
		MascotItemTypeIDLeg10     string `json:"mascot_item_type_id_leg_10"`
		MascotItemTypeIDLeg11     string `json:"mascot_item_type_id_leg_11"`
		MascotItemTypeIDLeg12     string `json:"mascot_item_type_id_leg_12"`
		MascotItemTypeIDLeg13     string `json:"mascot_item_type_id_leg_13"`
		MascotItemTypeIDLeg14     string `json:"mascot_item_type_id_leg_14"`
		MascotItemTypeIDLeg15     string `json:"mascot_item_type_id_leg_15"`
		MascotItemTypeIDLeg16     string `json:"mascot_item_type_id_leg_16"`
		MascotItemTypeIDLeg17     string `json:"mascot_item_type_id_leg_17"`
		MascotMessageEmotionLeg1  string `json:"mascot_message_emotion_leg_1"`
		MascotMessageEmotionLeg2  string `json:"mascot_message_emotion_leg_2"`
		MascotMessageEmotionLeg3  string `json:"mascot_message_emotion_leg_3"`
		MascotMessageEmotionLeg4  string `json:"mascot_message_emotion_leg_4"`
		MascotMessageEmotionLeg5  string `json:"mascot_message_emotion_leg_5"`
		MascotMessageEmotionLeg6  string `json:"mascot_message_emotion_leg_6"`
		MascotMessageEmotionLeg7  string `json:"mascot_message_emotion_leg_7"`
		MascotMessageEmotionLeg8  string `json:"mascot_message_emotion_leg_8"`
		MascotMessageEmotionLeg9  string `json:"mascot_message_emotion_leg_9"`
		MascotMessageEmotionLeg10 string `json:"mascot_message_emotion_leg_10"`
		MascotMessageEmotionLeg11 string `json:"mascot_message_emotion_leg_11"`
		MascotMessageEmotionLeg12 string `json:"mascot_message_emotion_leg_12"`
	} `json:"metadata"`
	LeagueID    string `json:"league_id"`
	IsOwner     bool   `json:"is_owner"`
	IsBot       bool   `json:"is_bot"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

/*
Via the user resource, you can GET the user object by either providing the
username or user_id of the user.

** Keep in mind that the username of a user can change over time, so if you are
storing information, you'll want to hold onto the user_id.**

usernameOrID (Required) : The user you want to fetch.
*/
func (c *Client) GetUser(usernameOrID string) (UserSimpleJSON, error) {
	// https://api.sleeper.app/v1/user/<username> or https://api.sleeper.app/v1/user/<user_id>
	lastfmURL := fmt.Sprintf("%s/user/%s", c.sleeperURL, usernameOrID)

	user := new(UserSimpleJSON)

	err := c.get(lastfmURL, user)

	if err != nil {
		return *user, err
	}

	return *user, nil
}
