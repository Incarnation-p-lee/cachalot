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
        the output layout for print, supported formats are [text json] (default "text")
  -pids string
        the comma separated pids for snapshot, -1 indicates all processes (default "-1")
  -sampling-count string
        the total count of sampling (default "10")
  -sampling-interval string
        the interval for each sampling, count in seconds (default "10")
  -sorted-by string
        the metrics to be sorted when print, supported metrics are [cpu memory threads] (default "cpu")
  -top-count string
        the top count of process to be printed (default "7")
```

The output may look like below.

```
=========================================================================
Print snapshot with iteration count 0
-------------------------------------------------------------------------
Timestamp: 2022-08-01 02:47:38.007390324 +0000 UTC m=+0.000308004
=========================================================================
Print snapshot network stat:
-------------------------------------------------------------------------
TCP4-Connections        0
TCP4-Unknown
TCP4-Established        0
TCP4-SynSent            0
TCP4-SynRecv            0
TCP4-FinWait1           0
TCP4-FinWait2           0
TCP4-TimeWait           0
TCP4-Close              0
TCP4-CloseWait          0
TCP4-LastACK            0
TCP4-Listen             0
TCP4-Closing            0
TCP4-NewSynRecv         0
TCP6-Connections        2
TCP6-Unknown
TCP6-Established        1
TCP6-SynSent            0
TCP6-SynRecv            0
TCP6-FinWait1           0
TCP6-FinWait2           0
TCP6-TimeWait           0
TCP6-Close              0
TCP6-CloseWait          0
TCP6-LastACK            0
TCP6-Listen             1
TCP6-Closing            0
TCP6-NewSynRecv         0
=========================================================================
Print snapshot PID cmdline:
-------------------------------------------------------------------------
PID             CmdLine
7               java-jarapp.jar
2534            ./cachalot
1               /bin/bash./start.sh
2525            bash
=========================================================================
Print snapshot processes details:
-------------------------------------------------------------------------
PID                     7       2534    1       2525
Threads                 25      5       1       1
CPUUsage                0.0%    0.0%    0.0%    0.0%
MemoryUsage             14.6%   2.1%    0.0%    0.0%
VmSize                  4700MB  687MB   0MB     7MB
VmRss                   675MB   5MB     0MB     3MB
VmData                  1669MB  40MB    0MB     1MB
VmStk                   0MB     0MB     0MB     0MB
VmExe                   0MB     0MB     0MB     0MB
VmLib                   20MB    0MB     0MB     1MB
TCP4-Connections        0       0       0       0
TCP4-Unknown
TCP4-Established        0       0       0       0
TCP4-SynSent            0       0       0       0
TCP4-SynRecv            0       0       0       0
TCP4-FinWait1           0       0       0       0
TCP4-FinWait2           0       0       0       0
TCP4-TimeWait           0       0       0       0
TCP4-Close              0       0       0       0
TCP4-CloseWait          0       0       0       0
TCP4-LastACK            0       0       0       0
TCP4-Listen             0       0       0       0
TCP4-Closing            0       0       0       0
TCP4-NewSynRecv         0       0       0       0
TCP6-Connections        2       0       0       0
TCP6-Unknown
TCP6-Established        1       0       0       0
TCP6-SynSent            0       0       0       0
TCP6-SynRecv            0       0       0       0
TCP6-FinWait1           0       0       0       0
TCP6-FinWait2           0       0       0       0
TCP6-TimeWait           0       0       0       0
TCP6-Close              0       0       0       0
TCP6-CloseWait          0       0       0       0
TCP6-LastACK            0       0       0       0
TCP6-Listen             1       0       0       0
TCP6-Closing            0       0       0       0
TCP6-NewSynRecv         0       0       0       0


=========================================================================
Print snapshot with iteration count 1
-------------------------------------------------------------------------
Timestamp: 2022-08-01 02:47:53.02793989 +0000 UTC m=+15.020857570
=========================================================================
Print snapshot network stat:
-------------------------------------------------------------------------
TCP4-Connections        0
TCP4-Unknown
TCP4-Established        0
TCP4-SynSent            0
TCP4-SynRecv            0
TCP4-FinWait1           0
TCP4-FinWait2           0
TCP4-TimeWait           0
TCP4-Close              0
TCP4-CloseWait          0
TCP4-LastACK            0
TCP4-Listen             0
TCP4-Closing            0
TCP4-NewSynRecv         0
TCP6-Connections        2
TCP6-Unknown
TCP6-Established        1
TCP6-SynSent            0
TCP6-SynRecv            0
TCP6-FinWait1           0
TCP6-FinWait2           0
TCP6-TimeWait           0
TCP6-Close              0
TCP6-CloseWait          0
TCP6-LastACK            0
TCP6-Listen             1
TCP6-Closing            0
TCP6-NewSynRecv         0
=========================================================================
Print snapshot PID cmdline:
-------------------------------------------------------------------------
PID             CmdLine
2534            ./cachalot
7               java-jarapp.jar
2525            bash
1               /bin/bash./start.sh
=========================================================================
Print snapshot processes details:
-------------------------------------------------------------------------
PID                     2534    7       2525    1
Threads                 9       25      1       1
CPUUsage                0.0%    0.0%    0.0%    0.0%
MemoryUsage             2.1%    14.6%   0.0%    0.0%
VmSize                  688MB   4700MB  7MB     0MB
VmRss                   9MB     675MB   3MB     0MB
VmData                  45MB    1669MB  1MB     0MB
VmStk                   0MB     0MB     0MB     0MB
VmExe                   0MB     0MB     0MB     0MB
VmLib                   0MB     20MB    1MB     0MB
TCP4-Connections        0       0       0       0
TCP4-Unknown
TCP4-Established        0       0       0       0
TCP4-SynSent            0       0       0       0
TCP4-SynRecv            0       0       0       0
TCP4-FinWait1           0       0       0       0
TCP4-FinWait2           0       0       0       0
TCP4-TimeWait           0       0       0       0
TCP4-Close              0       0       0       0
TCP4-CloseWait          0       0       0       0
TCP4-LastACK            0       0       0       0
TCP4-Listen             0       0       0       0
TCP4-Closing            0       0       0       0
TCP4-NewSynRecv         0       0       0       0
TCP6-Connections        0       2       0       0
TCP6-Unknown
TCP6-Established        0       1       0       0
TCP6-SynSent            0       0       0       0
TCP6-SynRecv            0       0       0       0
TCP6-FinWait1           0       0       0       0
TCP6-FinWait2           0       0       0       0
TCP6-TimeWait           0       0       0       0
TCP6-Close              0       0       0       0
TCP6-CloseWait          0       0       0       0
TCP6-LastACK            0       0       0       0
TCP6-Listen             0       1       0       0
TCP6-Closing            0       0       0       0
TCP6-NewSynRecv         0       0       0       0

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

### Memory Usage

Cachalot will leverage VmSize to caculate the memory usage from host machine. It also records other memory related metrics from [/proc/PID/status](https://www.kernel.org/doc/html/latest/filesystems/proc.html).
