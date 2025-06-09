Usage: nicedeck [COMMAND] [OPTIONS]

COMMAND:

version         print version information
help            print application help
programs        list available programs
platforms       list available platforms
shortcuts       list current user shortcuts
scraper         scrape data on SteamGridDB
launch          launch shortcut with given appID
modify          update or delete shortcut with given appID
install         install or update programs
remove          remove previously installed programs
sync-state      sync emulators state
process-roms    process emulators ROMs
server          start server for GUI usage

OPTIONS:

scraper:
  --search=[term]           search term

launch:
  --id=[appID]              shortcut appID

modify:
  --update                  update shortcut
  --delete                  delete shortcut
  --id=[appID]              shortcut appID
  --app-name=[value]        name for shortcut
  --start-dir=[value]       start directory for shortcut
  --exe=[value]             executable for shortcut
  --launch-options=[value]  launch options for shortcut
  --icon-url=[value]        icon URL for shortcut
  --logo-url=[value]        logo URL for shortcut
  --cover-url=[value]       cover URL for shortcut
  --banner-url=[value]      banner URL for shortcut
  --hero-url=[value]        hero URL for shortcut

install:
  --programs=[value,...]      list of programs to install

remove:
  --programs=[value,...]      list of programs to remove

sync-state:
  --platforms=[value,...]     platforms to sync emulators state
  --preferences=[value,...]   preferences when synchronizing state (dump-state|restore-state)

process-roms:
  --platforms=[value,...]     platforms to process the ROMs
  --preferences=[value,...]   preferences when processing ROMs (rebuild)

server:
  --gui=[value]           GUI mode (default|headless)
  --address=[value]       custom address for the server
  --dev                   enable development mode for static resources