# Journey

`Journey` is a CLI Nomad test harness, which implements a fancy progress bar based on [go-glint](https://github.com/mitchellh/go-glint)

## Export runtime arguments

```bash
export JOBS=1
export JOBSPEC=sleep.nomad
```

## Start a journey

```bash
journey start
```

## Stop a journey

```bash
journey stop
```
