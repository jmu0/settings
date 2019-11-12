# Settings
- Load: loads settings into struct or map
- environment variables overwrite settings in file
- command line args overwrite settings in file and environment variables

## Read settings from file
- .conf file (format: key value)
- .json file parse json
- .yml file parse yaml

## Read settings from Environment variables
<prg>_<key> = <value>
<prg>=executable name (all uppercase)
(underscore)
<key>setting key (all uppercase)

## Read settings from Command line args
<prg> --key=value

