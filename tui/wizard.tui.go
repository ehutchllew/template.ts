package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ehutchllew/template.ts/cmd/models"
)

type WizardAnswers struct {
	appName                  string
	appNameTextInput         textinput.Model
	appNameTextInputRendered bool

	cursor models.ToolNames

	selected map[models.ToolNames]bool
	styles   *Styles

	height int
	width  int
}

func New() *models.UserAnswers {
	logFile, logFileErr := tea.LogToFile("debug.log", "debug")
	if logFileErr != nil {
		log.Fatalf("log file err: %v", logFileErr)
	}
	defer logFile.Close()

	w := WizardAnswers{
		selected: map[models.ToolNames]bool{
			models.ES_BUILD:   false,
			models.ES_LINT:    false,
			models.JEST:       false,
			models.SWC:        false,
			models.TYPESCRIPT: false,
		},
		styles: DefaultStyles(),
	}

	p := tea.NewProgram(w)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error starting bubbletea: %v", err)
	}

	log.Println("have we hit this yet????")
	userAnswers := &models.UserAnswers{
		AppName:    w.appName,
		EsBuild:    w.selected[models.ES_BUILD],
		EsLint:     w.selected[models.ES_LINT],
		Jest:       w.selected[models.JEST],
		Swc:        w.selected[models.SWC],
		Typescript: w.selected[models.TYPESCRIPT],
	}

	return userAnswers
}

func (w WizardAnswers) Init() tea.Cmd {
	return nil
}

func (w WizardAnswers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// not working
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.height = msg.Height
		w.width = msg.Width
	}

	if w.appName == "" {
		if w.appNameTextInputRendered == false {
			w.appNameTextInput = textinput.New()
			w.appNameTextInput.Placeholder = "Type your App Name"
			w.appNameTextInput.Focus()
			w.appNameTextInputRendered = true
		}
		return w.updateTextInput(msg)
	} else {
		return w.updateSelector(msg)
	}
}

func (w WizardAnswers) View() string {
	if w.width == 0 {
		return "Loading..."
	}

	if w.appName == "" {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			w.appNameTextInput.Value(),
			w.styles.InputField.Render(w.appNameTextInput.View()),
		)
	}

	// header
	// s := "Select all the configurations that you would like to create\n\n"

	// cursor := " "

	return ""
}

func renderRow(cursor string, checked string, choice string) string {
	return fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
}

func (w WizardAnswers) updateSelector(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return w, tea.Quit
		case "up", "k":
			if w.cursor > 0 {
				w.cursor--
			}
		case "down", "j":
			if int(w.cursor) < len(w.selected)-1 {
				w.cursor++
			}
		case "enter", " ":
			b, ok := w.selected[w.cursor]
			if ok {
				w.selected[w.cursor] = !b
			} else {
				log.Fatalf("Couldn't access selected map. Accessed with key: (%d)", w.cursor)
			}
		}
	}

	return w, nil
}

func (w WizardAnswers) updateTextInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	w.appNameTextInput, cmd = w.appNameTextInput.Update(msg)
	log.Printf("msg:%+v", msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		case "enter":
			w.appNameTextInput.SetValue("done!")
			return w, nil
		}
	}
	return w, cmd
}
