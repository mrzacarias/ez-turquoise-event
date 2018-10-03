# EZ Turquoise Event

Helper to generate Turquoise's `event.json`

## Setup
1. Clone this repo under your `$GOPATH`
1. Run the setup and build script: `./script/setup_and_build.sh`

## Usage
```
-env string
    Vertical environment (staging|prod) (default "staging")
-loser string
    Loser leg (100% A/B) (default "NO LOSER SET")
-new string
    New version (will be scaled up) (default "REQUIRED")
-notify string
    List of people to notify (default "@jacob.murphy")
-old string
    Old version (will be scaled down) (default "REQUIRED")
-percentage string
    Percentage of instances with new version (10|33|50|100) (default "100")
-vert string
    Vertical name (auto|home|life) (default "home")
-winner string
    Winner leg (100% A/B) (default "NO WINNER SET")
```

## Running
`./script/run.sh -vert=home -env=staging -old=x2387.0.0 -new=x2396.0.0 -loser=a -winner=b`
