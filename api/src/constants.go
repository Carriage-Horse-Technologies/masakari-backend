package main

var ErrorSocketResponse = []byte(`{"action":"ERROR_MESSAGE","status":"NG","error": true}`)
var SampleResponse = []byte(`{"action":"UPDATE_CHARACTER_POS_EXAMPLE","characters":[{"pos_x":114.514,"pos_y":114.514,"url":"https://example.cloudfront.net/image/character_1.png","user_id":"example-user-id"},{"pos_x":810.514,"pos_y":810.514,"url":"https://example.cloudfront.net/image/character_1.png","user_id":"example-user-id"}]}`)
var SystemResponse = []byte(`{"action":"SYSTEM_MESSAGE","status":"OK","error": false}`)

const ID_PATH = "ID_PATH"
const CONNECTION_PATH = "connection_path"
const BASE_POS_PATH = "POS/"
const ROOM_ID = "last-hack"
