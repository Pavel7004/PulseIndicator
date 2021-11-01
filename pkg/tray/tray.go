package tray

import (
	"github.com/getlantern/systray"
	"github.com/pavel/PulseIndicator/pkg/icons"
)

func Start() {
	systray.Run(onReady, onExit)
}

func onReady() {
	icons := icons.Read("ePapirus")
	systray.SetIcon(icons.Low)
	systray.SetTooltip("Look at me, I'm a tooltip!")
}

func onExit() {

}
