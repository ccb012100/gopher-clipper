# GopherClip

Clipboard utility written in Go

## TODO

### Functionality

- [ ] Show stack of clips
- [ ] Copy item from stack
  - [ ] `mouse-1`
  - [ ] `Ctrl-C`
  - [ ] GUI button
- [ ] Delete from stack
  - [ ] `mouse-2`
  - [ ] GUI button
- [ ] Paste to top of stack
  - [ ] `mouse-2`
  - [ ] `Ctrl-V`
  - [ ] GUI button
- [ ] Cut from stack
  - [ ] `Ctrl-X`
  - [ ] GUI button
- [ ] Copy stack to clipboard
  - [ ] GUI button
  - [ ] `Ctrl-Shift-C`
- [ ] Export stack to file
  - either to `$HOME/Downloads/gopher-clip-stack-YYYY-MM-DD-HHMMSS.txt` or use a file save GUI
- [ ] Switch logging to [logrus](https://github.com/sirupsen/logrus)

#### Maybe

- [ ] Undo paste/delete on `Ctrl-Z`
- [ ] Redo paste/delete on `Ctrl-Y`
- [ ] Drag-and-drop stack items

### SQLite

- [ ] Sync clipboard to SQLite database
  - [ ] Copy
  - [ ] Delete
  - [ ] Paste

### Visual

- [ ] Set theme based on system
- [ ] Add multiple dark/light themes

### Settings

- [ ] Add settings file
  - [ ] theme
  - [ ] DB location
  - [ ] (maybe) shortcuts

## Attributions

![clipboard icon](Icon.png) Clipboard icon created by [Mehwish - Flaticon](https://www.flaticon.com/free-icons/clipboard).
