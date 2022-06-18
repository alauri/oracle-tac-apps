package cfg

import "github.com/spf13/viper"

type driver struct {
    Username string `toml:"username"`
    Password string `toml:"password"`
}

type database struct {
    Table string `toml:"table"`
}

type config struct {
    Driver driver `toml:"driver"`
    Database database `toml:"database"`
}


func ReadTOML(path string) config {
    // Read configuration TOML file
    var vp = viper.New()
    var ctoml config

    vp.SetConfigName("config")
    vp.SetConfigType("toml")
    vp.AddConfigPath(path)

    // Read the TOML file
    err := vp.ReadInConfig()
    if err != nil {
        panic(err)
    }

    // Convert the TOML file into struct
    err = vp.Unmarshal(&ctoml)
    if err != nil {
        panic(err)
    }
    return ctoml
}

func UpdateTOML(path string, label string, value string) {
    var vp = viper.New()

    vp.SetConfigName("config")
    vp.SetConfigType("toml")
    vp.AddConfigPath(path)

    vp.Set(label, value)
    vp.WriteConfig()
}
