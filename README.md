# Cachalot

Help developer get the snapshot of host machine runtime and/or container environment

## Badges

- [![Go](https://github.com/Incarnation-p-lee/cachalot/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/Incarnation-p-lee/cachalot/actions/workflows/go.yml)
- [![codecov](https://codecov.io/gh/Incarnation-p-lee/cachalot/branch/master/graph/badge.svg?token=kyWBu44Yuu)](https://codecov.io/gh/Incarnation-p-lee/cachalot)
- [![DeepSource](https://deepsource.io/gh/Incarnation-p-lee/cachalot.svg/?label=active+issues&show_trend=true&token=sfNFlwtPmXs8B7a9Dn7n0ERV)](https://deepsource.io/gh/Incarnation-p-lee/cachalot/?ref=repository-badge)
- [![Codacy Badge](https://app.codacy.com/project/badge/Grade/46a042f933084de2a04263e2cad1c25b)](https://www.codacy.com/gh/Incarnation-p-lee/cachalot/dashboard?utm_source=github.com&utm_medium=referral&utm_content=Incarnation-p-lee/cachalot&utm_campaign=Badge_Grade)
- [![Codiga Score](https://api.codiga.io/project/33659/score/svg)](https://app.codiga.io/project/33659/dashboard)
- [![Codiga Grade](https://api.codiga.io/project/33659/status/svg)](https://app.codiga.io/project/33659/dashboard)

## Usage

Use `./cachalot -h` to see more details about usage

```
Usage of ./cmd/cachalot/cachalot:
  -out string
        the output layout for print (default "text")
  -pids string
        the comma separated pids for snapshot, -1 indicates all processes (default "-1")
  -sampling-count string
        the total count of sampling (default "10")
  -sampling-interval string
        the interval for each sampling, count in seconds (default "10")
```

The output may look like below.

```
==========================================================
Print snapshot with iteration count 0
==========================================================
Timestamp: 2022-06-23 02:09:58.500744868 +0000 UTC m=+0.000049903
==========================================================
Total procesess count: 3
==========================================================
PID             CPUUsage        CmdLine
==========================================================
7               7.398%          java-jarapp.jar
1027            0.000%          bash
1038            0.000%          ./cachalot


==========================================================
Print snapshot with iteration count 1
==========================================================
Timestamp: 2022-06-23 02:10:13.510835481 +0000 UTC m=+15.010140416
==========================================================
Total procesess count: 3
==========================================================
PID             CPUUsage        CmdLine
==========================================================
7               7.314%          java-jarapp.jar
1038            0.000%          ./cachalot
1027            0.000%          bash
...
```

## Build

### Precondition

- [Ubuntu](https://ubuntu.com/)
- [Go 1.18](https://golang.google.cn/doc/go1.18)
- [Build Essential](https://pkgs.org/download/build-essential)

cachalot leverage only one makefile under the root folder of the project. The final binary will be built to `cmd/cachalot/cachalot` by commad `make`. Meanwhile, below commands of makefile are also supported.

```
make clean    # Cleanup building related files.
make release  # Build but stripped
make test     # Run the unit test
```

## Concept

### CPU Usage

#### Multiple Cores

Cachalot will count all the CPU ticks of host environment, so ticks of all the CPU cores will be counted for CPU usage. For example, there are 4 cores, 1 core is exhaused by one PID while others are free. Then the usage may be similar to `1 / 4 = 25%`.

#### Container

Cachalot will count all the CPU ticks of host machine, so the CPU usage will be more smaller than expectation. For example, there is one container with 600m limitation for CPU, running from one 8 cores host machine. Then useage may be similar to `0.6 / 8 = 7.5%`
