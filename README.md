# Pricy


## Install

```bash
brew tap stangirard/tap
brew install pricy
```

## Run

```bash
pricy
```

If you are using sso for credentials on aws

```bash
pricy --sso
```

## Usage

There are a couple of parameters that you can use
- `--details`: Show the details of the report with the pricing by service
- `--sso`: Use sso for credentials
- `--csv`: Output the report as csv to `reports.csv`
- `--evolution`: Show the evolution of the report as `evolution.csv`
- `--days`: Number of days to look back for
- `--interval`: Date Interval on which the report is generated (Default last 14 days) (Format: YYYY-MM-DD:YYYY-MM-DD)
- `--granularity`: Granularity of the report, can be `daily`,  `monthly`
- `--html`: Output the report as html to `pricy.html`
- `--prometheus`: Outputs as prometheus metrics on `http://localhost:2112/metrics` that can be scraped by prometheus

## Example

### HTML Report Generation 

```bash
pricy --sso --days 150 --granularity monthly --html
```

Generates a report for the price starting a 150 days ago

<p align="center">
<img src="docs/html-report.png" alt="html-report" width="40%">
<p align="center">


### Prometheus Metrics

```bash
pricy --sso --prometheus  
```

Generates the prometheus metrics for the price updating every 8 hours

<p align="center">
<img src="docs/prometheus.png" alt="prometheus" width="40%">
<p align="center">
