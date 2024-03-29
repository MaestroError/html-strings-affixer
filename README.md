# html-strings-affixer

CLI tool that finds texts in HTML and replaces it with suffixed and prefixed string. Originally developed to replace strings with localization function in blade file:  
`<p>Some nice string</p>` -> `<p>{{ __('Some nice string') }}</p>`  
But can be used in any file which contains html (".jsx", ".vue", ".twig") with same purpose. Of course, prefix and suffix are customizable

_JavaScript version is already available [here](https://github.com/MaestroError/hsa-js) ([hsa-js](https://github.com/MaestroError/hsa-js))_

#### Quick start

Check out a little [article](https://medium.com/towardsdev/translate-laravel-blade-files-easily-f3753f5395bd) on Medium about the tool.

# Navigation

- [Installation](#installation)
  - [Composer](#composer)
  - [Linux/Unix](#linuxunix)
  - [Windows](#windows)
  - [MacOS](#macos)
  - [From source](#from-source)
- [Features](#features)
- [Config file](#config-file)
  - [JSON object and descriptions](#json-object-and-descriptions)
- [Commands](#commands)
  - [Replace command options](#replace-command-options)
  - [Check command options](#check-command-options)
- [Usage](#usage)
  - [Just affix it!](#just-affix-it)
  - [Force replace](#force-replace)
  - [One File](#work-with-one-file)
  - [Too many warnings](#too-many-warning)
  - [Choose Executable](#executables)
- [Support](#support)

# Installation

### Composer

It is already available in composer now, you can require it:  
`composer require maestroerror/html-strings-affixer --dev`.  
After installation it is accessible by:

- Linux: `./vendor/bin/hsa`
- Windows: `./vendor/bin/hsawin`
- MacOS: `./vendor/bin/hsamac`
- MacOS arm64 (m1): `./vendor/bin/hsamacm1`

### Linux/Unix

**Step 1: Download**  
Download binary [file](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsa) with `wget https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsa`

**Step 2: Make it executable**  
Run in directory with downloaded file `chmod a+x ./hsa` (Use sudo if permission is required)

**Step 3: Execution**  
You can move it in directory where you need to use and just execute with `./hsa [command] [-args]` or Make it available viw command line.

**Step 4: Command line _(optional)_**  
Find your bin folder (`/usr/bin`, `/usr/sbin` or `~/bin`), or make it with `cd ~/ && mkdir bin` and Symlink (`ln -s $~/path/to/directory/hsa ~/bin/hsa`) or move/copy binary file in **bin** directory. After it you can use app in cli with "hsa" command. Try: `hsa check`

### Windows

Download binary [file](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsa.exe) or [Zip](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsawin.zip). Get "hsa.exe" file in needed directory and run with ".\hsa.exe"

### MacOS

Download an [archive](https://github.com/MaestroError/html-strings-affixer/releases/latest/download/hsamac.zip), Unzip it, get "hsamac" or "hsamacm1" file in needed directory, give to it executable permissions and run with `./hsamac` or `./hsamacm1`

### From source

If you have [golang](https://go.dev/doc/install) installed, you can clone this repo and run `go install` or `go build` for binary file

# Features

- Finds your HTML visible strings and replaces with affixed one
- Prefix and Suffix are customizable
  - By default set as "{{ __('" and "') }}"
- If string contains one of the warning characters, it not replaces, but prints out location:
  - So it will not replace string with variable like "Price: {{$price}}", if you have '{' in warning character, it will give you exact location(file:line) to make it translatable manually
- If string contains one of the ignoring characters, it just ignores string
  - For example, if your string is math expression like 4 _ 20 = 80 and you have "_" in ignoring characters, it will just ignore it
- This characters are set by default as:
  - ignore: "#", "\_", ">", "^", "\*", "="
  - warnings: "%", "{", "(", "}", ")", "$", "'"
- Ignore characters and warning characters are customizable from JSON config file ("ignore" and "warnings")

# Config file

You need to create "affixer.json" file in directory, from where you will run html-strings-affixer. You can find an [example](https://github.com/MaestroError/html-strings-affixer/blob/maestro/bin/affixer-example.json) in bin folder of this repository (bin/affixer-example.json).  
_Alternativly you can run app without config file in your working directory and it will offer you to create one from example_

### JSON object and descriptions

```
    // Scanning
    *(string) "folder" - just folder to scan
    *(array) "file_types" - Parses file only with given extensions
    (array) "ignore_names" - ignores files and folders with given names

    // Parse
    (array) "ignore" - ignores strings which contains given character (Group of characters not allowed, only single characters)
    (array) "warnings" - Warns about strings which contains given characters (not replaces)
    (array) "methods" - Uses only given parse methods. Available: text, placeholder, alt, title, hashtag

    // Replace
    *(string) "prefix" - Prefix to set
    *(string) "suffix" - Suffix to set
    (bool) "force" - If true, git status check is ignored

    // Report
    (bool) "detailed" - if true, detailed report will be displayed
    (bool) "warn_as_Table" - if true, warnings will be displayed as table
    (bool) "report" - if true, warnings will be saved as JSON file

    // Log
    (string) "log_folder" - folder to store logs
```

# Commands

Available commands:

- **replace** - Main command, which makes replacement of strings
- **check** - Checks folder and gives report with files and count of found strings
- **clear-log** - If you use log_folder config, logs are generated. this command clear all log files (Has no arguments)

Some configs you can pass as arguments, use `hsa [command] --help` to read more about command.

### _Replace_ command options

- -allowed="(string)" - allowed file types, separated by commas (**Required** for security reasons)
- -detailed - If passed, detailed report printed
- -file="(string)" - Use this argument to run command only on one file
- -folder="(string)" - Folder to scan (**Required**)
- -force - If passed, git status check is ignored
- -only="(string)" - Methods to use while parsing, separated by commas. Available: text, placeholder, alt, title, hashtag
- -prefix="(string)" - New prefix for strings (**Required**)
- -suffix="(string)" - New suffix for strings (**Required**)
- -report - if true, warnings will be saved as JSON file

### _Check_ command options

- -allowed="(string)" - allowed file types, separated by commas (**Required**)
- -file="(string)" - Use this argument to run command only on one file
- -folder="(string)" - Folder to scan (**Required**)
- -only="(string)" - Methods to use while parsing, separated by commas. Available: text, placeholder, alt, title, hashtag
- -report - if true, warnings will be saved as JSON file

# Usage

It is recommended to use tool with [config file](#config-file), where you can specify all needed info and use all features, that tool is providing. If you already have configured file you can easily run:

```
~$ ./vendor/bin/hsa check
# AND/OR
~$ ./vendor/bin/hsa replace
```

In this usage guide, we will cover some CLI use cases, but you can use them **alongside** with config file, because options provided from command line has **higher priority**. They re-write all passed options while specific run, but all **other configs** provided by JSON file stays same.

#### Just Affix it!

For example, you can: Easily affix all html strings with @lang directive in templates folder. To do so, run:

```
./vendor/bin/hsa replace -folder="resources/views" -allowed=".blade.php" -prefix="@lang('" -suffix="')"
```

in another words:

```
hsa replace -folder="[Path/To/Your/Folder]" -allowed="[Allowed, File, Extensions]" -prefix="[StringToPrepend]" -suffix="[StringToAppend]"
```

_Note: Keep in mind that "@" character is not a part of default ignore or warning chars, you need to specify it via config file_

#### Force replace

If you use git and have uncommitted changes, it is recommended to first commit or stash them and after run replace command, so hsa will ask you about it on every run. If you already tested the tool and you know what are you doing, you can use _-force_ parameter to skip that prompt.

```
./vendor/bin/hsa replace -folder="views" -force
```

_Note: For now, HSA has not UNDO command, so be careful while running replace_

#### Work with one file

Sometimes you will need to perform some commands only on single file, to reduce memory usage, get smaller printed information and/or avoid unwanted changes. For that purpose you can use option _-file_:

```
./vendor/bin/hsa replace -file="resources/view/contact.blade.php"
```

#### Too many warning

Sometimes with large project there are too many warning to check them via terminal. For such cases -report option could be helpful, it will generate JSON file with all warnings in it:

```
./vendor/bin/hsa check -report
```

Note that `-report` option is availabe for both, `replace` and `check` commands. Also, it can be set via JSON file as `report: true`.

#### Executables

There are several executables for different platforms, choose your one and use it. They are working the exactly same way, you are changing just the executable file name in command.

- Linux: `./vendor/bin/hsa`
- Windows: `./vendor/bin/hsawin`
- MacOS: `./vendor/bin/hsamac`
- MacOS arm64 (m1): `./vendor/bin/hsamacm1`

_Note: I am planning to create single command, which will work for any platform, but for now, it is as it is :smile:_

# Support

Support Our Work? 🌟 You can help us keep the code flowing by making a small donation. Every bit of support goes a long way in maintaining and improving our open-source contributions. Click the button below to contribute. Thank you for your generosity!

[<img src="https://github.com/MaestroError/resources/blob/maestro/buymeamilk/green-2.png" width="300px">](https://www.buymeacoffee.com/maestroerror)

Or use QR code:

[<img src="https://github.com/MaestroError/resources/blob/maestro/buymeamilk/qr-code.png" width="135px">](https://www.buymeacoffee.com/maestroerror)

---

---

### To Do

- Make check command to print files and found strings count +
- Create clear-log command and print logs size as message in bootstrap +
- Write documentation in readme file +
- Release #1 - 0.0.1 +
- register on packagist (Composer require) +

---

- Test linux binary on ubuntu +
- Add composer install instructions in docs +
- Add linux install instructions in documentation +
- Make windows exe/msi installer ?
- Standardize Json configs +
- Update example config file with all needed configs +
- Add warning characters and separate them from ignoring characters +
- Consider similar cases `Veelgestelde vragen over {{$serviceName}}` +
  - add OT-OT(Opening Tag) and CT-CT(Closing Tag) extractions in "text" group +
  - add CT-OT +
- Print warning characters strings as warnings (not replaced) with file and line on replace command +
- Add warnings count in check command +
- Update documentation +
- Generate files for bin folder +
- Next Release (Don't forget to add in release: hsa, hsa.exe, hsawin.zip and hsaInstaller.msi) +

---

- Use error handling instead some of default configs, when one is not specified (From CLI or Config file) +
  - Folder | File (not use default folder and check it before running commands) +
  - allowed +
  - prefix (for replace command) +
  - suffix (for replace command) +
- Fix error while using HSA out of git repo +
- Add success message function in reporter +
- Make warnings table in reporter +
- Make warnings as table controllable from config and add it in Docs +
- Test for latest changes +
- Next Release +

---

Updated 21/09/2023:

- Add --report option for detailed report (JSON) in check command ([issue](https://github.com/MaestroError/html-strings-affixer/issues/2)) +
- Add `@` and `{{--` in default ignore characters ([issue](https://github.com/MaestroError/html-strings-affixer/issues/3)) +
- Update regex to parse laravel comments correctly ([issue](https://github.com/MaestroError/html-strings-affixer/issues/5)) +
- Fix `index out of range [-1]` [issue](https://github.com/MaestroError/html-strings-affixer/issues/4) +
- Release +

---

- Make TrimSpaces controllable as configuration from JSON
- Handle errors to continue replacement on other files and add error location in output (report)
- Test for latest changes
- Next Release

### Upcoming

- Write benchmarks for app
- Make app available by npm ([article](https://blog.bywachira.com/post/create-cli-with-golang-and-publish-on-npm))
- Install and configure goreleaser
- Refactor app with Cobra package
- Use log files to undo last changes in folder_to_scan from log_folder (for dry run)
- add command: "watch" (for live updates) and "undo" (Undo last changes)
- Make single executable/command (Bridge) for any platform

#### Resources

- [goreleaser](https://goreleaser.com/) +
- [go-npm](https://github.com/sanathkr/go-npm)
- [Developing and publishing modules](https://go.dev/doc/modules/developing)
