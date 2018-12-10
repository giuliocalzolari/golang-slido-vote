# golang-slido-vote
[![Release](https://img.shields.io/github/release/giuliocalzolari/golang-slido-vote.svg?style=flat-square)](https://github.com/giuliocalzolari/golang-slido-vote/releases/latest)
[![Travis](https://img.shields.io/travis/giuliocalzolari/golang-slido-vote.svg?style=flat-square)](https://travis-ci.org/giuliocalzolari/golang-slido-vote)
[![Go Report Card](https://goreportcard.com/badge/github.com/giuliocalzolari/golang-slido-vote?style=flat-square)](https://goreportcard.com/report/github.com/giuliocalzolari/golang-slido-vote)


A simple command line tool to consume partitions of a topic and push to kinesis

### Installation

    git clone git@github.com:giuliocalzolari/golang-slido-vote.git
    cd ./golang-slido-vote/
    go build

    # or

    go run main.go -votes=42 -event=YOUREVENT  -question=111111


### Usage

    $ Usage of ./slido-vote:
      -event string
            event code (default "xxxxx")
      -flushlog string
            sets the flush trigger level (default "none")
      -log string
            sets the logging threshold (default "info")
      -question string
            question id (default "zzzzz")
      -stderr
            outputs to standard error (stderr)
      -votes int
            votes count (default 42)