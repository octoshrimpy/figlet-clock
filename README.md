# figlet-clock
_a nice figlet-powered clock for your terminal needs_


![image](https://user-images.githubusercontent.com/7097172/195013255-8e8caf35-0f31-4818-a4c7-ab6fd8423500.png)

Made to replace the default tmux clock in `C-t`:

1. Clone the repo
```
git clone https://github.com/octoshrimpy/figlet-clock
```

2. Add to your `tmux.conf`
```
unbind t
bind t popup -d '#{pane_current_path}' -xC -yC -w100% -h100% -E 'go run ~/path/to/clock.go'
```

You can change the `-w100% -h100%` to whatever size you want. By default it will exit on any keystroke, so you can just slap your keyboard when you get back! 
All your windows will be waiting for you, as this uses tmux popups and does not touch your layouts.


---

## known bugs and todos
- [ ] takes a few seconds to start running, and appears to be frozen whenever it's launched
- [ ] catch user settings in New() `clock.go:95`
- [ ] create settings structs to leverage charm styling
- [ ] add option to quit on any input, `q`, `esc`, or any combo of the three
- [ ] figlet fonts are not monospaced! clock jumps left and right by one char depending on the digits
- [ ] fix the help menu not rendering, add option to show/hide it

### wishlist
- [ ] figure out mouse input, and exit on mouse jiggle + options in settings
- [ ] simple password lock + options in settings
  * catch all exit sequences?
  * text input vs clicky pin?

---

## License

[MIT](https://github.com/charmbracelet/bubbletea/raw/master/LICENSE)

---


Made with 💜 for [Charm](https://charm.sh/)

<a href="https://charm.sh/"><img alt="The Charm logo" src="https://stuff.charm.sh/charm-badge.jpg" width="400"></a>
