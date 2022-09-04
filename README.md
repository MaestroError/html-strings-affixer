# html-strings-affixer
Finds texts in HTML and replaces it with suffixed and prefixed string. Originally developed to replace strings with localization function in blade file:      
`<p>Some nice string</p>` -> `<p>{{ __('Some nice string') }}</p>`     
But can be used in any file which contains html (".jsx", ".vue", ".twig") with same purpose. Of course, prefix and suffix are customizable

# Navigation
- [Installation](#installation)
    - [Composer](#composer)
    - [Linux/Unix](#linuxunix)
    - [Windows](#windows)
    - [From source](#from-source)
- [Features](#features)
- [Config file](#config-file)
- [Commands](#commands)

# Installation
### Composer
It is already available in composer now, you can require it:       
`composer require maestroerror/html-strings-affixer --dev`.     
After installation it is accessible by `./vendor/bin/hsa` and `./vendor/bin/hsawin` for Windows

### Linux/Unix
**Step 1: Download**   
Download binary [file](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsa) with `wget https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsa`
       
**Step 2: Make it executable**      
Run in directory with downloaded file `chmod a+x ./hsa` (Use sudo if permission is required)
       
**Step 3: Execution**       
You can move it in directory where you need to use and just execute with `./hsa [command] [-args]` or Make it available viw command line.
        
**Step 4: Command line *(optional)***     
Find your bin folder (`/usr/bin`, `/usr/sbin` or `~/bin`), or make it with `cd ~/ && mkdir bin` and Symlink (`ln -s $~/path/to/directory/hsa ~/bin/hsa`) or move/copy binary file in **bin** directory. After it you can use app in cli with "hsa" command. Try: `hsa check`

### Windows
Download binary [file](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsa.exe) or [Zip](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsawin.zip). Get "hsa.exe" file in needed directory and run with ".\hsa.exe" **OR** Download .msi [installer](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsaInstaller.msi) and follow installation process


### From source
If you have [golang](https://go.dev/doc/install) installed, you can clone this repo and run `go install` or `go build` for binary file

# Config file
You need to create "affixer.json" file in directory, from where you will run html-strings-affixer. You can find an [example](https://github.com/MaestroError/html-strings-affixer/blob/maestro/bin/affixer-example.json) in bin folder of this repository (bin/affixer-example.json).      
*Alternativly you can run app without config file in your working directory and it will offer you to create one from example*       
### JSON object and descriptions
```
    // Scanning
    (string) "folder" - just folder to scan
    (array) "file_types" - Parses file only with given extensions
    (array) "ignore_names" - ignores files and folders with given names
    
    // Parse
    (array) "ignore" - ignores strings which contains given character
    (array) "methods" - Uses only given parse methods. Available: text, placeholder, alt, title, hashtag

    // Replace
    (string) "prefix" - Prefix to set
    (string) "suffix" - Suffix to set
    (bool) "force" - If true, git status check is ignored

    // Report
    (bool) "detailed" - if true, detailed report will be displayed

    // Log
    (string) "log_folder" - folder to store logs
```

# Commands
Available commands:
- replace - Main command, which makes replacement of strings
- check - Checks folder and gives report with files and count of found strings
- clear-log - If you use log_folder config, logs are generated. this command clear all log files    

Some configs you can pass as arguments, use `hsa [command] --help` to read more about command. 


### To Do
- Make check command to print files and found strings count +
- Create clear-log command and print logs size as message in bootstrap +
- Write documentation in readme file
- Release #1 - 0.0.1 +
- register on packagist (Composer require) +
-------
- Test linux binary on ubuntu +
- Add composer install instructions in docs +
- Add linux install instructions in documentation +
- Make windows exe/msi installer ?
- Standardize Json configs +
- Update example config file with all needed configs +
- Add warning characters and separate them from ignoring characters +
- Consider similar cases `Veelgestelde vragen over {{$serviceName}}`
    - add OT-OT(Opening Tag) and CT-CT(Closing Tag) extractions in "text" group
- Print warning characters strings as warnings (not replaced) with file and line on replace command
- Generate files for bin folder
- Next Release (Don't forget to add in release: hsa, hsa.exe, hsawin.zip and hsaInstaller.msi)
-------
- Add --expand option for detailed info in check command
- Add --report option for detailed report (JSON) in check command
- Next Release
-------
- Next Release


### Upcoming
- Write benchmarks for app
- Make app available by npm ([article](https://blog.bywachira.com/post/create-cli-with-golang-and-publish-on-npm)) 
- Install and configure goreleaser
- Refactor app with Cobra package
- Use log files to undo last changes in folder_to_scan from log_folder (for dry run)
- add command: "watch" (for live updates) and "undo" (Undo last changes)

#### Resources
- [goreleaser](https://goreleaser.com/) +
- [go-npm](https://github.com/sanathkr/go-npm)
- [Developing and publishing modules](https://go.dev/doc/modules/developing)