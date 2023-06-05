package main

import (
	"encoding/json"
	"fmt"
)

func CreateGameSession(token string) (string, string, error) {
	body := &CreateGameSessionBody{
		AliasId:               GlobalConfig.AliasId,
		FleetId:               GlobalConfig.FleetId,
		Name:                  "AutoCreateGameSession",
		CreatorId:             "",
		MaxClientSessionCount: 20,
		IdempotencyToken:      "",
		ServerSessionData:     "",
	}
	url := fmt.Sprintf(CreateGameSessionUrlPattern, GlobalConfig.FleetManagerAddr, GlobalConfig.FleetManagerPort, GlobalConfig.ProjectId)
	rsp, err := HttpClient.Post(url, body, token)
	if err != nil {
		Logger.Errorf("create game session failed: %s", err)
		return "", "", nil
	}
	session := &RespServerSession{}
	err = json.Unmarshal(rsp, session)
	if err != nil {
		Logger.Errorf("failed to unmarshal game session: %s", err)
		return "", "", nil
	}
	if session.ServerSessions.ServerSessionId == "" {
		Logger.Errorf("failed to create game session: %s", rsp)
		return "", "", fmt.Errorf("failed to create game session: %s", rsp)
	}
	return session.ServerSessions.ServerSessionId, session.ServerSessions.FleetId, nil
}

func QueryGameSession(gameSessionId string, fleetId string, token string) (string, error) {
	url := fmt.Sprintf(QueryGameSessionUrlPattern, GlobalConfig.FleetManagerAddr, GlobalConfig.FleetManagerPort, GlobalConfig.ProjectId, gameSessionId)
	queryParam := make(map[string]string)
	queryParam["fleet_id"] = fleetId
	rsp, err := HttpClient.Get(url, queryParam, token)
	if err != nil {
		Logger.Errorf("failed to get game session request: %s", err)
		return "", nil
	}
	session := &RespServerSession{}
	err = json.Unmarshal(rsp, session)
	if err != nil {
		Logger.Errorf("failed to unmarshal game session request: %s", err)
		return "", nil
	}
	return session.ServerSessions.State, nil
}

func Login() string {
	loginBody := &LoginBody{
		UserName: GlobalConfig.LoginName,
		Password: GlobalConfig.LoginPassword,
	}
	url := fmt.Sprintf(LoginUrlPattern, GlobalConfig.FleetManagerAddr, GlobalConfig.FleetManagerPort)
	rsp, err := HttpClient.Post(url, loginBody, "")
	if err != nil {
		Logger.Errorf("failed to get game session request: %s", err)
		return ""
	}
	loginResp := &LoginResp{}
	err = json.Unmarshal(rsp, loginResp)
	if err != nil {
		Logger.Errorf("failed to unmarshal game session request: %s", err)
		return ""
	}
	Logger.Infof("login success")
	return loginResp.AuthToken
}
