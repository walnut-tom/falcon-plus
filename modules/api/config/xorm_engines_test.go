package config

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("debug", true)
	viper.SetDefault("gen_sql", true)
	viper.SetDefault("db.alarms", "postgres://postgres:postgres@127.0.0.1/alarms?sslmode=disable")
	viper.SetDefault("db.dashboard", "postgres://postgres:postgres@127.0.0.1/dashboard?sslmode=disable")
	viper.SetDefault("db.falcon_portal", "postgres://postgres:postgres@127.0.0.1/falcon_portal?sslmode=disable")
	viper.SetDefault("db.graph", "postgres://postgres:postgres@127.0.0.1/graph?sslmode=disable")
	viper.SetDefault("db.uic", "postgres://postgres:postgres@127.0.0.1/uic?sslmode=disable")
	InitEngine(true, viper.GetViper())
}

func TestEngines(t *testing.T) {
	tests := []struct {
		name string
		want Engine
	}{
		{
			name: "Engine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Engines(); got == nil {
				t.Errorf("Engines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitEngine(t *testing.T) {
	type args struct {
		loggerlevel bool
		vip         *viper.Viper
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "db.alarm",
			args: args{
				loggerlevel: true,
				vip:         viper.GetViper(),
			},
		}, {
			name: "db.dashboard",
			args: args{
				loggerlevel: true,
				vip:         viper.GetViper(),
			},
		}, {
			name: "db.falcon_portal",
			args: args{
				loggerlevel: true,
				vip:         viper.GetViper(),
			},
		}, {
			name: "db.graph",
			args: args{
				loggerlevel: true,
				vip:         viper.GetViper(),
			},
		}, {
			name: "db.uic",
			args: args{
				loggerlevel: true,
				vip:         viper.GetViper(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitEngine(tt.args.loggerlevel, tt.args.vip); (err != nil) != tt.wantErr {
				t.Errorf("InitEngine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
