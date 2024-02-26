package api

import (
	"context"
	"encoding/json"
	"fake-server/config"
	"fake-server/gsemanager"
	"fake-server/logger"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

const (
	SUCCESS      = 0
	SUCCESSMSG   = "success"
	DelaySeconds = 5
)

type httpProcess struct {
	httpPort     int
	HttpPortChan chan int
}

func NewHttpProcess() *httpProcess {
	h := &httpProcess{
		httpPort:     0,
		HttpPortChan: make(chan int),
	}
	return h
}

type response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (h *httpProcess) registerApi() {
	http.HandleFunc("/gse/login", h.Login)
	http.HandleFunc("/gse/logout", h.LoginOut)
	http.HandleFunc("/gse/terminate-game-server-session", h.TerminateSession)
	http.HandleFunc("/gse/end-process", h.EndProcess)
	http.HandleFunc("/gse/describe-player-sessions", h.DescribePlayerSessions)
	http.HandleFunc("/gse/update-player-session-policy", h.UpdatePlayerSessionCreationPolicy)
	http.HandleFunc("/gse/report-custom-data", h.ReportCustomData)
	http.HandleFunc("/gse/set-process-health-status", h.SetHealthStatus)
	http.HandleFunc("/", h.HelloWorld)
}

func (h *httpProcess) StartHttpServer() error {
	port := h.GenerateHttpRandomPort(config.GlobalConfig.HttpStartPort, config.GlobalConfig.HttpEndPort)
	listen, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Logger.Errorf("http fail to listen", zap.Error(err))
		return fmt.Errorf("http fail to listen")
	}
	addr := listen.Addr().String()
	portStr := strings.Split(addr, ":")[1]
	h.httpPort, err = strconv.Atoi(portStr)
	if err != nil {
		logger.Logger.Errorf("http fail to get port", zap.Error(err))
		return fmt.Errorf("http fail to get port")
	}
	logger.Logger.Infof("http listen port is", zap.Int("port", h.httpPort))

	h.registerApi()
	logger.Logger.Infof("start http server success")
	go http.Serve(listen, nil)
	return nil
}

func (h *httpProcess) GetHttpPort() int {
	return h.httpPort
}

func (h *httpProcess) writeResp(code int32, message string, result interface{}) (string, error) {
	resp := &response{
		Code:    code,
		Message: message,
		Result:  result,
	}

	resultStr, err := json.Marshal(resp)
	return string(resultStr), err
}

func (h *httpProcess) getContext() context.Context {
	return context.Background()
}

