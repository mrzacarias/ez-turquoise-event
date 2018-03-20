# EZ Turquoise Event

Helper to generate Turquoise's `event.json`

## Setup
1. Clone this repo under your `$GOPATH`
1. Run the setup and build script: `./script/setup_and_build.sh`

## Usage
```
-vert string
  Vertical name (home|life) (default "home")
-env string
  Vertical environment (staging|prod) (default "staging")
-new string
  New version (will be scaled up) (default "REQUIRED")
-old string
  Old version (will be scaled down) (default "REQUIRED")
-notify string
  List of people to notify (default "@jacob.murphy")
-percentage string
  Percentage of instances with new version (10|33|50|100) (default "100")
```

## Running
`bin/ez-turquoise-event -vert=life -env=staging -new=x2150.0.0 -old=x2132.0.0 -notify="@..." percentage=100`

That will create `life_staging_x2150.0.0_100.json`:
```
Slack messages:
  At lynx-life-deploy and consumer-facing-eng:
  @here Lynx life x2150.0.0 has been deployed to staging at 100%:
  https://github.com/adharmonics/lynx-life/releases/tag/x2150.0.0
  @... please check your changes and message me if they look good

Event.json by percentage:
- - - - - - - - - - CUT HERE FOR 10% - - - - - - - - - -
[
  {
    "appName": "lynx-life-staging-sidekiq",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-worker": 9,
      "lynx-life-automated-x2150.0.0-worker": 1
    }
  },
  {
    "appName": "lynx-life-staging-web",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-web": 9,
      "lynx-life-automated-x2150.0.0-web": 1
    }
  }
]

- - - - - - - - - - CUT HERE FOR 33% - - - - - - - - - -
[
  {
    "appName": "lynx-life-staging-sidekiq",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-worker": 2,
      "lynx-life-automated-x2150.0.0-worker": 1
    }
  },
  {
    "appName": "lynx-life-staging-web",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-web": 2,
      "lynx-life-automated-x2150.0.0-web": 1
    }
  }
]

- - - - - - - - - - CUT HERE FOR 33% A/B - - - - - - - -
[
  {
    "appName": "lynx-life-staging-sidekiq",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-worker": 1,
      "lynx-life-automated-x2150.0.0a-worker": 1,
      "lynx-life-automated-x2150.0.0b-worker": 1
    }
  },
  {
    "appName": "lynx-life-staging-web",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-web": 1,
      "lynx-life-automated-x2150.0.0a-web": 1,
      "lynx-life-automated-x2150.0.0b-web": 1
    }
  }
]

- - - - - - - - - - CUT HERE FOR 50% - - - - - - - - - -
[
  {
    "appName": "lynx-life-staging-sidekiq",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-worker": 1,
      "lynx-life-automated-x2150.0.0-worker": 1
    }
  },
  {
    "appName": "lynx-life-staging-web",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-web": 1,
      "lynx-life-automated-x2150.0.0-web": 1
    }
  }
]

- - - - - - - - - - CUT HERE FOR 100% - - - - - - - - - -
[
  {
    "appName": "lynx-life-staging-sidekiq",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-worker": 1,
      "lynx-life-automated-x2150.0.0-worker": 1
    }
  },
  {
    "appName": "lynx-life-staging-web",
    "launchConfigurations": {
      "lynx-life-automated-x2132.0.0-web": 1,
      "lynx-life-automated-x2150.0.0-web": 1
    }
  },
  {
    "appName": "lynx-life-staging-web",
    "launchConfigurations": {
      "lynx-life-automated-x2150.0.0-web": 1,
      "lynx-life-automated-x2132.0.0-web": 0
    }
  },
  {
    "appName": "lynx-life-staging-sidekiq",
    "launchConfigurations": {
      "lynx-life-automated-x2150.0.0-worker": 1,
      "lynx-life-automated-x2132.0.0-worker": 0
    }
  }
]
```
