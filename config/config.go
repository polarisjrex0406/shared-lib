package config

import (
	"fmt"
	"sync"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

var (
	instance *ConfigType
	once     sync.Once
)

// LoadConfig loads the configuration from .env and command-line flags.
func LoadConfig() (*ConfigType, error) {
	var cfg ConfigType
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	fp := flags.NewParser(&cfg, flags.Default)
	// Parse flags
	if _, err := fp.Parse(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetConfig returns the singleton instance of Config.
func GetConfig() (*ConfigType, error) {
	var err error
	once.Do(func() {
		instance, err = LoadConfig()
	})
	return instance, err
}

type ConfigType struct {
	Command struct {
		Migrate bool `long:"migrate"`
		Seed    bool `long:"seed"`
	}

	Server struct {
		Port int `long:"server-port" env:"SERVER_PORT" default:"8080"`
	}

	JWT struct {
		PrivateKey string `long:"jwt-private-key" env:"JWT_PRIVATE_KEY"`
		Duration   struct {
			AccessTokenInMin   int `long:"jwt-duration-access-token-in-min" env:"JWT_DURATION_ACCESS_TOKEN_IN_MIN" default:"15"`
			RefreshTokenInHour int `long:"jwt-duration-refresh-token-in-hour" env:"JWT_DURATION_REFRESH_TOKEN_IN_HOUR" default:"24"`
			KeepLogInDay       int `long:"jwt-duration-keep-log-in-day" env:"JWT_DURATION_KEEP_LOG_IN_DAY" default:"30"`
		}
	}

	Recaptcha struct {
		URL       string `long:"recaptcha-url" env:"RECAPTCHA_URL" default:"https://www.google.com/recaptcha/api/siteverify"`
		SecretKey string `long:"recaptcha-secret-key" env:"RECAPTCHA_SECRET_KEY" default:"YOUR SECRET KEY"`
		SiteKey   string `long:"recaptcha-site-key" env:"RECAPTCHA_SITE_KEY" default:"YOUR SITE KEY"`
	}

	Postgres struct {
		DBName   string `long:"postgres-db-name" env:"POSTGRES_DB_NAME" default:"mimicproxy"`
		User     string `long:"postgres-user" env:"POSTGRES_USER" default:"postgres"`
		Password string `long:"postgres-password" env:"POSTGRES_PASSWORD"`
		Host     string `long:"postgres-host" env:"POSTGRES_HOST" default:"localhost"`
		Port     int    `long:"postgres-port" env:"POSTGRES_PORT" default:"5432"`
		SSLMode  string `long:"postgres-ssl-mode" env:"POSTGRES_SSL_MODE" default:"disable"`
	}

	InfluxDB struct {
		Host         string `long:"influxdb-host" env:"INFLUXDB_HOST" default:"localhost"`
		Port         int    `long:"influxdb-port" env:"INFLUXDB_PORT" default:"8086"`
		Token        string `long:"influxdb-token" env:"INFLUXDB_TOKEN"`
		Organization string `long:"influxdb-org" env:"INFLUXDB_ORG" default:"omimic"`
		Bucket       string `long:"influxdb-bucket" env:"INFLUXDB_BUCKET" default:"mimicproxy"`
	}

	Email struct {
		MailTrap struct {
			URL    string `long:"email-mailtrap-url" env:"EMAIL_MAILTRAP_URL"`
			APIKey string `long:"email-mailtrap-api-key" env:"EMAIL_MAILTRAP_API_KEY"`
		}
		TemplatePath string `long:"email-template-path" env:"EMAIL_TEMPLATE_PATH"`
	}

	Frontend struct {
		BaseURL string `long:"frontend-base-url" env:"FRONTEND_BASE_URL" default:"frontend_url"`
	}

	FileStorage struct {
		Local struct {
			BasePath string `long:"file-storage-local-base-path" env:"FILE_STORAGE_LOCAL_BASE_PATH" default:"base_path"`
		}
	}

	AntiVirus struct {
		ClamAV struct {
			BaseURL         string `long:"anti-virus-clamav-base-url" env:"ANTI_VIRUS_CLAMAV_BASE_URL" default:"http://localhost:8081"`
			ScanEndpoint    string `long:"anti-virus-clamav-scan-endpoint" env:"ANTI_VIRUS_CLAMAV_SCAN_ENDPOINT" default:"/api/v1/scan"`
			VersionEndpoint string `long:"anti-virus-clamav-version-endpoint" env:"ANTI_VIRUS_CLAMAV_VERSION_ENDPOINT" default:"/api/v1/version"`
			DBSignEndpoint  string `long:"anti-virus-clamav-db-sign-endpoint" env:"ANTI_VIRUS_CLAMAV_DB_SIGN_ENDPOINT" default:"/api/v1/dbsignatures"`
		}
	}

	Payment struct {
		Cryptomus struct {
			APIKey     string `long:"payment-cryptomus-api-key" env:"PAYMENT_CRYPTOMUS_API_KEY"`
			MerchantID string `long:"payment-cryptomus-merchant-id" env:"PAYMENT_CRYPTOMUS_MERCHANT_ID"`
			PaymentURL string `long:"payment-cryptomus-payment-url" env:"PAYMENT_CRYPTOMUS_PAYMENT_URL" default:"https://api.cryptomus.com/v1/payment"`
			InfoURL    string `long:"payment-cryptomus-info-url" env:"PAYMENT_CRYPTOMUS_INFO_URL" default:"https://api.cryptomus.com/v1/payment/info"`
			Webhook    struct {
				URL       string `long:"payment-cryptomus-webhook-url" env:"PAYMENT_CRYPTOMUS_WEBHOOK_URL" default:"/api/balance/cryptomus/webhook"`
				AllowedIP string `long:"payment-cryptomus-webhook-allowed-ip" env:"PAYMENT_CRYPTOMUS_WEBHOOK_ALLOWED_IP" default:"91.227.144.54"`
			}
		}
	}

	Provider struct {
		TTProxy struct {
			BaseURL          string `long:"provider-ttp-base-url" env:"PROVIDER_TTP_BASEURL" default:"https://api.ttproxy.com/v1/subLicense/"`
			License          string `long:"provider-ttp-license" env:"PROVIDER_TTP_LICENSE"`
			Secret           string `long:"provider-ttp-secret" env:"PROVIDER_TTP_SECRET"`
			ProxyCredentials struct {
				Host string `long:"provider-ttp-proxy-cred-host" env:"PROVIDER_TTP_PROXY_CRED_HOST" default:"dynamic.ttproxy.com"`
				Port int    `long:"provider-ttp-proxy-cred-port" env:"PROVIDER_TTP_PROXY_CRED_PORT" default:"10001"`
			}
		}
		DataImpulse struct {
			BaseURL          string `long:"provider-di-base-url" env:"PROVIDER_DI_BASEURL" default:"https://api.dataimpulse.com/provider/"`
			Login            string `long:"provider-di-login" env:"PROVIDER_DI_LOGIN"`
			Password         string `long:"provider-di-password" env:"PROVIDER_DI_PASSWORD"`
			ProxyCredentials struct {
				Host string `long:"provider-di-proxy-cred-host" env:"PROVIDER_DI_PROXY_CRED_HOST" default:"gw.dataimpulse.com"`
				Port int    `long:"provider-di-proxy-cred-port" env:"PROVIDER_DI_PROXY_CRED_PORT" default:"823"`
			}
		}
		Proxyverse struct {
			ProxyCredentials struct {
				Host     string `long:"provider-pv-proxy-cred-host" env:"PROVIDER_PV_PROXY_CRED_HOST" default:"51.81.93.42"`
				Port     int    `long:"provider-pv-proxy-cred-port" env:"PROVIDER_PV_PROXY_CRED_PORT" default:"9200"`
				Username string `long:"provider-pv-proxy-cred-username" env:"PROVIDER_PV_PROXY_CRED_USERNAME"`
				Password string `long:"provider-pv-proxy-cred-password" env:"PROVIDER_PV_PROXY_CRED_PASSWORD"`
			}
		}
		Databay struct {
			ProxyCredentials struct {
				Host     string `long:"provider-db-proxy-cred-host" env:"PROVIDER_DB_PROXY_CRED_HOST" default:"resi-global-gateways.databay.com"`
				Port     int    `long:"provider-db-proxy-cred-port" env:"PROVIDER_DB_PROXY_CRED_PORT" default:"7676"`
				Username string `long:"provider-db-proxy-cred-username" env:"PROVIDER_DB_PROXY_CRED_USERNAME"`
				Password string `long:"provider-db-proxy-cred-password" env:"PROVIDER_DB_PROXY_CRED_PASSWORD"`
			}
		}
	}
}
