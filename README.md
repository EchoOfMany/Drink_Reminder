# Simple Drink reminder

Just small desktop project to practice Golang and create something to help the better half drink enough

**NOT USED** for my reference use: `go build -ldflags -H=windowsgui` to build so that the cmd does show when the exe is run
use `fyne package -os windows -icon drink.png` to build. Does the above and adds program icon

### TO DO

- [x] restart timer on completion
- [ ] add manual start/stop ( possibly pause )
- [ ] allow time to be edited in app
- [ ] refactor so its not a mess

### Known issue

- Issue with the first completion where the desktop notification fires twice, further completes work as expected
