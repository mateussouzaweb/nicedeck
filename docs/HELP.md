Usage: nicedeck [COMMAND] [OPTIONS]

COMMAND:

version         print version information
help            print application help
programs        list available programs
platforms       list available platforms
shortcuts       list current user shortcuts
scrape          scrape data on SteamGridDB
launch          launch shortcut with given ID
modify          update or delete shortcut with given ID
install         install or update programs
remove          remove previously installed programs
backup-state    backup emulators state
restore-state   restore emulators state
process-roms    process emulators ROMs
server          start server for GUI usage (default)

OPTIONS:

scrape:
  --search=[term]             search term

launch:
  --id=[ID]                   shortcut ID

modify:
  --update                    update shortcut
  --delete                    delete shortcut
  --id=[ID]                   shortcut ID
  --program=[value]           program for shortcut
  --name=[value]              name for shortcut
  --description=[value]       description for shortcut
  --start-directory=[value]   start directory for shortcut
  --executable=[value]        executable for shortcut
  --launch-options=[value]    launch options for shortcut
  --icon-url=[value]          icon URL for shortcut
  --logo-url=[value]          logo URL for shortcut
  --cover-url=[value]         cover URL for shortcut
  --banner-url=[value]        banner URL for shortcut
  --hero-url=[value]          hero URL for shortcut
  --tags=[value,value]        tags for shortcut, comma separated

install:
  --programs=[value,...]      list of programs to install
  --preferences=[value,...]   preferences when installing programs

remove:
  --programs=[value,...]      list of programs to remove
  --preferences=[value,...]   preferences when removing programs

backup-state:
  --platforms=[value,...]     platforms to backup emulators state
  --preferences=[value,...]   preferences when synchronizing state

restore-state:
  --platforms=[value,...]     platforms to restore emulators state
  --preferences=[value,...]   preferences when synchronizing state

process-roms:
  --platforms=[value,...]     platforms to process the ROMs
  --preferences=[value,...]   preferences when processing ROMs (rebuild)

server:
  --gui=[value]               GUI mode (default|headless)
  --address=[value]           custom address for the server
  --dev                       enable development mode for static resources