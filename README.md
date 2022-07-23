# html-strings-affixer
Finds texts in HTML and replaces it with suffixed and prefixed string


### To Do
- Make check command to print files and found strings count
- Create clear-log command and print logs size as message in bootstrap
- Write documentation in readme file
- Release #1 - 0.0.1
- register on packagist (Composer require)
- Add warning characters and separate them from ignoring characters
- Print warning characters strings as warnings (not replaced) with file and line on replace command
- Add --expand option for detailed info in check command
- Add --report option for detailed report (JSON) in check command
- Write benchmarks for app
- Release #2 - 0.0.2
- .deb installable package (apt-get)
- exe/msi installer


### Upcoming
- Refactor app with Cobra package
- Use log files to undo last changes in folder_to_scan from log_folder (for dry run)
- add command: "watch" (for live updates) and "undo" (Undo last changes)