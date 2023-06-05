package main


type Config struct {
	FleetManagerAddr 	string
	FleetManagerPort	string
	ProjectId			string

	AliasId				string
	FleetId				string
	LoginName			string
	LoginPassword		string

	BatchCreateGameSession int
	CreateGameSessionSleepSecond int

	ProcessRunSleepSeconds		int

}

const (
	DefaultLoginName = "admin"
	DefaultLoginPassword = ""
	DefaultBatchCreateGameSession = 10
	DefaultCreateGameSessionSleepSecond = 2
	DefaultProcessRunSleepSeconds = 500
)

const (
	CreateGameSessionUrlPattern = "https://%s:%s/v1/%s/server-sessions"
	QueryGameSessionUrlPattern = "https://%s:%s/v1/%s/server-sessions/%s"
	LoginUrlPattern = "https://%s:%s/v1/user/login"
)

const (
	GameSessionActiveState = "ACTIVE"
	GameSessionErrorState = "ERROR"
)

var GlobalConfig = &Config{}
var HttpClient = &Https{}

type CreateGameSessionBody struct {
	FleetId                 string     `json:"fleet_id" validate:"required,min=1,max=64"`
	AliasId                 string     `json:"alias_id" validate:"required,min=1,max=64"`
	CreatorId               string     `json:"creator_id" validate:"min=0,max=1024"`
	Name                    string     `json:"name" validate:"min=0,max=1024"`
	MaxClientSessionCount     int        `json:"max_client_session_count" validate:"required,gte=1,lte=1024"`
	IdempotencyToken        string     `json:"idempotency_token" validate:"min=0,max=48"`
	ServerSessionData       string     `json:"server_session_data" validate:"min=0,max=4096"`
	ServerSessionProperties []Property `json:"server_session_properties" validate:"omitempty,dive,min=0,max=16"`
}

type Property struct {
	Key   string `json:"key" validate:"required,min=1,max=64"`
	Value string `json:"value" validate:"required,min=1,max=1024"`
}

type ServerSession struct {
	ServerSessionId                         string     `json:"server_session_id"`
	Name                                    string     `json:"name"`
	CreatorId                               string     `json:"creator_id"`
	FleetId                                 string     `json:"fleet_id"`
	Properties                              []Property `json:"server_session_properties"`
	ServerSessionData                       string     `json:"server_session_data"`
	CurrentClientSessionCount               int        `json:"current_client_session_count"`
	MaxClientSessionCount                   int        `json:"max_client_session_count"`
	State                                   string     `json:"state"`
	StateReason                             string     `json:"state_reason"`
	IpAddress                               string     `json:"ip_address"`
	Port                                    int        `json:"port"`
	ClientSessionCreationPolicy             string     `json:"client_session_creation_policy"`
	ServerSessionProtectionPolicy           string     `json:"server_session_protection_policy"`
	ServerSessionProtectionTimeLimitMinutes int        `json:"server_session_protection_time_limit_minutes"`
}

type RespServerSession struct {
	ServerSessions ServerSession 	`json:"server_session"`
}

type LoginBody struct {
	UserName  string 		`json:"username"`
	Password  string		`json:"password"`
}

type LoginResp struct {
	UserName		string 		`json:"username"`
	Id				string		`json:"id"`
	AuthToken		string		`json:"Auth-Token"`
	Activation		int			`json:"Activation"`
	UserType		int			`json:"UserType"`
}