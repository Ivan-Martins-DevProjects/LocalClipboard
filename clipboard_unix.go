//go:build !windows

package main

import (
	"errors"
	"os/exec"
	"runtime"
)

func (a *App) writeNativeClipboard(text string) error {
	if runtime.GOOS == "darwin" {
		return a.runCommand("pbcopy", text)
	}

	// No Linux, tentamos os utilitários mais comuns em ordem
	commands := [][]string{
		{"xclip", "-selection", "clipboard"},
		{"xsel", "--clipboard", "--input"},
		{"wl-copy"}, // Para usuários de Wayland puro (ex: Fedora modern)
	}

	for _, cmd := range commands {
		if _, err := exec.LookPath(cmd[0]); err == nil {
			return a.runCommand(cmd[0], text, cmd[1:]...)
		}
	}

	return errors.New("nenhum utilitário de clipboard encontrado (instale xclip, xsel ou wl-clipboard)")
}

func (a *App) runCommand(name string, text string, args ...string) error {
	cmd := exec.Command(name, args...)
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = in.Write([]byte(text))
	in.Close()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
