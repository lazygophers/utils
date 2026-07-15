//go:build !debug && !dev && !development && !alpha && !beta && !release && !prod && !production && !test && !canary

package app

var Env = Debug
