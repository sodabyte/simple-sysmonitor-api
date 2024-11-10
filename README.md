This is a personal project for learning Go, exposing system monitoring data via an HTTP API for easy integration with web UIs.

## Notes

- The binary is mainly run in Termux. It may work on other Linux machines, but this hasn't been fully tested.
- To successfully retrieve CPU and partition data, the application must be run with **superuser** privileges. (Yes, my device is rooted with magisk.)
- The API will panic if the `/cpu-usage` endpoint is accessed without superuser privileges.

## Endpoints

- `GET /processes` - List of running processes.
- `GET /disk-usage` - Disk usage statistics.
- `GET /disk-io` - Disk input/output statistics.
- `GET /partitions` - System partition information.
- `GET /memory-usage` - Memory usage statistics.
- `GET /cpu-usage` - CPU usage statistics.

## Contribution

Feel free to open issues or pull requests if you'd like to contribute.
