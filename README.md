# RXJH-EMU PUBLIC VERSION
RXJH Emulator Public Project, Go lang server emulator of RXJH game
Development on client from [steam (scion of fate)](https://store.steampowered.com/app/747970/Scions_of_Fate/)

Libraries usage:
 - [easytcp](https://github.com/DarthPestilane/easytcp)
 - [Gorm](https://gorm.io/)
 
# Usage

dev run:

```
go run main.go --type login
```

```
go run main.go --type game
```

build:
```
go build -o server.exe // or no extension for other platform
```

Run Executable:

```
server.exe --type login or game
```
