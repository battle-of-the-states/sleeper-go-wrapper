package sleeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStats(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	playerID := "4039"

	// Cooper Kupp
	res, err := client.UnstableGetPlayerInfo(NFL, PlayerStats, playerID, SeasonTypeOpt(RegularSeason), SeasonOpt("2021"), WeekOpt(9))

	if err != nil {
		t.Error(err)
	}

	require.Equal(t, LosAngelesRams, NFLTeamAbbreviation(res.Team))
	require.Equal(t, NFL, res.Sport)
	require.Equal(t, RegularSeason, SeasonType(res.SeasonType))
	require.Equal(t, "2021", res.Season)
	require.Equal(t, playerID, res.PlayerID)
	require.Equal(t, Tennessee, NFLTeamAbbreviation(*res.Opponent))
	require.Equal(t, "202110932", res.GameID)
	require.Equal(t, "2021-11-07", *res.Date)
	require.Equal(t, "sportradar", res.Company)
	require.Equal(t, "stat", res.Category)

	stats := res.Stats

	require.Equal(t, float32(23), stats["tm_st_snp"])
	require.Equal(t, float32(4), stats["st_snp"])
	require.Equal(t, float32(7.31), stats["rec_ypt"])
	require.Equal(t, float32(8.64), stats["rec_ypr"])
	require.Equal(t, float32(71), stats["rec_yar"])
	require.Equal(t, float32(13), stats["rec_tgt"])
}

func TestGetPositionStats(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	season := "2022"
	week := 7

	// Cooper Kupp
	res, err := client.UnstableGetPlayersInfoByPosition(NFL, PlayerStats, season, week, SeasonTypeOpt(RegularSeason), PositionOpt(QuarterBack))

	if err != nil {
		t.Error(err)
	}

	qb1Info := res[0]

	require.Equal(t, Tennessee, NFLTeamAbbreviation(qb1Info.Team))
	require.Equal(t, NFL, qb1Info.Sport)
	require.Equal(t, RegularSeason, SeasonType(qb1Info.SeasonType))
	require.Equal(t, "2022", qb1Info.Season)
	require.Equal(t, "1049", qb1Info.PlayerID)
	require.Equal(t, Indianapolis, NFLTeamAbbreviation(*qb1Info.Opponent))
	require.Equal(t, "202210734", qb1Info.GameID)
	require.Equal(t, "2022-10-23", *qb1Info.Date)
	require.Equal(t, "sportradar", qb1Info.Company)
	require.Equal(t, "stat", qb1Info.Category)

	qb1 := qb1Info.Player

	require.Equal(t, 11, qb1.YearsExp)
	require.Equal(t, Tennessee, NFLTeamAbbreviation(qb1.Team))
	require.Equal(t, QuarterBack, FFPosition(qb1.Position))
	require.Equal(t, int64(1672349449646), qb1.NewsUpdated)
	require.Equal(t, "Active", (*qb1.Metadata)["injury_override_regular_2020_5"])
	require.Equal(t, "Tannehill", qb1.LastName)
	require.Equal(t, "IR", *qb1.InjuryStatus)
	require.Nil(t, qb1.InjuryStartDate)
	require.Nil(t, qb1.InjuryNotes)
	require.Nil(t, qb1.InjuryBodyPart)
	require.Equal(t, "Ryan", qb1.FirstName)
	require.Equal(t, QuarterBack, FFPosition(qb1.FantasyPositions[0]))

	stats := qb1Info.Stats

	require.Equal(t, float32(84), stats["pass_air_yd"])
	require.Equal(t, float32(27), stats["pass_lng"])
	require.Equal(t, float32(58), stats["off_snp"])
	require.Equal(t, float32(65), stats["cmp_pct"])
	require.Equal(t, float32(4), stats["rush_att"])
	require.Equal(t, float32(8), stats["rush_lng"])
}

func TestGetProjections(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	playerID := "4039"

	// Cooper Kupp
	res, err := client.UnstableGetPlayerInfo(NFL, PlayerProjections, playerID, SeasonTypeOpt(RegularSeason), SeasonOpt("2021"), WeekOpt(9))

	if err != nil {
		t.Error(err)
	}

	require.Equal(t, LosAngelesRams, NFLTeamAbbreviation(res.Team))
	require.Equal(t, NFL, res.Sport)
	require.Equal(t, RegularSeason, SeasonType(res.SeasonType))
	require.Equal(t, "2021", res.Season)
	require.Equal(t, playerID, res.PlayerID)
	require.Equal(t, Tennessee, NFLTeamAbbreviation(*res.Opponent))
	require.Equal(t, "202110932", res.GameID)
	require.Equal(t, "2021-11-07", *res.Date)
	require.Equal(t, "rotowire", res.Company)
	require.Equal(t, "proj", res.Category)

	stats := res.Stats

	require.Equal(t, float32(4.1), stats["rush_yd"])
	require.Equal(t, float32(0.41), stats["rush_fd"])
	require.Equal(t, float32(0.66), stats["rush_att"])
	require.Equal(t, float32(112.44), stats["rec_yd"])
	require.Equal(t, float32(11.1), stats["rec_tgt"])
	require.Equal(t, float32(0.67), stats["rec_td"])
}

func TestGetPositionProjections(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	season := "2022"
	week := 7

	// Cooper Kupp
	res, err := client.UnstableGetPlayersInfoByPosition(NFL, PlayerProjections, season, week, SeasonTypeOpt(RegularSeason), PositionOpt(QuarterBack))

	if err != nil {
		t.Error(err)
	}

	qb1Info := res[0]

	require.Equal(t, Tennessee, NFLTeamAbbreviation(qb1Info.Team))
	require.Equal(t, NFL, qb1Info.Sport)
	require.Equal(t, RegularSeason, SeasonType(qb1Info.SeasonType))
	require.Equal(t, "2022", qb1Info.Season)
	require.Equal(t, "1049", qb1Info.PlayerID)
	require.Equal(t, Indianapolis, NFLTeamAbbreviation(*qb1Info.Opponent))
	require.Equal(t, "202210734", qb1Info.GameID)
	require.Equal(t, "2022-10-23", *qb1Info.Date)
	require.Equal(t, "rotowire", qb1Info.Company)
	require.Equal(t, "proj", qb1Info.Category)

	qb1 := qb1Info.Player

	require.Equal(t, 11, qb1.YearsExp)
	require.Equal(t, Tennessee, NFLTeamAbbreviation(qb1.Team))
	require.Equal(t, QuarterBack, FFPosition(qb1.Position))
	require.Equal(t, int64(1672349449646), qb1.NewsUpdated)
	require.Equal(t, "Active", (*qb1.Metadata)["injury_override_regular_2020_5"])
	require.Equal(t, "Tannehill", qb1.LastName)
	require.Equal(t, "IR", *qb1.InjuryStatus)
	require.Nil(t, qb1.InjuryStartDate)
	require.Nil(t, qb1.InjuryNotes)
	require.Nil(t, qb1.InjuryBodyPart)
	require.Equal(t, "Ryan", qb1.FirstName)
	require.Equal(t, QuarterBack, FFPosition(qb1.FantasyPositions[0]))

	stats := qb1Info.Stats

	require.Equal(t, float32(17.11), stats["rush_yd"])
	require.Equal(t, float32(0.17), stats["rush_td"])
	require.Equal(t, float32(1.71), stats["rush_fd"])
	require.Equal(t, float32(16.97), stats["pass_cmp"])
	require.Equal(t, float32(3.8), stats["rush_att"])
	require.Equal(t, float32(203.25), stats["pass_yd"])
}
