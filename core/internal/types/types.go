// Code generated by goctl. DO NOT EDIT.
package types

type LonginRequest struct {
	Name     string `json:"name"`
	Password string `json:"Password"`
}

type LoginReply struct {
	Token string `json:"token"`
}
