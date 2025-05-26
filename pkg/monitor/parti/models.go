package parti

import "time"

type ApiResponse struct {
	Status      string      `json:"status"`
	IsModerator interface{} `json:"is_moderator"`
	ChannelInfo struct {
		Channel struct {
			LivestreamID      int    `json:"livestream_id"`
			UserID            int    `json:"user_id"`
			ChannelArn        string `json:"channel_arn"`
			IngestEndpoint    string `json:"ingest_endpoint"`
			PlaybackURL       string `json:"playback_url"`
			PlaybackAuthToken string `json:"playback_auth_token"`
			StreamKeyArn      string `json:"stream_key_arn"`
			StreamKey         string `json:"stream_key"`
			CreatedAt         int    `json:"created_at"`
			IsLive            bool   `json:"is_live"`
			CurrentEventID    int    `json:"current_event_id"`
		} `json:"channel"`
		Stream struct {
			Health      string    `json:"health"`
			StartTime   time.Time `json:"start_time"`
			State       string    `json:"state"`
			StreamID    string    `json:"stream_id"`
			ViewerCount int       `json:"viewer_count"`
		} `json:"stream,omitempty"`
		VisibilityType      string `json:"visibility_type"`
		LivestreamEventInfo struct {
			LivestreamPreviewFile string `json:"livestream_preview_file"`
			EventName             string `json:"event_name"`
			EventDescription      string `json:"event_description"`
		} `json:"livestream_event_info,omitempty"`
	} `json:"channel_info"`
	IsStreamingLiveNow bool `json:"is_streaming_live_now"`
}
