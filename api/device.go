package api

// The Device struct describes an available playback device
type Device struct {
	ID            string `json:"id"`
	IsActive      bool   `json:"is_active"`
	IsRestricted  bool   `json:"is_restricted"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	VolumePercent int    `json:"volume_percent"`
}

// The Devices struct is needed because the /devices endpoint returns the devices wrapped in a root object
type Devices struct {
	Devices []Device `json:"devices"`
}
