package api

import (
	"fmt"
	"net/http"
	"strings"

	"git.dtluna.net/dtluna/syncthing-cli/config"
)

const (
	configPath = "/rest/system/config"

	getConfigError = "getting config: {{err}}"
	setConfigError = "setting config: {{err}}"
)

func Indent(s string, depth int) string {
	sep := strings.Repeat(" ", depth)
	return "\n" + sep + strings.Join(
		strings.Split(s, "\n"),
		"\n"+sep,
	)
}

func indentStringSlice(slice []string, depth int) string {
	if len(slice) == 0 {
		return "None"
	}

	sep := strings.Repeat(" ", depth)
	lines := []string{}

	for i, el := range slice {
		lines = append(lines, sep+fmt.Sprintf("%v. %s", i+1, el))
	}
	return "\n" + strings.Join(lines, "\n")
}

type FolderDevice struct {
	DeviceID string
}

func (fd FolderDevice) String() string {
	return fmt.Sprint("ID: ", fd.DeviceID)
}

func indentFolderDevices(fds []FolderDevice, depth int) string {
	if len(fds) == 0 {
		return "None"
	}

	sep := strings.Repeat(" ", depth)
	lines := []string{}

	for i, el := range fds {
		line := strings.ReplaceAll(fmt.Sprintf("%v. %s", i+1, el), "\n", "\n"+sep)
		lines = append(lines, line)
	}
	return "\n" + strings.Join(lines, "\n")
}

type VersioningInfoParams map[string]interface{}

func (vip VersioningInfoParams) String() string {
	if len(vip) == 0 {
		return "None"
	}

	slice := []string{}
	for key, value := range vip {
		slice = append(slice, fmt.Sprintf("%v: %v", key, value))
	}
	return strings.Join(slice, "\n")
}

type VersioningInfo struct {
	Type   string
	Params VersioningInfoParams
}

func (vi VersioningInfo) String() string {
	return fmt.Sprintf(
		`Type: %v
Params: %v`,
		vi.Type,
		Indent(vi.Params.String(), 2),
	)
}

type Folder struct {
	ID      string
	Label   string
	Path    string
	Type    string
	Devices []FolderDevice

	RescanIntervalS  int
	FSWatcherEnabled bool
	Order            string
	IgnorePerms      bool
	IgnoreDelete     bool
	MinDiskFreePct   int
	MinDiskFree      string
	Versioning       VersioningInfo
	AutoNormalize    bool

	Copiers int
	Pullers int
	Hashers int

	ScanProgressIntervalS int
	PullerSleepS          int
	PullerPauseS          int
	MaxConflicts          int
	DisableSparseFiles    bool
	DisableTempIndexes    bool
	Fsync                 bool
	Invalid               string
}

func (f Folder) String() string {
	var viString string
	if f.Versioning.Type == "" {
		viString = "None"
	} else {
		viString = Indent(f.Versioning.String(), 2)
	}

	return fmt.Sprintf(
		`ID: %v
Label: %v
Path: %v
Type: %v
Devices: %v
Versioning: %v`,
		f.ID,
		f.Label,
		f.Path,
		f.Type,
		indentFolderDevices(f.Devices, 2),
		viString,
	)
}

type Device struct {
	DeviceID    string
	Name        string
	Addresses   []string
	Compression string
	CertName    string
	Introducer  bool
}

func (d Device) String() string {
	return fmt.Sprintf(
		`ID: %v
Name: %v
Addresses: %v
Compression: %v
Certificate name: %v
Introducer: %v`,
		d.DeviceID,
		d.Name,
		indentStringSlice(d.Addresses, 2),
		d.Compression,
		d.CertName,
		d.Introducer,
	)
}

type GUIConfig struct {
	Enabled             bool
	Address             string
	User                string
	UseTLS              bool
	InsecureAdminAccess bool
	Theme               string
	APIKey              string
}

func (gc GUIConfig) String() string {
	return fmt.Sprintf(
		`Enabled: %v
Address: %v
User: %v
TLS used: %v
Insecure admin access: %v
Theme: %v`,
		gc.Enabled,
		gc.Address,
		gc.User,
		gc.UseTLS,
		gc.InsecureAdminAccess,
		gc.Theme,
	)
}

