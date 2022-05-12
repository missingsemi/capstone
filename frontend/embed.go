package frontend

import "embed"

//go:embed index.html scripts/*
var Frontend embed.FS
