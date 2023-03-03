package sleeper

// Sport ensures users give valid options for sports
type Sport string

// Bracket ensures users give valid options for the playoff brackets
type Bracket string

// SeasonType are the fantasy sport season types Sleeper expects ("pre", "post", or "regular")
type SeasonType string

// LeagueStatus is the current status of the league ("pre_draft", "drafting", "in_season", or "complete")
type LeagueStatus string

// PlayerInfoType gives the option between stats and projections
type PlayerInfoType string

// FFPosition gives all positions Sleeper recognizes for Fantasy Football
type FFPosition string

// NFLTeamAbbreviation gives all NFL team name abbreviations
type NFLTeamAbbreviation string

// TransactionType are the different transaction types recognized by Sleeper
type TransactionType string

const (
	SLEEPER_BASE_URL             = "https://api.sleeper.app/v1"
	SLEEPER_BASE_URL_UNSTABLE    = "https://api.sleeper.app"
	SLEEPER_BASE_CDN             = "https://sleepercdn.com/"
	SLEEPER_AVATAR_URL           = "https://sleepercdn.com/avatars"
	SLEEPER_AVATAR_THUMBNAIL_URL = "https://sleepercdn.com/avatars/thumbs"

	// NFL is the nfl sport tag
	NFL Sport = "nfl"
	// NBA is the nba sport tag
	NBA Sport = "nba"
	// LCS is the lcs sport tag
	LCS Sport = "lcs"
	// MLB is the mlb sport tag
	MLB Sport = "mlb"

	// WinnersBracket is the winners bracket
	WinnersBracket Bracket = "winners_bracket"
	// LosersBracket is the losers bracket
	LosersBracket Bracket = "losers_bracket"

	// Preseason is the season type for the preseason
	Preseason SeasonType = "pre"
	// Postseason is the season type for the postseason
	Postseason SeasonType = "post"
	// RegularSeason is the season type for the regular season
	RegularSeason SeasonType = "regular"
	// OffSeason is the season type for the off season
	OffSeason SeasonType = "off"

	// StatusPreDraft is the league status for pre_draft
	StatusPreDraft LeagueStatus = "pre_draft"
	// StatusDrafting is the league status for drafting
	StatusDrafting LeagueStatus = "drafting"
	// StatusInSeason is the league status for in_season
	StatusInSeason LeagueStatus = "in_season"
	// StatusComplete is the league status for complete
	StatusComplete LeagueStatus = "complete"

	// TrnTrade is a trade transaction
	TrnTrade TransactionType = "trade"
	// TrnFreeAgent is a free agent transaction
	TrnFreeAgent TransactionType = "free_agent"
	// TrnWaiver is a waiver transaction
	TrnWaiver TransactionType = "waiver"

	// PlayerStats is the expected value by Sleeper for stats endpoints
	PlayerStats PlayerInfoType = "stats"
	// PlayerProjections is the expected value by Sleeper for projections endpoints
	PlayerProjections PlayerInfoType = "projections"

	QuarterBack  FFPosition = "QB"
	RunningBack  FFPosition = "RB"
	WideReceiver FFPosition = "WR"
	TightEnd     FFPosition = "TE"
	Kicker       FFPosition = "K"
	Defense      FFPosition = "DEF"

	Arizona            NFLTeamAbbreviation = "ARI"
	Atlanta            NFLTeamAbbreviation = "ATL"
	Baltimore          NFLTeamAbbreviation = "BAL"
	Buffalo            NFLTeamAbbreviation = "BUF"
	Carolina           NFLTeamAbbreviation = "CAR"
	Chicago            NFLTeamAbbreviation = "CHI"
	Cincinnati         NFLTeamAbbreviation = "CIN"
	Cleveland          NFLTeamAbbreviation = "CLE"
	Dallas             NFLTeamAbbreviation = "DAL"
	Denver             NFLTeamAbbreviation = "DEN"
	Detroit            NFLTeamAbbreviation = "DET"
	GreenBay           NFLTeamAbbreviation = "GB"
	Houston            NFLTeamAbbreviation = "HOU"
	Indianapolis       NFLTeamAbbreviation = "IND"
	Jacksonville       NFLTeamAbbreviation = "JAX"
	KansasCity         NFLTeamAbbreviation = "KC"
	LosAngelesChargers NFLTeamAbbreviation = "LAC"
	LosAngelesRams     NFLTeamAbbreviation = "LAR"
	LasVegas           NFLTeamAbbreviation = "LV"
	Miami              NFLTeamAbbreviation = "MIA"
	Minnesota          NFLTeamAbbreviation = "MIN"
	NewEngland         NFLTeamAbbreviation = "NE"
	NewOrleans         NFLTeamAbbreviation = "NO"
	NewYorkGiants      NFLTeamAbbreviation = "NYG"
	NewYorkJets        NFLTeamAbbreviation = "NYJ"
	Philadelphia       NFLTeamAbbreviation = "PHI"
	Pittsburgh         NFLTeamAbbreviation = "PIT"
	Seattle            NFLTeamAbbreviation = "SEA"
	SanFrancisco       NFLTeamAbbreviation = "SF"
	TampaBay           NFLTeamAbbreviation = "TB"
	Tennessee          NFLTeamAbbreviation = "TEN"
	Washington         NFLTeamAbbreviation = "WAS"
)