type Options struct {
	ListenAddresses                     []string
	GlobalAnnounceServers               []string
	GlobalAnnounceEnabled               bool
	LocalAnnounceEnabled                bool
	LocalAnnouncePort                   int
	LocalAnnounceMCAddr                 string
	MaxSendKbps                         int
	MaxRecvKbps                         int
	ReconnectionIntervalS               int
	RelaysEnabled                       bool
	RelayReconnectIntervalM             int
	StartBrowser                        bool
	NatEnabled                          bool
	NatLeaseMinutes                     int
	NatRenewalMinutes                   int
	NatTimeoutSeconds                   int
	UrAccepted                          int
	UrUniqueID                          string
	UrURL                               string
	UrPostInsecurely                    bool
	UrInitialDelayS                     int
	RestartOnWakeup                     bool
	AutoUpgradeIntervalH                int
	KeepTemporariesH                    int
	CacheIgnoredFiles                   bool
	ProgressUpdateIntervalS             int
	LimitBandwidthInLan                 bool
	MinHomeDiskFreePct                  int
	ReleasesURL                         string
	AlwaysLocalNets                     []string
	OverwriteRemoteDeviceNamesOnConnect bool
	TempIndexMinBlocks                  int
}

func (o Options) String() string {
	return fmt.Sprintf(
		`Listen addresses: %v
Global announce servers: %v
Global announce enabled: %v
Local announce enabled: %v
Local announce port: %v
Local announce MAC address: %v
Max sending speed: %v kbit/s
Max receiving speed: %v kbit/s
Reconnection interval: %v s
Relays enabled: %v
Relay reconnect interval: %v min
Start browser: %v
NAT enabled: %v
NAT lease: %v min
NAT renewal: %v min
NAT timeout: %v s
Usage reports accepted: %v
Usage report unique ID: %v
Usage report URL: %v
Post usage reports insecurely: %v
Usage report initial delay: %v s
Restart on wakeup: %v
Auto upgrade interval: %v h
Keep temporary failed transfers: %v h
Cache ignored files: %v
Progress update interval: %v s
Limit bandwidth in LAN: %v
Minimal home disk free space: %v%%
Releases URL: %v
Always local networks: %v
Overwrite remote device names on connect: %v
`,
		indentStringSlice(o.ListenAddresses, 2),
		indentStringSlice(o.GlobalAnnounceServers, 2),
		o.GlobalAnnounceEnabled,
		o.LocalAnnounceEnabled,
		o.LocalAnnouncePort,
		o.LocalAnnounceMCAddr,
		o.MaxSendKbps,
		o.MaxRecvKbps,
		o.ReconnectionIntervalS,
		o.RelaysEnabled,
		o.RelayReconnectIntervalM,
		o.StartBrowser,
		o.NatEnabled,
		o.NatLeaseMinutes,
		o.NatRenewalMinutes,
		o.NatTimeoutSeconds,
		o.UrAccepted,
		o.UrUniqueID,
		o.UrURL,
		o.UrPostInsecurely,
		o.UrInitialDelayS,
		o.RestartOnWakeup,
		o.AutoUpgradeIntervalH,
		o.KeepTemporariesH,
		o.CacheIgnoredFiles,
		o.ProgressUpdateIntervalS,
		o.LimitBandwidthInLan,
		o.MinHomeDiskFreePct,
		o.ReleasesURL,
		indentStringSlice(o.AlwaysLocalNets, 2),
		o.OverwriteRemoteDeviceNamesOnConnect,
	)
}

type STConfig struct {
	Version        int
	Folders        []Folder
	Devices        []Device
	GUI            GUIConfig
	Options        Options
	IgnoredDevices []string
	IgnoredFolders []string
}

func (c STConfig) String() string {
	return fmt.Sprintf(
		`Version: %v

GUI: %v

Options: %v

Ignored devices: %v

Ignored folders: %v`,
		c.Version,
		Indent(c.GUI.String(), 2),
		Indent(c.Options.String(), 2),
		indentStringSlice(c.IgnoredDevices, 2),
		indentStringSlice(c.IgnoredFolders, 2),
	)
}

func GetConfig(cfg *config.Config) (*STConfig, error) {
	req := newClient(cfg).Request().Path(configPath)
	resp, err := req.Send()
	if err != nil {
		return nil, wrapError(err, getConfigError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, wrapError(err, getConfigError)
	}

	stc := new(STConfig)
	err = resp.JSON(stc)
	if err != nil {
		return nil, wrapError(err, getConfigError)
	}

	return stc, nil
}

func SetConfig(cfg *config.Config, stconfig *STConfig) (*STConfig, error) {
	req := newClient(cfg).Request().Path(configPath).
		Method(http.MethodPost).JSON(stconfig)
	resp, err := req.Send()
	if err != nil {
		return nil, wrapError(err, setConfigError)
	}

	err = checkResponseOK(resp)
	if err != nil {
		return nil, wrapError(err, setConfigError)
	}

	stc := new(STConfig)
	err = resp.JSON(stc)
	if err != nil {
		return nil, wrapError(err, setConfigError)
	}

	return stc, nil
}
