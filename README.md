# Settings
Load() loads settings from file (.conf, .json or .yml), overrides these settings with environment variables and then overrides these settings with command line arguments.
## Read settings from file
- .conf file (format: key value)
- .json file parse json
- .yml file parse yaml

## Read settings from Environment variables
Format: executable name in all caps, underscore, key in all caps.
    export APP_KEY=value

## Read settings from Command line args
    `<app>` --key=value

# Example
    var s map[string]string

    s = make(map[string]string)
	settings.Load("settings.conf", &s)
