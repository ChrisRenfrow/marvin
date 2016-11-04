package marvin

import (
	"fmt"
	"net/http"
	"net/url"

	"gopkg.in/ini.v1"

	"github.com/riking/homeapi/marvin/database"
	"github.com/riking/homeapi/marvin/slack"
)

// ModuleID is a string constant identifying a module.
type ModuleID string

type ModuleState int

const (
	_ ModuleState = iota
	ModuleStateConstructed
	ModuleStateLoaded
	ModuleStateEnabled
	ModuleStateDisabled
	ModuleStateErrorLoading
	ModuleStateErrorEnabling
)

type Module interface {
	Identifier() ModuleID

	// Load should declare dependencies.
	Load(t Team)
	// Enable has dependencies available.
	Enable(t Team)
	// Disable should shut down and unregister all resources.
	Disable(t Team)
}

type ModuleStatus interface {
	Instance() Module
	State() ModuleState
	// Returns non-nil if Degraded returns true.
	Err() error

	IsLoaded() bool
	IsEnabled() bool
	Degraded() bool
}

type ModuleConfig interface {
	// Get gets a module configuration value. The error will be set on database errors.
	// Get() will panic if the key was not initialized with Add() or AddProtect().
	Get(key string) (string, error)
	// GetIsDefault gets a module configuration value, but does not require the key
	// have been initialized.
	//
	// 1) If the key was not initialized with Add(), value is the empty string,
	//    isDefault is true, and err is ErrConfNoDefault.
	// 2) If the key was initialized, but has no override, value is the default value,
	//    isDefault is true, and err is nil.
	// 3) If the key has an override, value is the override, isDefault is false, and
	//    err is nil.
	GetIsDefault(key string) (value string, isDefault bool, err error)
	// GetIsDefaultNotProtected acts like GetIsDefault, but returns
	// ("", false, ErrConfProtected) if the key is protected.
	GetIsDefaultNotProtected(key string) (value string, isDefault bool, err error)
	// Set sets an override for the given configuration key.
	Set(key, value string) error
	// SetDefault resets the configuration for the given key to the default.
	SetDefault(key string) error
	// Add initializes the default value for a key for use with Get().
	// This must be called during the module Load phase.
	Add(key, defaultValue string)
	// Add initializes the default value for a key for use with Get(), and also sets
	// the key as protected.
	// This must be called during the module Load phase.
	AddProtect(key, defaultValue string, protect bool)

	// ListDefaults returns the defaults map.
	// This cannot be called during the module Load phase.
	ListDefaults() map[string]string
	// ListDefaults returns the protected-keys map.
	// This cannot be called during the module Load phase.
	ListProtected() map[string]bool
}

// ErrConfProtected is an error return from ModuleConfig.GetIsDefaultNotProtected.
type ErrConfProtected struct {
	Key string
}

// Error implements the error interface.
func (e ErrConfProtected) Error() string {
	return fmt.Sprintf("%s is a protected configuration value. Viewing is restricted to admin DMs.", e.Key)
}

// ErrConfNoDefault is an error return from ModuleConfig.GetIsDefault.
type ErrConfNoDefault struct {
	Key string
}

// Error implements the error interface.
func (e ErrConfNoDefault) Error() string {
	return fmt.Sprintf("%s had no default set.", e.Key)
}

// AccessLevel represents the level of rights a user has.
type AccessLevel int

const (
	AccessLevelBlacklisted AccessLevel = iota
	AccessLevelNormal
	AccessLevelChannelAdmin
	AccessLevelAdmin
	AccessLevelController
)

// ActionSource represents the cause of actions or commands.
type ActionSource interface {
	UserID() slack.UserID
	ChannelID() slack.ChannelID
	AccessLevel() AccessLevel
	ArchiveLink() string

	SendCmdReply(result CommandResult) CommandResult
}

// TeamConfig is loaded from the config.ini file.
type TeamConfig struct {
	TeamDomain   string
	ClientID     string
	ClientSecret string
	VerifyToken  string
	DatabaseURL  string
	UserToken    string
	LogChannel   slack.ChannelID
}

