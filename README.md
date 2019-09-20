# EZ Turquoise Event

Helper to generate Turquoise's `event.json`

## Setup
1. Clone this repo under your `$GOPATH`
1. Run the setup and build script: `./script/setup_and_build.sh`

## Usage
```
Usage of bin/ez-turquoise-event:
  -vert string
      Vertical name (auto|home|life) (default "life")
  -env string
    	Vertical environment (staging|prod) (default "staging")
  -loser string
    	Loser leg (100% A/B) (default "a")
  -winner string
    	Winner leg (100% A/B) (default "b")
  -new string
    	New version (will be scaled up) (default "REQUIRED")
  -old string
    	Old version (will be scaled down) (default "REQUIRED")
```

## Running
`./script/run.sh -vert=auto -env=prod -old=v20190910.0b -new=v20190912.0`