func (h *httpProcess) Login(w http.ResponseWriter, req *http.Request) {
	playSessionId := req.URL.Query().Get("playerSessionId")
	gameSessionId := req.URL.Query().Get("gameSessionId")

	if playSessionId == "" {
		resp, _ := h.writeResp(http.StatusBadRequest, "playerSessionId cant be empty", nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	if gameSessionId == "" {
		resp, _ := h.writeResp(http.StatusBadRequest, "gameSessionId cant be empty", nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	gseManager := gsemanager.GetGseManager()
	_, err := gseManager.AcceptPlayerSession(playSessionId, gameSessionId)

	if err != nil {
		code := int32(http.StatusInternalServerError)
		errMsg := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errMsg = st.Message()
			code = int32(st.Code())
		}
		resp, _ := h.writeResp(code, errMsg, nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	successMsg, _ := h.writeResp(SUCCESS, SUCCESSMSG, nil)
	fmt.Fprintf(w, "%s", successMsg)
}

func (h *httpProcess) LoginOut(w http.ResponseWriter, req *http.Request) {
	playSessionId := req.URL.Query().Get("playerSessionId")
	gameSessionId := req.URL.Query().Get("gameSessionId")

	if playSessionId == "" {
		resp, _ := h.writeResp(http.StatusBadRequest, "playerSessionId cant be empty", nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	gseManager := gsemanager.GetGseManager()
	_, err := gseManager.RemovePlayerSession(playSessionId, gameSessionId)
	if err != nil {
		code := int32(http.StatusInternalServerError)
		errMsg := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errMsg = st.Message()
			code = int32(st.Code())
		}
		resp, _ := h.writeResp(code, errMsg, nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	successMsg, _ := h.writeResp(SUCCESS, SUCCESSMSG, nil)
	fmt.Fprintf(w, "%s", successMsg)
}

func (h *httpProcess) TerminateSession(w http.ResponseWriter, req *http.Request) {
	gseManager := gsemanager.GetGseManager()
	gameSessionId := req.URL.Query().Get("gameSessionId")
	_, err := gseManager.TerminateGameServerSession(gameSessionId)

	if err != nil {
		code := int32(http.StatusInternalServerError)
		errMsg := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errMsg = st.Message()
			code = int32(st.Code())
		}
		resp, _ := h.writeResp(code, errMsg, nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	successMsg, _ := h.writeResp(SUCCESS, SUCCESSMSG, nil)
	fmt.Fprintf(w, "%s", successMsg)
}

func (h *httpProcess) EndProcess(w http.ResponseWriter, req *http.Request) {
	gseManager := gsemanager.GetGseManager()
	_, err := gseManager.ProcessEnding()
	if err != nil {
		code := int32(http.StatusInternalServerError)
		errMsg := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errMsg = st.Message()
			code = int32(st.Code())
		}
		resp, _ := h.writeResp(code, errMsg, nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	successMsg, _ := h.writeResp(SUCCESS, SUCCESSMSG, nil)
	fmt.Fprintf(w, "%s", successMsg)
	logger.Logger.Infof("process will end delay %d seconds", DelaySeconds)
	h.EndingProcessDelaySeconds()
}

func (h *httpProcess) DescribePlayerSessions(w http.ResponseWriter, req *http.Request) {
	gameServerSessionId := req.URL.Query().Get("gameServerSessionId")
	playerId := req.URL.Query().Get("playerId")
	playerSessionId := req.URL.Query().Get("playerSessionId")
	playerSessionStatusFilter := req.URL.Query().Get("playerSessionStatusFilter")
	nextToken := req.URL.Query().Get("nextToken")
	limitStr := req.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitStr)

	gseManager := gsemanager.GetGseManager()
	resp, err := gseManager.DescribePlayerSessions(gameServerSessionId, playerId, playerSessionId, playerSessionStatusFilter,
		nextToken, int32(limit))

	logger.Logger.Infof("DescribePlayerSessions resp is ", zap.Any("resp", resp))

	if err != nil {
		code := int32(http.StatusInternalServerError)
		errMsg := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errMsg = st.Message()
			code = int32(st.Code())
		}
		resp, _ := h.writeResp(code, errMsg, nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	result, _ := h.writeResp(SUCCESS, SUCCESSMSG, resp)
	fmt.Fprintf(w, "%s", result)
}

func (h *httpProcess) UpdatePlayerSessionCreationPolicy(w http.ResponseWriter, req *http.Request) {
	newPolicy := req.URL.Query().Get("newPlayerSessionCreationPolicy")
	gameSessionId := req.URL.Query().Get("gameSessionId")

	gseManager := gsemanager.GetGseManager()
	_, err := gseManager.UpdatePlayerSessionCreationPolicy(newPolicy, gameSessionId)

	if err != nil {
		code := int32(http.StatusInternalServerError)
		errMsg := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errMsg = st.Message()
			code = int32(st.Code())
		}
		resp, _ := h.writeResp(code, errMsg, nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	successMsg, _ := h.writeResp(SUCCESS, SUCCESSMSG, nil)
	fmt.Fprintf(w, "%s", successMsg)
}

func (h *httpProcess) ReportCustomData(w http.ResponseWriter, req *http.Request) {
	currentCustomCountStr := req.URL.Query().Get("currentCustomCount")
	maxCustomCountStr := req.URL.Query().Get("maxCustomCount")

	currentCustomCount, errCurrent := (strconv.Atoi(currentCustomCountStr))
	maxCustomCount, errMax := (strconv.Atoi(maxCustomCountStr))

	if errCurrent != nil || errMax != nil {
		resp, _ := h.writeResp(http.StatusBadRequest, "currentCustomCount 或者 maxCustomCount必须是整数", nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	gseManager := gsemanager.GetGseManager()
	_, err := gseManager.ReportCustomData(int32(currentCustomCount), int32(maxCustomCount))

	if err != nil {
		code := int32(http.StatusInternalServerError)
		errMsg := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errMsg = st.Message()
			code = int32(st.Code())
		}
		resp, _ := h.writeResp(code, errMsg, nil)
		fmt.Fprintf(w, "%s", resp)
		return
	}

	successMsg, _ := h.writeResp(SUCCESS, SUCCESSMSG, nil)
	fmt.Fprintf(w, "%s", successMsg)
}

func (h *httpProcess) SetHealthStatus(w http.ResponseWriter, req *http.Request) {
	statusStr := req.URL.Query().Get("healthStatus")
	status, _ := (strconv.Atoi(statusStr))

	if status == 0 {
		rpcServerIns.healthStatus = false
	} else {
		rpcServerIns.healthStatus = true
	}

	successMsg, _ := h.writeResp(SUCCESS, SUCCESSMSG, nil)
	fmt.Fprintf(w, "%s", successMsg)
}

func (h *httpProcess) HelloWorld(w http.ResponseWriter, req *http.Request) {
	successMsg, _ := h.writeResp(SUCCESS, "hello,world", nil)
	fmt.Fprintf(w, "%s", successMsg)
}

func (h *httpProcess) GenerateHttpRandomPort(startPort int, endPort int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(endPort-startPort) + startPort
}

func (h *httpProcess) EndingProcessDelaySeconds() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(DelaySeconds * time.Second)
		logger.Logger.Infof("process terminated by http request, server port: %d", h.httpPort)
		os.Exit(1)
	}()
}
