SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

# kernel-spdx-ids

kernel-spdx-ids can do the following:
1. optionally obtain and unpack a Linux® kernel source tarball
2. scan it for [SPDX® short-form license IDs](https://spdx.org/ids/)
3. produce a license report
4. optionally produce a corresponding SPDX document

Step 1 is optional. You can also download and unpack the kernel yourself.
See `Sample usage` below.

Steps 2-4 are basically just a wrapper around functionality provided by
[spdx/tools-golang](https://github.com/spdx/tools-golang/).

Please note that this is NOT performing any sort of license scanning, beyond
looking for SPDX short-form license IDs. Licenses will not be detected for
kernel source files that do not contain SPDX-License-Identifier tags.

For related information, see:
- [Linux kernel licensing rules](https://www.kernel.org/doc/html/latest/process/license-rules.html)
- [SPDX short-form license IDs](https://spdx.org/ids/)
- [SPDX license list](https://spdx.org/licenses/)

## Usage

### Sample usage

1. Download and unpack version 4.18 of the kernel at, e.g., `/tmp/linux-4.18/`
2. `go run kernel-spdx-ids.go -v 4.18 -p /tmp/linux-4.18 -s ./linux-4.18.spdx`

### Flags

`go run kernel-spdx-ids.go` with the following parameters:

* `-v <version>`, e.g. `-v 4.18`: required. kernel version to scan.
* `-d`: optional. if included, will download the kernel tarball of the version
  specified in the `-v` flag.
    * NOTE: May not work correctly on macOS systems. The kernel tarball will
      download correctly, but the unpacker will fail because it crashes on
      files in the same directory that have the same case-insensitive filename.
* `-p <path>`, e.g. `-p /path/to/kernel/`: required. path to the kernel source
  directory to be scanned. if `-d` is present, this is also where
  the kernel tarball will be unpacked after downloading.
* `-s <spdx-path>`, e.g. `-s linux-4.18.spdx`: optional. destination path to
  output SPDX file.

### >= 4.x kernels only

Currently it assumes you are looking for a 4.x or higher kernel. To my
knowledge, pre-4.x kernels did not contain SPDX short-form license IDs (at least
not in significant number).

## Contributing

Code contributions, PRs, issues, etc. are welcome. Please see `CONTRIBUTING.md`.

## Licenses

You may use kernel-spdx-ids under your choice of:
* Apache License, version 2.0 (**Apache-2.0**); OR
* GNU General Public License, version 2.0 or later (**GPL-2.0-or-later**).

## Trademarks

SPDX is a registered trademark of The Linux Foundation. Linux is a registered trademark of Linus Torvalds.
