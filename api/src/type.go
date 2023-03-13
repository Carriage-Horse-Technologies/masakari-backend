package main

import "github.com/gorilla/websocket"

type GetScoreResponse struct {
	Cost      int `json:"cost"`
	Score     int `json:"score"`
	UserCount int `json:"user_count"`
}

type PostScoreRequest struct {
	Score  float64 `json:"score"`
	UserId string  `json:"user_id"`
}

type PostScoreResponse struct {
	Message string `json:"message"`
}

type GetPodsCountResponse struct {
	Count int `json:"count"`
}
type connection struct {
	ws   *websocket.Conn
	send chan []byte
}
type subscription struct {
	conn *connection
	room string
}
type ByteBroadCast struct {
	Message []byte
	Type    int
	Conn    *websocket.Conn
}
type message struct {
	data []byte
	room string
}

type hub struct {
	rooms      map[string]map[*connection]bool
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}

type ErrorObject struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}
type Request struct {
	Action  string  `json:"action"`
	UserId  string  `json:"user_id"`
	Message string  `json:"message"`
	Name    string  `json:"name"`
	PosX    float32 `json:"pos_x"`
	PosY    float32 `json:"pos_y"`
}

type User struct {
	UserId    string  `json:"user_id"`
	UserImage string  `json:"url"`
	PosX      float32 `json:"pos_x"`
	PosY      float32 `json:"pos_y"`
}

type UserPosReponse struct {
	Action     string `json:"action"`
	Characters []User `json:"characters"`
}
type ProcessStatus struct {
	Name    string  `json:"name"`
	CPU     float32 `json:"cpu"`
	Memory  float32 `json:"memory"`
	Traffic int     `json:"traffic"`
}
type ProcessStatusResponse struct {
	Action string        `json:"action"`
	Status ProcessStatus `json:"status"`
}

type ChatResponse struct {
	Action  string `json:"action"`
	UserId  string `json:"user_id"`
	Message string `json:"message"`
	Name    string `json:"name"`
}
