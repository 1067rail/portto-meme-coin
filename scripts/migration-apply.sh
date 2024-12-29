#! /bin/sh

atlas migrate apply --dir "file://migrations" --url "postgres://user:pass@localhost:5432/meme-coin?sslmode=disable"