func LoadTeamConfig(sec *ini.Section) *TeamConfig {
	c := &TeamConfig{}
	c.TeamDomain = sec.Key("TeamDomain").String()
	c.ClientID = sec.Key("ClientID").String()
	c.ClientSecret = sec.Key("ClientSecret").String()
	c.VerifyToken = sec.Key("VerifyToken").String()
	c.DatabaseURL = sec.Key("DatabaseURL").String()
	c.UserToken = sec.Key("UserToken").String()
	c.LogChannel = slack.ChannelID(sec.Key("LogChannel").String())
	return c
}

type SubCommand interface {
	Handle(t Team, args *CommandArguments) CommandResult
	Help(t Team, args *CommandArguments) CommandResult
}

type SubCommandFunc func(t Team, args *CommandArguments) CommandResult

type CommandRegistration interface {
	RegisterCommand(name string, c SubCommand)
	RegisterCommandFunc(name string, c SubCommandFunc, help string) SubCommand
	UnregisterCommand(name string)
}

type HTTPDoer interface {
	Do(*http.Request) (http.Response, error)
}

type SendMessage interface {
	SendMessage(channelID slack.ChannelID, message string) (slack.MessageTS, slack.RTMRawMessage, error)
	SendComplexMessage(channelID slack.ChannelID, message url.Values) (slack.MessageTS, error)
}

// Team represents a Slack team, and is the "god object" for Marvin.
//
// Its implementation is in the marvin/slack/controller package.
type Team interface {
	// Domain returns the leftmost component of the Slack domain name.
	Domain() string
	DB() *database.Conn
	TeamConfig() *TeamConfig
	ModuleConfig(mod ModuleID) ModuleConfig

	// BotUser returns the user ID that Marvin is signed in as.
	BotUser() slack.UserID

	// EnableModules loads every module and attempts to transition them to
	// the state listed in the configuration.
	EnableModules()

	// DependModule places the instance of the requested module in the
	// given pointer.
	//
	// If the requested module is already enabled, the pointer is
	// filled immediately and the function returns 1. If the requested
	// module has errored, the pointer is left alone and the function
	// returns -2.
	//
	// During loading, when the requested module has not been enabled
	// yet, the function returns 0 and remembers the pointer. If the
	// requested module is not known, the function returns -1.
	DependModule(self Module, dependencyID ModuleID, ptr *Module) int
	// GetModule returns the Module instance for a module directly.
	// TODO - return its state?
	GetModule(modID ModuleID) Module
	// GetAllModules() returns the status of all modules.
	GetAllModules() []ModuleStatus
	// GetAllModules() returns the status of all enabled modules.
	GetAllEnabledModules() []ModuleStatus

	SendMessage
	ReactMessage(msgID slack.MessageID, emojiName string) error
	// SlackAPIPost makes a Slack API call by adding the token to the form.
	// TODO could do more for the caller, like unmarshalling
	SlackAPIPost(method string, form url.Values) (*http.Response, error)

	ArchiveURL(msgID slack.MessageID) string

	OnEveryEvent(mod ModuleID, f func(slack.RTMRawMessage))
	OnEvent(mod ModuleID, event string, f func(slack.RTMRawMessage))
	OnNormalMessage(mod ModuleID, f func(slack.RTMRawMessage))
	OffAllEvents(mod ModuleID)

	CommandRegistration
	DispatchCommand(args *CommandArguments) CommandResult

	ReportError(err error, source ActionSource)

	ChannelName(channel slack.ChannelID) string
	UserName(user slack.UserID) string
	UserLevel(user slack.UserID) AccessLevel
	GetIM(user slack.UserID) (slack.ChannelID, error)
	PublicChannelInfo(channel slack.ChannelID) (*slack.Channel, error)
	PrivateChannelInfo(channel slack.ChannelID) (*slack.Channel, error)
}

type ShockyInstance interface {
	TeamConfig(teamDomain string) TeamConfig
	ModuleConfig(team TeamConfig) ModuleConfig
	DB(team TeamConfig) *database.Conn

	SendChannelSlack(team Team, channel string, message slack.OutgoingSlackMessage)
	SendPrivateSlack(team Team, user string, message slack.OutgoingSlackMessage)
}