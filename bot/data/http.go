package data

type Session struct {
	New        bool
	SessionId  string
	Attributes map[string]string
}

type System struct {
	User        User
	Application Application
	Device      Device
}

type User struct {
	UserId      string
	AccessToken string
	UserInfo    UserInfo
}

type UserInfo struct {
}

type Application struct {
	ApplicationId string
}

type Device struct {
	DeviceId            string
	SupportedInterfaces map[string]interface{}
}

type Context struct {
	System      System
	AudioPlayer AudioPlayerContext
	VideoPlayer VideoPlayerContext
}

type AudioPlayerContext struct {
	Token                string
	OffsetInMilliseconds int32 `json:"offsetInMilliSeconds,omitempty"`
	PlayActivity         string
}

type VideoPlayerContext struct {
	Token                string
	OffsetInMilliseconds int32 `json:"offsetInMilliseconds,omitempty"`
	PlayActivity         string
}

type baseRequest struct {
	Type      string
	RequestId string
	Timestamp string
}

// 公共请求体
type RequestPart struct {
	Version string
	Session Session
	Context Context
	Request baseRequest
}

// 事件请求
type EventRequest struct {
	Request struct {
		baseRequest
		Token string
		Url   string `json:"url,omitempty"`
		Name  string `json:"name,omitempty"`
	}
}

type AudioPlayerEventRequest struct {
	Request struct {
		baseRequest
		Token                string
		OffsetInMilliseconds int32 `json:"offsetInMilliSeconds,omitempty"` //Audio Player Event
	}
}

type VideoPlayerEventRequest struct {
	Request struct {
		baseRequest
		Token                string
		OffsetInMilliseconds int32 `json:"offsetInMilliseconds,omitempty"` //Audio Player Event
	}
}

// 打开请求
type LaunchRequest struct {
	Request baseRequest
}

// 退出请求
type SessionEndedRequestBody struct {
	baseRequest
	Reason string
}

type SessionEndedRequest struct {
	Request SessionEndedRequestBody
}

// intent request
type Query struct {
	Type     string
	Original string
}

type Slot struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
}

type Intent struct {
	Name               string          `json:"name"`
	Slots              map[string]Slot `json:"slots"`
	ConfirmationStatus string          `json:"confirmationStatus"`
}

type IntentRequestBody struct {
	baseRequest
	Query       Query
	DialogState string
	Intents     []Intent
}

type IntentRequest struct {
	Request IntentRequestBody
}

//--------response------
type SessionResponse struct {
	Attributes map[string]string `json:"attributes"`
}

type ContextResponse struct {
	Intent Intent `json:"intent"`
}

type DialogDirective struct {
	Type          string `json:"type"`
	SlotToElicit  string `json:"slotToElicit,omitempty"`
	SlotToConfirm string `json:"slotToConfirm,omitempty"`
	UpdatedIntent Intent `json:"updatedIntent"`
}

type Speech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Ssml string `json:"ssml,omitempty"`
}
