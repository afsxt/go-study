package label_service

import "base-server/models"

//-----------------------------------------------------------------------------

type Label struct {
	ID    int
	Text  string
	Audio string
}

func (l *Label) Add() error {
	label := map[string]interface{}{
		"id":    l.ID,
		"text":  l.Text,
		"audio": l.Audio,
	}

	err := models.AddLabel(label)
	return err
}
