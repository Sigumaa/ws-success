package main

import (
	"github.com/Sigumaa/ws-success/model/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ChangeVolume(c echo.Context) error {
	res := domain.WebSocketResponse[domain.ChangeVolume]{
		Type: "CHANGE_VOLUME",
		Payload: domain.ChangeVolume{
			Volume: 50,
		},
	}
	JSON <- res
	return c.String(http.StatusOK, "ChangeVolumeが呼び出されました")
}

func MemoUpdate(c echo.Context) error {
	res := domain.WebSocketResponse[domain.MemoUpdate]{
		Type: "MEMO_UPDATE",
		Payload: domain.MemoUpdate{
			Contents: "Hello",
		},
	}
	JSON <- res
	return c.String(http.StatusOK, "MemoUpdateが呼び出されました")
}

func SpeakerChange(c echo.Context) error {
	res := domain.WebSocketResponse[domain.SpeakerChange]{
		Type: "SPEAKER_CHANGE",
		Payload: domain.SpeakerChange{
			SpeakerQuotaTypeID: "1",
		},
	}
	JSON <- res
	return c.String(http.StatusOK, "SpeakerChangeが呼び出されました")
}

func SceneChange(c echo.Context) error {
	res := domain.WebSocketResponse[domain.SceneChange]{
		Type: "SCENE_CHANGE",
		Payload: domain.SceneChange{
			SceneType: "1",
			Volume:    50,
		},
	}
	JSON <- res
	return c.String(http.StatusOK, "SceneChangeが呼び出されました")
}

func MessageUpdate(c echo.Context) error {
	res := domain.WebSocketResponse[domain.MessageUpdate]{
		Type: "MESSAGE_UPDATE",
		Payload: domain.MessageUpdate{
			Message: []string{"Hello", "World"},
		},
	}
	JSON <- res
	return c.String(http.StatusOK, "MessageUpdateが呼び出されました")
}

func Comments(c echo.Context) error {
	res := domain.WebSocketResponse[domain.Comments]{
		Type: "COMMENTS",
		Payload: domain.Comments{
			Platform: "Twitter",
			Name:     "400",
			IconURL:  "400x400.jpg",
			Content:  "Hello",
		},
	}
	JSON <- res
	return c.String(http.StatusOK, "Commentsが呼び出されました")
}
