# EZ Events json

Helper to generate Turquoise's `event.json`

```
Usage:
-env string
    Vertical environment (staging|prod) (default "staging")
-new string
    New version (will be scaled up) (default "REQUIRED")
-notify string
    List of people to notify (default "@jacob.murphy")
-old string
    Old version (will be scaled down) (default "REQUIRED")
-percentage string
    Percentage of instances with new version (10|33|50|100) (default "100")
-vert string
    Vertical name (home|life) (default "home")
```

Ex.:
`go run main.go -vert=life -env=staging -new=x2150.0.0 -old=x2132.0.0 -notify="@..." percentage=100`
