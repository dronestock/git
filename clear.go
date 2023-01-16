package main

func (p *plugin) clear() (undo bool, err error) {
	if undo = !p.clearable() || p.pulling(); undo {
		return
	}

	return
}
