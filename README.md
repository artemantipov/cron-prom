# cron-pron
Simple Go-cron based on https://github.com/robfig/cron with prometheus metric of failed jobs count.

### Possible environment variables
* CRON_JOB_* - define cron job (e.g. `CRON_JOB_Bla="* * * * * do foo --bar`) (required)
* METRICS_PORT - default is *1221* (optional)
* METRICS_URL - default */metrics* (optional)
* METRIC_PREFIX - default metric name is *cron_jobs_failed* (e.g `METRIC_PREFIX="foo_"` > *foo_cron_jobs_failed*) (optional)


