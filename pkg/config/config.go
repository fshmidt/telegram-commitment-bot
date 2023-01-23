package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string
	Messages      Messages
	BoltDb        string `mapstructure:"db_file"`
}

type Messages struct {
	Responses Responses
	Errors    Errors
	Scales    Scales
	Parts     Parts
}

type Errors struct {
	Default        string `mapstructure:"default"`
	NoCommit       string `mapstructure:"noCommit"`
	NoDate         string `mapstructure:"noDate"`
	BadYear        string `mapstructure:"badYear"`
	BadMonth       string `mapstructure:"badMonth"`
	BadDay         string `mapstructure:"badDay"`
	ExistingCommit string `mapstructure:"existingCommit"`
	JsonError      string `mapstructure:"jsonError"`
	TooSoon        string `mapstructure:"tooSoon"`
	TooLongError   string `mapstructure:"tooLongError"`
}

type Responses struct {
	Start           string `mapstructure:"Start"`
	Commit          string `mapstructure:"Commit"`
	Unknown         string `mapstructure:"Unknown"`
	NoCommits       string `mapstructure:"NoCommits"`
	Nonsense        string `mapstructure:"Nonsense"`
	Ok1             string `mapstructure:"Ok1"`
	Ok2             string `mapstructure:"Ok2"`
	Ok3             string `mapstructure:"Ok3"`
	Ok4             string `mapstructure:"Ok4"`
	Done            string `mapstructure:"Done"`
	Congrats        string `mapstructure:"Congrats"`
	TooManyCommits  string `mapstructure:"TooManyCommits"`
	NoDoneCommits   string `mapstructure:"NoDoneCommits"`
	DeleteHeader    string `mapstructure:"DeleteHeader"`
	CommitIsDeleted string `mapstructure:"CommitIsDeleted"`
	Initiate        string `mapstructure:"Initiate"`
}

type Scales struct {
	Start      string `mapstructure:"start"`
	Ten        string `mapstructure:"ten"`
	Twenty     string `mapstructure:"twenty"`
	Thirty     string `mapstructure:"thirty"`
	Forty      string `mapstructure:"forty"`
	Fifty      string `mapstructure:"fifty"`
	Sixty      string `mapstructure:"sixty"`
	Seventy    string `mapstructure:"seventy"`
	Eighty     string `mapstructure:"eighty"`
	Ninety     string `mapstructure:"ninety"`
	NinetyFive string `mapstructure:"ninetyFive"`
	End        string `mapstructure:"end"`
}

type Parts struct {
	Hundredth  string `mapstructure:"hundredth"`
	Tenth      string `mapstructure:"tenth"`
	Fifth      string `mapstructure:"fifth"`
	Third      string `mapstructure:"third"`
	Half       string `mapstructure:"half"`
	SixtySix   string `mapstructure:"sixtySix"`
	SixtyNine  string `mapstructure:"sixtyNine"`
	Eighty     string `mapstructure:"eighty"`
	Ninety     string `mapstructure:"ninety"`
	NinetyFive string `mapstructure:"ninetyFive"`
	End        string `mapstructure:"end"`
}

func Init() (*Config, error) {

	//os.Setenv("TOKEN", "")

	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading env")
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.scales", &cfg.Messages.Scales); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.parts", &cfg.Messages.Parts); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("token")

	return nil
}
