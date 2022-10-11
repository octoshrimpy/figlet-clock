package main

import (
  "fmt"
  "os"
  "sync"
  "time"

  "github.com/charmbracelet/bubbles/help"
  "github.com/charmbracelet/bubbles/key"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
  "github.com/mbndr/figlet4go"
  //       "fmt"
  //       "io"
  //       "log"
  //       "strings"
  //       "github.com/76creates/stickers"
  //       "github.com/charmbracelet/bubbles/viewport"
  //       "github.com/charmbracelet/glamour"
  //       "github.com/knipferrc/teacup/code"
  //       "github.com/knipferrc/teacup/statusbar"
  //       "github.com/treilik/bubbleboxer"
)

// necessary for keeping track of time
var (
  lastID int
  idMtx  sync.Mutex
)

// so we can access it later when it self updates
func nextID() int {
  idMtx.Lock()
  defer idMtx.Unlock()
  lastID++
  return lastID
}

type window struct {
  Width     int
  Height    int
}

// declare the data that the clock will need
type model struct {
  font      string    // name of font
  x_offset  int       // offset from left edge
  y_offset  int       // offset from top edge
  pos       string    // position of clock,
                      // T = top, C = center, B = bottom
                      // L = left, C = center, R = right
                      // examples: TC, BL, CC, BR
  msg     string      // cute message under the clock
  help    help.Model
  keymap  keymap
  id      int
  view    window
  use_12  bool
  ascii   *figlet4go.AsciiRender
}

// so we can leave
type keymap struct {
  quit key.Binding
}

// get self id
func (m model) ID() int {
  return m.id
}

// generating a new clock on app spawn
func New() model {
  // @todo catch user inputs here, to create new model
  ascii := figlet4go.NewAsciiRender()

  // return the built model
  return model{
    id:       nextID(),
    font:     "default",
    pos:      "CC",
    msg:      "strive to be kind to yourself and others",
    x_offset: 0,
    y_offset: 0,
    use_12: true,
    ascii: ascii,
  }
}

// generate the help menu at the bottom
// @todo later on add controls for changing font, size, color, etc
func (m model) HelpView() string {
  return "\n" + m.help.ShortHelpView([]key.Binding{
    m.keymap.quit,
  })
}

// base init
func (m model) Init() tea.Cmd {
  return nil
}

// returns the stringified version of the app output
func (m model) View() string {

  // get curent time in computer's set timezone
  hou, min, sec := time.Now().Local().Clock()

  // 12-hour time setting
  if m.use_12 {
    hou = hou % 12
  }

  s_hou := fmt.Sprintf("%02d", hou)
  s_min := fmt.Sprintf("%02d", min)
  s_sec := fmt.Sprintf("%02d", sec)

  // convert it all to a string (Itoa is a terrible name)
  time := s_hou + " : " + s_min + " : " + s_sec

  //generate figlet render using the time we got
  ascii := figlet4go.NewAsciiRender()
  // ascii.LoadFont
  figletOpts := figlet4go.NewRenderOptions()
  
  ascii.LoadFont("fonts/")

  figletOpts.FontName = "doom"
  // if m.font == "default" {
  // }

  renderStr, _ := ascii.RenderOpts(time, figletOpts)

  gloss := lipgloss.NewStyle().
            Align(lipgloss.Center).
            Foreground(lipgloss.Color("#7D56F4")).
            Render(renderStr)

  block := lipgloss.
              Place(m.view.Width, m.view.Height, lipgloss.Center, lipgloss.Center, gloss)

  // return string
  return block
}

// runs every second, since that's what we've told it to do over at tick()
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    
		// commenting out all this stuff, close program if any input, as this is a screensaver
		// @todo convert this over to a settings toggle
		
		// switch {
    // // if quit key pressed
    // case key.Matches(msg, m.keymap.quit):
      // m.quitting = true
      return m, tea.Quit
    // }

  case tea.WindowSizeMsg:
    m.view.Width = msg.Width
    m.view.Height = msg.Height
    return m, nil
  }

  // return with tick every second, so it updates itself
  return m, tick(m.id, time.Second)
}

func main() {
  // @todo we're not rendering help menu, oops
  m := model{
    keymap: keymap{
      quit: key.NewBinding(
        key.WithKeys("ctrl+c", "q"),
        key.WithHelp("q", "quit"),
      ),
    },
  }

  // create new bubbletea app
  p := tea.NewProgram(m,

    // use full size of terminal
    tea.WithAltScreen(),

    // turn on mouse support
    tea.WithMouseAllMotion(),
  )

  // insert spongebob meme: SeLf-DoCuMeNtInG
  if err := p.Start(); err != nil {
    fmt.Println("something borked", err)
    os.Exit(1)
  }
}

// ju-lee, do the thing!
func tick(id int, d time.Duration) tea.Cmd {
  return tea.Tick(d, func(_ time.Time) tea.Msg {
    return nil
  })
}
