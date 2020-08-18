go 1.14

module github.com/jinlongchen/viper

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/hashicorp/hcl v1.0.0
	github.com/jinlongchen/crypt v0.0.0-20200818165202-ff1ea6e6083e
	github.com/magiconair/properties v1.8.1
	github.com/mitchellh/mapstructure v1.3.3
	github.com/pelletier/go-toml v1.8.0
	github.com/spf13/afero v1.3.4
	github.com/spf13/cast v1.3.1
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/subosito/gotenv v1.2.0
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
	github.com/jinlongchen/crypt v0.0.0-20200818165202-ff1ea6e6083e => ../crypt
)
